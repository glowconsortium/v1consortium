// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"time"
	"v1consortium/internal/model"
)

type (
	ISessionManager interface {
		CreateUserSession(ctx context.Context, userId string, email string, sessionID string, accessToken string, refreshToken string) error
		DestroyUserSession(ctx context.Context) error
		GetUserIdFromSession(ctx context.Context) string
		// GetAccessTokenFromSession retrieves the access token from session
		GetAccessTokenFromSession(ctx context.Context) string
		// GetRefreshTokenFromSession retrieves the refresh token from session
		GetRefreshTokenFromSession(ctx context.Context) string
		// GetEmailFromSession retrieves the email from session
		GetEmailFromSession(ctx context.Context) string
		ValidateSession(ctx context.Context) bool
		// ValidateAndSyncSession performs comprehensive session validation and synchronization with Supabase
		ValidateAndSyncSession(ctx context.Context) error
		// GetSessionInfo returns comprehensive session information
		GetSessionInfo(ctx context.Context) map[string]interface{}
		// RefreshSessionTokens manually refreshes the session tokens
		RefreshSessionTokens(ctx context.Context) error
		// ValidateSessionAge checks if the session is too old and needs re-authentication
		ValidateSessionAge(ctx context.Context, maxAge time.Duration) error
		// ValidateAndSyncSessionWithResponse performs validation and returns detailed response like proto service
		ValidateAndSyncSessionWithResponse(ctx context.Context) (*model.SessionValidationResponse, error)
		// GetActiveSessionsInfo returns information about active sessions for a user
		GetActiveSessionsInfo(ctx context.Context, userID string) ([]*model.SessionInfo, error)
		// RevokeSession revokes a specific session
		RevokeSession(ctx context.Context, sessionID string) error
		// RevokeAllSessions revokes all sessions for a user except optionally the current one
		RevokeAllSessions(ctx context.Context, userID string, exceptSessionID string) (int32, error)
		// ValidateTokenRequest validates a token and returns its metadata
		ValidateTokenRequest(ctx context.Context, token string, expectedType string) (*model.TokenValidationResponse, error)
		// CreateTokenForUser creates a token for a specific user and type
		CreateTokenForUser(ctx context.Context, userID string, tokenType string, expiresAt time.Time, metadata map[string]string) (*model.TokenCreateResponse, error)
		// CleanupExpiredTokens removes expired tokens from storage
		CleanupExpiredTokens(ctx context.Context) (int32, error)
		// CreateSession creates a new session (RPC: CreateSession)
		CreateSession(ctx context.Context, req *model.CreateSessionRequest) (*model.CreateSessionResponse, error)
		// ValidateSessionRPC validates an existing session (RPC: ValidateSession)
		ValidateSessionRPC(ctx context.Context, req *model.ValidateSessionRequest) (*model.SessionValidationResponse, error)
		// RefreshSession refreshes session tokens (RPC: RefreshSession)
		RefreshSession(ctx context.Context, req *model.RefreshSessionRequest) (*model.RefreshSessionResponse, error)
		// RevokeSessionRPC revokes a specific session (RPC: RevokeSession)
		RevokeSessionRPC(ctx context.Context, req *model.RevokeSessionRequest) (*model.RevokeSessionResponse, error)
		// GetActiveSessions retrieves active sessions for a user (RPC: GetActiveSessions)
		GetActiveSessions(ctx context.Context, req *model.GetActiveSessionsRequest) (*model.GetActiveSessionsResponse, error)
		// RevokeAllSessionsRPC revokes all sessions for a user (RPC: RevokeAllSessions)
		RevokeAllSessionsRPC(ctx context.Context, req *model.RevokeAllSessionsRequest) (*model.RevokeAllSessionsResponse, error)
		// CreateToken creates a new token (RPC: CreateToken)
		CreateToken(ctx context.Context, req *model.CreateTokenRequest) (*model.TokenCreateResponse, error)
		// ValidateToken validates a token (RPC: ValidateToken)
		ValidateToken(ctx context.Context, req *model.ValidateTokenRequest) (*model.TokenValidationResponse, error)
		// RevokeToken revokes a specific token (RPC: RevokeToken)
		RevokeToken(ctx context.Context, req *model.RevokeTokenRequest) (*model.RevokeTokenResponse, error)
		// CleanupExpiredTokensRPC cleans up expired tokens (RPC: CleanupExpiredTokens)
		CleanupExpiredTokensRPC(ctx context.Context, req *model.CleanupExpiredTokensRequest) (*model.CleanupExpiredTokensResponse, error)
	}
)

var (
	localSessionManager ISessionManager
)

func SessionManager() ISessionManager {
	if localSessionManager == nil {
		panic("implement not found for interface ISessionManager, forgot register?")
	}
	return localSessionManager
}

func RegisterSessionManager(i ISessionManager) {
	localSessionManager = i
}
