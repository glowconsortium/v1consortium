package supabaseservice

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"
	"v1consortium/internal/pkg/supabaseclient"
	"v1consortium/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
)

func new() service.ISupabaseService {
	return &sSupabaseService{}
}

func init() {
	service.RegisterSupabaseService(new())
}

type sSupabaseService struct{}

func (s *sSupabaseService) GetAnonClient(ctx context.Context) (*supabase.Client, error) {
	config := s.GetSupabaseConfig(ctx)

	client, err := supabaseclient.NewSupabaseClient(*config)
	if err != nil {
		return nil, err
	}
	return client.GetAnonClient(), nil
}

func (s *sSupabaseService) GetServiceClient(ctx context.Context) (*supabase.Client, error) {
	config := s.GetSupabaseConfig(ctx)

	client, err := supabaseclient.NewSupabaseClient(*config)
	if err != nil {
		return nil, err
	}
	return client.GetServiceClient(), nil
}

func (s *sSupabaseService) GetUserAuthenticatedClient(ctx context.Context, accessToken string) (*supabase.Client, error) {
	// Validate access token format and presence
	if accessToken == "" {
		return nil, errors.New("access token cannot be empty")
	}

	// Basic JWT format validation (should have 3 parts separated by dots)
	tokenParts := strings.Split(accessToken, ".")
	if len(tokenParts) != 3 {
		return nil, errors.New("invalid access token format")
	}

	anonclient, err := s.GetAnonClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get anonymous client: %w", err)
	}

	anonclient.UpdateAuthSession(types.Session{
		AccessToken: accessToken,
	})

	return anonclient, nil
}

func (s *sSupabaseService) GetSupabaseConfig(ctx context.Context) *supabaseclient.SupabaseConfig {

	Url := g.Cfg().MustGet(ctx, "supabase.url").String()
	PublicApiKey := g.Cfg().MustGet(ctx, "supabase.publicApiKey").String()
	SecretApiKey := g.Cfg().MustGet(ctx, "supabase.secretApiKey").String()

	g.Log().Debugf(ctx, "Supabase Config - URL: %s, PublicApiKey: %s, SecretApiKey: %s", Url, PublicApiKey, SecretApiKey)

	return &supabaseclient.SupabaseConfig{
		Url:          Url,
		PublicApiKey: PublicApiKey,
		SecretApiKey: SecretApiKey,
	}
}

