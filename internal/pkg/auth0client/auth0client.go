package auth0client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gopkg.in/auth0.v5/management"
)

// Config holds the Auth0 configuration
type Config struct {
	Domain       string
	ClientID     string
	ClientSecret string
	Audience     string
	Connection   string // Database connection name
}

// Client represents the Auth0 client
type Client struct {
	config     *Config
	management *management.Management
	httpClient *http.Client
}

// User represents an Auth0 user
type User struct {
	UserID        string                 `json:"user_id"`
	Email         string                 `json:"email"`
	Name          string                 `json:"name"`
	Picture       string                 `json:"picture"`
	EmailVerified bool                   `json:"email_verified"`
	CreatedAt     time.Time              `json:"created_at"`
	UpdatedAt     time.Time              `json:"updated_at"`
	AppMetadata   map[string]interface{} `json:"app_metadata"`
	UserMetadata  map[string]interface{} `json:"user_metadata"`
}

// TokenClaims represents JWT token claims
type TokenClaims struct {
	Sub           string   `json:"sub"`
	Email         string   `json:"email"`
	EmailVerified bool     `json:"email_verified"`
	Name          string   `json:"name"`
	Picture       string   `json:"picture"`
	Permissions   []string `json:"permissions"`
	Scope         string   `json:"scope"`
	jwt.RegisteredClaims
}

// NewClient creates a new Auth0 client
func NewClient(config *Config) (*Client, error) {
	if config.Domain == "" || config.ClientID == "" || config.ClientSecret == "" {
		return nil, fmt.Errorf("missing required Auth0 configuration")
	}

	// Create management client
	mgmt, err := management.New(
		config.Domain,
		management.WithClientCredentials(config.ClientID, config.ClientSecret),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create Auth0 management client: %w", err)
	}

	return &Client{
		config:     config,
		management: mgmt,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}, nil
}

// ValidateToken validates a JWT token from Auth0
func (c *Client) ValidateToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Verify signing method
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// Get the JWKS and return the key
		return c.getSigningKey(token)
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		// Validate audience
		if c.config.Audience != "" {
			if !contains(claims.Audience, c.config.Audience) {
				return nil, fmt.Errorf("invalid audience")
			}
		}

		// Validate issuer
		expectedIssuer := fmt.Sprintf("https://%s/", c.config.Domain)
		if claims.Issuer != expectedIssuer {
			return nil, fmt.Errorf("invalid issuer")
		}

		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// GetUser retrieves a user by ID from Auth0
func (c *Client) GetUser(ctx context.Context, userID string) (*User, error) {
	user, err := c.management.User.Read(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return c.convertUser(user), nil
}

// GetUserByEmail retrieves a user by email from Auth0
func (c *Client) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	query := fmt.Sprintf("email:\"%s\"", email)
	users, err := c.management.User.Search(management.Query(query))
	if err != nil {
		return nil, fmt.Errorf("failed to search user by email: %w", err)
	}

	if len(users.Users) == 0 {
		return nil, fmt.Errorf("user not found")
	}

	return c.convertUser(users.Users[0]), nil
}

// CreateUser creates a new user in Auth0
func (c *Client) CreateUser(ctx context.Context, email, password, name string) (*User, error) {
	user := &management.User{
		Email:      &email,
		Password:   &password,
		Name:       &name,
		Connection: &c.config.Connection,
	}

	err := c.management.User.Create(user)
	if err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return c.convertUser(user), nil
}

