package model

import (
	"time"
	"v1consortium/internal/model/entity"
)

// Session Management Request/Response Models

// CreateSessionRequest represents the request to create a new session
type CreateSessionRequest struct {
	UserID    string    `json:"user_id"`
	IPAddress string    `json:"ip_address"`
	UserAgent string    `json:"user_agent"`
	ExpiresAt time.Time `json:"expires_at"`
}

// CreateSessionResponse represents the response from session creation
type CreateSessionResponse struct {
	SessionID    string `json:"session_id"`
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginResponse struct {
	Session   *CreateSessionResponse `json:"session"`
	User      *entity.UserProfiles   `json:"user"`
	ExpiresAt time.Time              `json:"expires_at"`
}

// ValidateSessionRequest represents the request to validate a session
type ValidateSessionRequest struct {
	SessionID   string `json:"session_id"`
	AccessToken string `json:"access_token"`
}

// SessionValidationResponse represents the response from session validation
type SessionValidationResponse struct {
	Valid           bool      `json:"valid"`
	UserID          string    `json:"user_id"`
	Email           string    `json:"email"`
	OrganizationID  string    `json:"organization_id"`
	Role            string    `json:"role"`
	Permissions     []string  `json:"permissions"`
	ExpiresAt       time.Time `json:"expires_at"`
	TokensRefreshed bool      `json:"tokens_refreshed"`
	NewAccessToken  string    `json:"new_access_token,omitempty"`
	NewRefreshToken string    `json:"new_refresh_token,omitempty"`
}

// RefreshSessionRequest represents the request to refresh a session
type RefreshSessionRequest struct {
	SessionID    string `json:"session_id"`
	RefreshToken string `json:"refresh_token"`
}

// RefreshSessionResponse represents the response from session refresh
type RefreshSessionResponse struct {
	NewAccessToken  string    `json:"new_access_token"`
	NewRefreshToken string    `json:"new_refresh_token"`
	ExpiresAt       time.Time `json:"expires_at"`
}

// RevokeSessionRequest represents the request to revoke a session
type RevokeSessionRequest struct {
	SessionID string `json:"session_id"`
}

// RevokeSessionResponse represents the response from session revocation
type RevokeSessionResponse struct {
	Message string `json:"message"`
}

// GetActiveSessionsRequest represents the request to get active sessions
type GetActiveSessionsRequest struct {
	UserID string `json:"user_id"`
}

// SessionInfo represents information about a user session
type SessionInfo struct {
	SessionID    string    `json:"session_id"`
	UserID       string    `json:"user_id"`
	IPAddress    string    `json:"ip_address"`
	UserAgent    string    `json:"user_agent"`
	CreatedAt    time.Time `json:"created_at"`
	LastAccessed time.Time `json:"last_accessed"`
	ExpiresAt    time.Time `json:"expires_at"`
	IsCurrent    bool      `json:"is_current"`
}

// GetActiveSessionsResponse represents the response from getting active sessions
type GetActiveSessionsResponse struct {
	Sessions []*SessionInfo `json:"sessions"`
}

// RevokeAllSessionsRequest represents the request to revoke all sessions
type RevokeAllSessionsRequest struct {
	UserID          string `json:"user_id"`
	ExceptSessionID string `json:"except_session_id"`
}

// RevokeAllSessionsResponse represents the response from revoking all sessions
type RevokeAllSessionsResponse struct {
	RevokedCount int32  `json:"revoked_count"`
	Message      string `json:"message"`
}

// Token Management Request/Response Models

// CreateTokenRequest represents the request to create a token
type CreateTokenRequest struct {
	UserID    string            `json:"user_id"`
	TokenType string            `json:"token_type"` // "access", "refresh", "reset", "verification", "api"
	ExpiresAt time.Time         `json:"expires_at"`
	Metadata  map[string]string `json:"metadata"`
}

// TokenCreateResponse represents the response from token creation
type TokenCreateResponse struct {
	Token   string `json:"token"`
	TokenID string `json:"token_id"`
}

// ValidateTokenRequest represents the request to validate a token
type ValidateTokenRequest struct {
	Token        string `json:"token"`
	ExpectedType string `json:"expected_type"`
}

// TokenValidationResponse represents the response from token validation
type TokenValidationResponse struct {
	Valid     bool              `json:"valid"`
	UserID    string            `json:"user_id"`
	TokenType string            `json:"token_type"`
	ExpiresAt time.Time         `json:"expires_at"`
	Metadata  map[string]string `json:"metadata"`
}

// RevokeTokenRequest represents the request to revoke a token
type RevokeTokenRequest struct {
	Token string `json:"token"`
}

// RevokeTokenResponse represents the response from token revocation
type RevokeTokenResponse struct {
	Message string `json:"message"`
}

// CleanupExpiredTokensRequest represents the request to cleanup expired tokens
type CleanupExpiredTokensRequest struct {
	// Empty - cleanup all expired tokens
}

// CleanupExpiredTokensResponse represents the response from cleanup
type CleanupExpiredTokensResponse struct {
	CleanedCount int32 `json:"cleaned_count"`
}