func (s *sSupabaseService) GetUser(ctx context.Context, userID uuid.UUID) (*types.AdminGetUserResponse, error) {
	client, err := s.GetServiceClient(ctx)
	if err != nil {
		return nil, err
	}
	user, err := client.Auth.AdminGetUser(types.AdminGetUserRequest{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *sSupabaseService) CreateUser(ctx context.Context, email, password string, userMetadata map[string]interface{}) (*types.AdminCreateUserResponse, error) {
	client, err := s.GetServiceClient(ctx)
	if err != nil {
		return nil, err
	}
	user, err := client.Auth.AdminCreateUser(types.AdminCreateUserRequest{
		Email:        email,
		Password:     &password,
		UserMetadata: userMetadata,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *sSupabaseService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	client, err := s.GetServiceClient(ctx)
	if err != nil {
		return err
	}
	err = client.Auth.AdminDeleteUser(types.AdminDeleteUserRequest{
		UserID: userID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *sSupabaseService) UpdateUser(ctx context.Context, userID uuid.UUID, email *string, password *string, userMetadata map[string]interface{}) (*types.AdminUpdateUserResponse, error) {
	client, err := s.GetServiceClient(ctx)
	if err != nil {
		return nil, err
	}
	user, err := client.Auth.AdminUpdateUser(types.AdminUpdateUserRequest{
		UserID:       userID,
		Email:        *email,
		Password:     *password,
		UserMetadata: userMetadata,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *sSupabaseService) SignIn(ctx context.Context, email, password string) (*types.TokenResponse, error) {
	// Validate input parameters
	if email == "" {
		return nil, errors.New("email cannot be empty")
	}
	if password == "" {
		return nil, errors.New("password cannot be empty")
	}

	// Basic email format validation
	if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
		return nil, errors.New("invalid email format")
	}

	client, err := s.GetAnonClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get anonymous client: %w", err)
	}

	session, err := client.Auth.SignInWithEmailPassword(email, password)
	if err != nil {
		g.Log().Errorf(ctx, "Sign-in failed for email %s: %v", email, err)
		return nil, fmt.Errorf("authentication failed: %w", err)
	}

	// Validate the returned session
	if session == nil {
		return nil, errors.New("authentication succeeded but no session returned")
	}
	if session.AccessToken == "" {
		return nil, errors.New("authentication succeeded but no access token returned")
	}

	g.Log().Infof(ctx, "User signed in successfully: %s", email)
	return session, nil
}

func (s *sSupabaseService) SignUp(ctx context.Context, email, password string, userMetadata map[string]interface{}) (*types.SignupResponse, error) {
	client, err := s.GetAnonClient(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Auth.Signup(types.SignupRequest{
		Email:    email,
		Password: password,
		Data:     userMetadata,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// social signup
func (s *sSupabaseService) SignUpWithProvider(ctx context.Context, provider string, redirectTo string, scopes string) (*types.AuthorizeResponse, error) {
	client, err := s.GetAnonClient(ctx)
	if err != nil {
		return nil, err
	}

	var providerType types.Provider
	var flowType types.FlowType

	switch strings.ToLower(provider) {
	case "google":
		providerType = types.ProviderGoogle
		flowType = types.FlowImplicit
	case "microsoft":
		providerType = types.ProviderAzure
		flowType = types.FlowImplicit
	default:
		return nil, errors.New("unsupported provider")
	}

	// Validate redirect URL format (basic check)
	if redirectTo == "" || !(strings.HasPrefix(redirectTo, "http://") || strings.HasPrefix(redirectTo, "https://")) {
		return nil, errors.New("invalid redirect URL format")
	}

	resp, err := client.Auth.Authorize(types.AuthorizeRequest{
		Provider: providerType,
		FlowType: flowType,
		Scopes:   scopes,
	})

	return resp, err
}

func (s *sSupabaseService) RefreshToken(ctx context.Context, refreshToken string) (*types.TokenResponse, error) {
	// Validate refresh token
	if refreshToken == "" {
		return nil, errors.New("refresh token cannot be empty")
	}

	client, err := s.GetAnonClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get anonymous client: %w", err)
	}

	session, err := client.Auth.RefreshToken(refreshToken)
	if err != nil {
		g.Log().Errorf(ctx, "Token refresh failed: %v", err)
		return nil, fmt.Errorf("token refresh failed: %w", err)
	}

	// Validate the refreshed session
	if session == nil {
		return nil, errors.New("token refresh succeeded but no session returned")
	}
	if session.AccessToken == "" {
		return nil, errors.New("token refresh succeeded but no access token returned")
	}

	g.Log().Debugf(ctx, "Token refreshed successfully")
	return session, nil
}

func (s *sSupabaseService) SignOut(ctx context.Context, accessToken string) error {
	client, err := s.GetUserAuthenticatedClient(ctx, accessToken)
	if err != nil {
		return err
	}
	err = client.Auth.Logout()
	if err != nil {
		return err
	}
	return nil
}

func (s *sSupabaseService) GetUserFromToken(ctx context.Context, accessToken string) (*types.UserResponse, error) {
	// Validate access token
	if accessToken == "" {
		return nil, errors.New("access token cannot be empty")
	}

	client, err := s.GetUserAuthenticatedClient(ctx, accessToken)
	if err != nil {
		return nil, fmt.Errorf("failed to get authenticated client: %w", err)
	}

	user, err := client.Auth.GetUser()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to get user from token: %v", err)
		return nil, fmt.Errorf("failed to get user from token: %w", err)
	}

	// Validate user response
	if user == nil {
		return nil, errors.New("user retrieval succeeded but no user returned")
	}
	if user.ID == uuid.Nil {
		return nil, errors.New("user retrieved but has invalid ID")
	}

	return user, nil
}

// ValidateTokenExpiry checks if the access token is expired
func (s *sSupabaseService) ValidateTokenExpiry(ctx context.Context, accessToken string) error {
	if accessToken == "" {
		return errors.New("access token cannot be empty")
	}

	user, err := s.GetUserFromToken(ctx, accessToken)
	if err != nil {
		return fmt.Errorf("token validation failed: %w", err)
	}

	if user == nil {
		return errors.New("invalid token: no user found")
	}

	return nil
}

// ValidateAndRefreshSession validates the current session and refreshes if needed
func (s *sSupabaseService) ValidateAndRefreshSession(ctx context.Context, accessToken, refreshToken string) (*types.TokenResponse, error) {
	if accessToken == "" {
		return nil, errors.New("access token cannot be empty")
	}
	if refreshToken == "" {
		return nil, errors.New("refresh token cannot be empty")
	}

	// First try to validate the current access token
	err := s.ValidateTokenExpiry(ctx, accessToken)
	if err == nil {
		// Token is still valid, return existing session
		return &types.TokenResponse{
			Session: types.Session{
				AccessToken:  accessToken,
				RefreshToken: refreshToken,
			},
		}, nil
	}

	g.Log().Debugf(ctx, "Access token expired or invalid, attempting refresh")

	// Token is expired or invalid, try to refresh
	newSession, err := s.RefreshToken(ctx, refreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh expired token: %w", err)
	}

	return newSession, nil
}

// ValidateSessionSecurity performs additional security checks on the session
func (s *sSupabaseService) ValidateSessionSecurity(ctx context.Context, accessToken string, expectedUserID uuid.UUID) error {
	if accessToken == "" {
		return errors.New("access token cannot be empty")
	}
	if expectedUserID == uuid.Nil {
		return errors.New("expected user ID cannot be empty")
	}

	user, err := s.GetUserFromToken(ctx, accessToken)
	if err != nil {
		return fmt.Errorf("failed to validate session security: %w", err)
	}

	// Check if the token belongs to the expected user
	if user.ID != expectedUserID {
		g.Log().Warningf(ctx, "Session security violation: token user ID %s does not match expected %s", user.ID, expectedUserID)
		return errors.New("session security violation: user ID mismatch")
	}

	// Check if user account is confirmed
	if user.EmailConfirmedAt == nil {
		return errors.New("user email not confirmed")
	}

	// Check if user is not banned
	if user.BannedUntil != nil && user.BannedUntil.After(time.Now()) {
		return fmt.Errorf("user is banned until: %v", user.BannedUntil)
	}

	return nil
}