// UpdateUser updates a user's information
func (c *Client) UpdateUser(ctx context.Context, userID string, updates map[string]interface{}) (*User, error) {
	user := &management.User{}

	if name, ok := updates["name"].(string); ok {
		user.Name = &name
	}
	if picture, ok := updates["picture"].(string); ok {
		user.Picture = &picture
	}
	if appMetadata, ok := updates["app_metadata"].(map[string]interface{}); ok {
		user.AppMetadata = appMetadata
	}
	if userMetadata, ok := updates["user_metadata"].(map[string]interface{}); ok {
		user.UserMetadata = userMetadata
	}

	err := c.management.User.Update(userID, user)
	if err != nil {
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	return c.GetUser(ctx, userID)
}

// DeleteUser deletes a user from Auth0
func (c *Client) DeleteUser(ctx context.Context, userID string) error {
	err := c.management.User.Delete(userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	return nil
}

// SetUserMetadata sets app metadata for a user
func (c *Client) SetUserMetadata(ctx context.Context, userID string, metadata map[string]interface{}) error {
	user := &management.User{
		AppMetadata: metadata,
	}

	err := c.management.User.Update(userID, user)
	if err != nil {
		return fmt.Errorf("failed to set user metadata: %w", err)
	}

	return nil
}

// AssignRoles assigns roles to a user
func (c *Client) AssignRoles(ctx context.Context, userID string, roleIDs []string) error {
	roles := make([]*management.Role, len(roleIDs))
	for i, roleID := range roleIDs {
		roles[i] = &management.Role{ID: &roleID}
	}

	err := c.management.User.AssignRoles(userID, roles)
	if err != nil {
		return fmt.Errorf("failed to assign roles: %w", err)
	}

	return nil
}

// RemoveRoles removes roles from a user
func (c *Client) RemoveRoles(ctx context.Context, userID string, roleIDs []string) error {
	roles := make([]*management.Role, len(roleIDs))
	for i, roleID := range roleIDs {
		roles[i] = &management.Role{ID: &roleID}
	}

	err := c.management.User.RemoveRoles(userID, roles)
	if err != nil {
		return fmt.Errorf("failed to remove roles: %w", err)
	}

	return nil
}

// GetUserRoles retrieves roles assigned to a user
func (c *Client) GetUserRoles(ctx context.Context, userID string) ([]*management.Role, error) {
	roles, err := c.management.User.Roles(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to get user roles: %w", err)
	}

	return roles.Roles, nil
}

// SendPasswordResetEmail sends a password reset email
func (c *Client) SendPasswordResetEmail(ctx context.Context, email string) error {
	ticket := &management.Ticket{
		Email: &email,
	}

	err := c.management.Ticket.ChangePassword(ticket)
	if err != nil {
		return fmt.Errorf("failed to send password reset email: %w", err)
	}

	return nil
}

// SendVerificationEmail sends an email verification
// Note: This functionality may need to be implemented differently based on your Auth0 setup
func (c *Client) SendVerificationEmail(ctx context.Context, userID string) error {
	// This would typically be done through Auth0's Actions or Rules
	// or by calling the Auth0 Management API directly
	return fmt.Errorf("send verification email not implemented - use Auth0 dashboard or custom implementation")
}

// BlockUser blocks a user
func (c *Client) BlockUser(ctx context.Context, userID string) error {
	blocked := true
	user := &management.User{
		Blocked: &blocked,
	}

	err := c.management.User.Update(userID, user)
	if err != nil {
		return fmt.Errorf("failed to block user: %w", err)
	}

	return nil
}

// UnblockUser unblocks a user
func (c *Client) UnblockUser(ctx context.Context, userID string) error {
	blocked := false
	user := &management.User{
		Blocked: &blocked,
	}

	err := c.management.User.Update(userID, user)
	if err != nil {
		return fmt.Errorf("failed to unblock user: %w", err)
	}

	return nil
}

// Helper methods

func (c *Client) convertUser(u *management.User) *User {
	user := &User{}

	if u.ID != nil {
		user.UserID = *u.ID
	}
	if u.Email != nil {
		user.Email = *u.Email
	}
	if u.Name != nil {
		user.Name = *u.Name
	}
	if u.Picture != nil {
		user.Picture = *u.Picture
	}
	if u.EmailVerified != nil {
		user.EmailVerified = *u.EmailVerified
	}
	if u.CreatedAt != nil {
		user.CreatedAt = *u.CreatedAt
	}
	if u.UpdatedAt != nil {
		user.UpdatedAt = *u.UpdatedAt
	}
	if u.AppMetadata != nil {
		user.AppMetadata = u.AppMetadata
	}
	if u.UserMetadata != nil {
		user.UserMetadata = u.UserMetadata
	}

	return user
}

func (c *Client) getSigningKey(token *jwt.Token) (interface{}, error) {
	// Get the kid from token header
	kid, ok := token.Header["kid"].(string)
	if !ok {
		return nil, fmt.Errorf("kid not found in token header")
	}

	// Fetch JWKS from Auth0
	jwksURL := fmt.Sprintf("https://%s/.well-known/jwks.json", c.config.Domain)
	resp, err := c.httpClient.Get(jwksURL)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch JWKS: %w", err)
	}
	defer resp.Body.Close()

	var jwks struct {
		Keys []struct {
			Kid string   `json:"kid"`
			Kty string   `json:"kty"`
			Use string   `json:"use"`
			N   string   `json:"n"`
			E   string   `json:"e"`
			X5c []string `json:"x5c"`
		} `json:"keys"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		return nil, fmt.Errorf("failed to decode JWKS: %w", err)
	}

	// Find the key with matching kid
	for _, key := range jwks.Keys {
		if key.Kid == kid {
			if len(key.X5c) > 0 {
				// Parse certificate
				cert := "-----BEGIN CERTIFICATE-----\n" + key.X5c[0] + "\n-----END CERTIFICATE-----"
				return jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			}
		}
	}

	return nil, fmt.Errorf("signing key not found")
}

// ExtractTokenFromHeader extracts the Bearer token from Authorization header
func ExtractTokenFromHeader(authHeader string) (string, error) {
	if authHeader == "" {
		return "", fmt.Errorf("authorization header is empty")
	}

	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", fmt.Errorf("invalid authorization header format")
	}

	return parts[1], nil
}

// contains checks if a slice of strings contains a target string
func contains(slice []string, target string) bool {
	for _, item := range slice {
		if item == target {
			return true
		}
	}
	return false
}
