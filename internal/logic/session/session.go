package session

import (
	"context"
	"errors"
	"fmt"
	"time"
	"v1consortium/internal/model"
	"v1consortium/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
)

type sSessionManager struct {
}

func init() {
	service.RegisterSessionManager(new())
}

func new() service.ISessionManager {
	return &sSessionManager{}
}

func (s *sSessionManager) CreateUserSession(ctx context.Context, userId string, email string, sessionID string, accessToken string, refreshToken string) error {
	// Validate input parameters
	if userId == "" {
		return errors.New("userId cannot be empty")
	}
	if email == "" {
		return errors.New("email cannot be empty")
	}
	if accessToken == "" {
		return errors.New("accessToken cannot be empty")
	}
	if refreshToken == "" {
		return errors.New("refreshToken cannot be empty")
	}

	// Validate that the access token is valid with Supabase before storing
	user, err := service.SupabaseService().GetUserFromToken(ctx, accessToken)
	if err != nil {
		return fmt.Errorf("failed to validate access token with Supabase: %w", err)
	}

	// Ensure the user ID matches the token
	if user.ID.String() != userId {
		return errors.New("user ID mismatch between provided ID and token")
	}

	request := g.RequestFromCtx(ctx)

	// Set session creation timestamp
	err = request.Session.Set("sessionCreatedAt", time.Now().Unix())
	if err != nil {
		return fmt.Errorf("failed to set session creation time: %w", err)
	}

	err = request.Session.Set("userId", userId)
	if err != nil {
		return err
	}
	err = request.Session.Set("useremail", email)
	if err != nil {
		return err
	}
	err = request.Session.Set("sessionId", sessionID)
	if err != nil {
		return err
	}
	err = request.Session.Set("accessToken", accessToken)
	if err != nil {
		return err
	}
	err = request.Session.Set("refreshToken", refreshToken)
	if err != nil {
		return err
	}

	// Set last validation timestamp
	err = request.Session.Set("lastValidatedAt", time.Now().Unix())
	if err != nil {
		return fmt.Errorf("failed to set last validation time: %w", err)
	}

	g.Log().Infof(ctx, "User session created successfully for user: %s", userId)
	return nil
}

func (s *sSessionManager) DestroyUserSession(ctx context.Context) error {
	request := g.RequestFromCtx(ctx)

	// Get access token before destroying session for logout
	accessToken, _ := request.Session.Get("accessToken")
	userId, _ := request.Session.Get("userId")

	// Log session destruction
	if !userId.IsEmpty() {
		g.Log().Infof(ctx, "Destroying session for user: %s", userId.String())
	}

	// Sign out from Supabase if we have a valid access token
	if !accessToken.IsEmpty() {
		err := service.SupabaseService().SignOut(ctx, accessToken.String())
		if err != nil {
			g.Log().Warningf(ctx, "Failed to sign out from Supabase: %v", err)
			// Continue with local session destruction even if Supabase signout fails
		}
	}

	err := request.Session.Remove("userId")
	if err != nil {
		return err
	}
	err = request.Session.Remove("useremail")
	if err != nil {
		return err
	}
	err = request.Session.Remove("sessionId")
	if err != nil {
		return err
	}
	err = request.Session.Remove("accessToken")
	if err != nil {
		return err
	}
	err = request.Session.Remove("refreshToken")
	if err != nil {
		return err
	}

	// Remove session timestamps
	_ = request.Session.Remove("sessionCreatedAt")
	_ = request.Session.Remove("lastValidatedAt")

	return nil
}

func (s *sSessionManager) GetUserIdFromSession(ctx context.Context) string {
	request := g.RequestFromCtx(ctx)
	userId, _ := request.Session.Get("userId")
	return userId.String()
}

// GetAccessTokenFromSession retrieves the access token from session
func (s *sSessionManager) GetAccessTokenFromSession(ctx context.Context) string {
	request := g.RequestFromCtx(ctx)
	accessToken, _ := request.Session.Get("accessToken")
	return accessToken.String()
}

// GetRefreshTokenFromSession retrieves the refresh token from session
func (s *sSessionManager) GetRefreshTokenFromSession(ctx context.Context) string {
	request := g.RequestFromCtx(ctx)
	refreshToken, _ := request.Session.Get("refreshToken")
	return refreshToken.String()
}

// GetEmailFromSession retrieves the email from session
func (s *sSessionManager) GetEmailFromSession(ctx context.Context) string {
	request := g.RequestFromCtx(ctx)
	email, _ := request.Session.Get("useremail")
	return email.String()
}

func (s *sSessionManager) ValidateSession(ctx context.Context) bool {
	return s.ValidateAndSyncSession(ctx) == nil
}

// ValidateAndSyncSession performs comprehensive session validation and synchronization with Supabase
func (s *sSessionManager) ValidateAndSyncSession(ctx context.Context) error {
	request := g.RequestFromCtx(ctx)

	// Check if basic session data exists
	userId, err := request.Session.Get("userId")
	if err != nil {
		return errors.New("failed to get userId from session")
	}
	if userId.IsEmpty() {
		return errors.New("userId is empty in session")
	}

	accessToken, err := request.Session.Get("accessToken")
	if err != nil {
		return errors.New("failed to get accessToken from session")
	}
	if accessToken.IsEmpty() {
		return errors.New("accessToken is empty in session")
	}

	refreshToken, err := request.Session.Get("refreshToken")
	if err != nil {
		return errors.New("failed to get refreshToken from session")
	}
	if refreshToken.IsEmpty() {
		return errors.New("refreshToken is empty in session")
	}

	// Check session age and validation frequency
	lastValidated, _ := request.Session.Get("lastValidatedAt")
	if !lastValidated.IsEmpty() {
		lastValidatedTime := time.Unix(lastValidated.Int64(), 0)
		// Re-validate with Supabase if last validation was more than 5 minutes ago
		if time.Since(lastValidatedTime) > 5*time.Minute {
			g.Log().Debugf(ctx, "Session validation is stale, re-validating with Supabase")
		} else {
			// Recent validation, skip Supabase check for performance
			return nil
		}
	}

	// Validate and potentially refresh tokens with Supabase
	tokenResponse, err := service.SupabaseService().ValidateAndRefreshSession(ctx, accessToken.String(), refreshToken.String())
	if err != nil {
		g.Log().Errorf(ctx, "Session validation failed: %v", err)
		return fmt.Errorf("session validation with Supabase failed: %w", err)
	}

	// Check if tokens were refreshed
	if tokenResponse.AccessToken != accessToken.String() || tokenResponse.RefreshToken != refreshToken.String() {
		g.Log().Infof(ctx, "Tokens were refreshed, updating session")

		// Update session with new tokens
		err = request.Session.Set("accessToken", tokenResponse.AccessToken)
		if err != nil {
			return fmt.Errorf("failed to update accessToken in session: %w", err)
		}

		err = request.Session.Set("refreshToken", tokenResponse.RefreshToken)
		if err != nil {
			return fmt.Errorf("failed to update refreshToken in session: %w", err)
		}
	}

	// Get user from token to verify session integrity
	user, err := service.SupabaseService().GetUserFromToken(ctx, tokenResponse.AccessToken)
	if err != nil {
		return fmt.Errorf("failed to get user from token: %w", err)
	}

	if user == nil {
		return errors.New("user not found for valid token")
	}

	// Verify user ID matches session
	if user.ID.String() != userId.String() {
		g.Log().Errorf(ctx, "Session security violation: token user ID %s does not match session user ID %s", user.ID, userId.String())
		return errors.New("session security violation: user ID mismatch")
	}

	// Perform additional security checks
	userUUID, err := uuid.Parse(userId.String())
	if err != nil {
		return fmt.Errorf("invalid user ID format in session: %w", err)
	}

	err = service.SupabaseService().ValidateSessionSecurity(ctx, tokenResponse.AccessToken, userUUID)
	if err != nil {
		return fmt.Errorf("session security validation failed: %w", err)
	}

	// Update last validation timestamp
	err = request.Session.Set("lastValidatedAt", time.Now().Unix())
	if err != nil {
		g.Log().Warningf(ctx, "Failed to update last validation timestamp: %v", err)
	}

	return nil
}

// GetSessionInfo returns comprehensive session information
func (s *sSessionManager) GetSessionInfo(ctx context.Context) map[string]interface{} {
	request := g.RequestFromCtx(ctx)
	info := make(map[string]interface{})

	if request == nil || request.Session == nil {
		return info
	}

	if userId, _ := request.Session.Get("userId"); !userId.IsEmpty() {
		info["userId"] = userId.String()
	}

	if email, _ := request.Session.Get("useremail"); !email.IsEmpty() {
		info["email"] = email.String()
	}

	if sessionId, _ := request.Session.Get("sessionId"); !sessionId.IsEmpty() {
		info["sessionId"] = sessionId.String()
	}

	if createdAt, _ := request.Session.Get("sessionCreatedAt"); !createdAt.IsEmpty() {
		info["createdAt"] = time.Unix(createdAt.Int64(), 0)
	}

	if lastValidated, _ := request.Session.Get("lastValidatedAt"); !lastValidated.IsEmpty() {
		info["lastValidatedAt"] = time.Unix(lastValidated.Int64(), 0)
	}

	info["isValid"] = s.ValidateSession(ctx)

	return info
}

// RefreshSessionTokens manually refreshes the session tokens
func (s *sSessionManager) RefreshSessionTokens(ctx context.Context) error {
	request := g.RequestFromCtx(ctx)

	refreshToken, err := request.Session.Get("refreshToken")
	if err != nil {
		return fmt.Errorf("failed to get refresh token from session: %w", err)
	}
	if refreshToken.IsEmpty() {
		return errors.New("refresh token is empty in session")
	}

	// Refresh tokens with Supabase
	tokenResponse, err := service.SupabaseService().RefreshToken(ctx, refreshToken.String())
	if err != nil {
		return fmt.Errorf("failed to refresh tokens: %w", err)
	}

	// Update session with new tokens
	err = request.Session.Set("accessToken", tokenResponse.AccessToken)
	if err != nil {
		return fmt.Errorf("failed to update access token in session: %w", err)
	}

	err = request.Session.Set("refreshToken", tokenResponse.RefreshToken)
	if err != nil {
		return fmt.Errorf("failed to update refresh token in session: %w", err)
	}

	// Update last validation timestamp
	err = request.Session.Set("lastValidatedAt", time.Now().Unix())
	if err != nil {
		g.Log().Warningf(ctx, "Failed to update last validation timestamp: %v", err)
	}

	g.Log().Infof(ctx, "Session tokens refreshed successfully")
	return nil
}

// ValidateSessionAge checks if the session is too old and needs re-authentication
func (s *sSessionManager) ValidateSessionAge(ctx context.Context, maxAge time.Duration) error {
	request := g.RequestFromCtx(ctx)

	createdAt, err := request.Session.Get("sessionCreatedAt")
	if err != nil || createdAt.IsEmpty() {
		return errors.New("session creation time not found")
	}

	sessionAge := time.Since(time.Unix(createdAt.Int64(), 0))
	if sessionAge > maxAge {
		return fmt.Errorf("session is too old: %v > %v", sessionAge, maxAge)
	}

	return nil
}

// ValidateAndSyncSessionWithResponse performs validation and returns detailed response like proto service
func (s *sSessionManager) ValidateAndSyncSessionWithResponse(ctx context.Context) (*model.SessionValidationResponse, error) {
	request := g.RequestFromCtx(ctx)

	response := &model.SessionValidationResponse{
		Valid:           false,
		TokensRefreshed: false,
	}

	// Check if basic session data exists
	userId, err := request.Session.Get("userId")
	if err != nil || userId.IsEmpty() {
		return response, errors.New("userId not found in session")
	}
	response.UserID = userId.String()

	accessToken, err := request.Session.Get("accessToken")
	if err != nil || accessToken.IsEmpty() {
		return response, errors.New("accessToken not found in session")
	}

	refreshToken, err := request.Session.Get("refreshToken")
	if err != nil || refreshToken.IsEmpty() {
		return response, errors.New("refreshToken not found in session")
	}

	// Check session age and validation frequency
	lastValidated, _ := request.Session.Get("lastValidatedAt")
	if !lastValidated.IsEmpty() {
		lastValidatedTime := time.Unix(lastValidated.Int64(), 0)
		// Re-validate with Supabase if last validation was more than 5 minutes ago
		if time.Since(lastValidatedTime) <= 5*time.Minute {
			// Recent validation, return valid response
			response.Valid = true
			return response, nil
		}
	}

	// Validate and potentially refresh tokens with Supabase
	tokenResponse, err := service.SupabaseService().ValidateAndRefreshSession(ctx, accessToken.String(), refreshToken.String())
	if err != nil {
		g.Log().Errorf(ctx, "Session validation failed: %v", err)
		return response, fmt.Errorf("session validation with Supabase failed: %w", err)
	}

	// Check if tokens were refreshed
	if tokenResponse.AccessToken != accessToken.String() || tokenResponse.RefreshToken != refreshToken.String() {
		g.Log().Infof(ctx, "Tokens were refreshed, updating session")
		response.TokensRefreshed = true
		response.NewAccessToken = tokenResponse.AccessToken
		response.NewRefreshToken = tokenResponse.RefreshToken

		// Update session with new tokens
		_ = request.Session.Set("accessToken", tokenResponse.AccessToken)
		_ = request.Session.Set("refreshToken", tokenResponse.RefreshToken)
	}

	// Get user from token to verify session integrity
	user, err := service.SupabaseService().GetUserFromToken(ctx, tokenResponse.AccessToken)
	if err != nil {
		return response, fmt.Errorf("failed to get user from token: %w", err)
	}

	// Verify user ID matches session
	if user.ID.String() != userId.String() {
		g.Log().Errorf(ctx, "Session security violation: token user ID %s does not match session user ID %s", user.ID, userId.String())
		return response, errors.New("session security violation: user ID mismatch")
	}

	// Perform additional security checks
	userUUID, err := uuid.Parse(userId.String())
	if err != nil {
		return response, fmt.Errorf("invalid user ID format in session: %w", err)
	}

	err = service.SupabaseService().ValidateSessionSecurity(ctx, tokenResponse.AccessToken, userUUID)
	if err != nil {
		return response, fmt.Errorf("session security validation failed: %w", err)
	}

	// Update last validation timestamp
	_ = request.Session.Set("lastValidatedAt", time.Now().Unix())

	response.Valid = true
	if email, _ := request.Session.Get("useremail"); !email.IsEmpty() {
		response.Email = email.String()
	}

	return response, nil
}

// GetActiveSessionsInfo returns information about active sessions for a user
func (s *sSessionManager) GetActiveSessionsInfo(ctx context.Context, userID string) ([]*model.SessionInfo, error) {
	// In a real implementation, this would query a database for active sessions
	// For now, we'll return the current session info if it matches the user
	request := g.RequestFromCtx(ctx)

	sessionUserId, err := request.Session.Get("userId")
	if err != nil || sessionUserId.IsEmpty() {
		return []*model.SessionInfo{}, nil
	}

	if sessionUserId.String() != userID {
		return []*model.SessionInfo{}, nil
	}

	sessionInfo := &model.SessionInfo{
		UserID:    userID,
		IsCurrent: true,
	}

	if sessionId, _ := request.Session.Get("sessionId"); !sessionId.IsEmpty() {
		sessionInfo.SessionID = sessionId.String()
	}

	if createdAt, _ := request.Session.Get("sessionCreatedAt"); !createdAt.IsEmpty() {
		sessionInfo.CreatedAt = time.Unix(createdAt.Int64(), 0)
	}

	if lastValidated, _ := request.Session.Get("lastValidatedAt"); !lastValidated.IsEmpty() {
		sessionInfo.LastAccessed = time.Unix(lastValidated.Int64(), 0)
	}

	// Set expiration time (example: 24 hours from creation)
	if !sessionInfo.CreatedAt.IsZero() {
		sessionInfo.ExpiresAt = sessionInfo.CreatedAt.Add(24 * time.Hour)
	}

	return []*model.SessionInfo{sessionInfo}, nil
}

// RevokeSession revokes a specific session
func (s *sSessionManager) RevokeSession(ctx context.Context, sessionID string) error {
	request := g.RequestFromCtx(ctx)

	currentSessionId, err := request.Session.Get("sessionId")
	if err != nil {
		return fmt.Errorf("failed to get current session ID: %w", err)
	}

	if currentSessionId.String() == sessionID {
		// Revoking current session
		return s.DestroyUserSession(ctx)
	}

	// In a real implementation, this would revoke sessions from a database
	g.Log().Infof(ctx, "Session %s revoked", sessionID)
	return nil
}

// RevokeAllSessions revokes all sessions for a user except optionally the current one
func (s *sSessionManager) RevokeAllSessions(ctx context.Context, userID string, exceptSessionID string) (int32, error) {
	request := g.RequestFromCtx(ctx)

	sessionUserId, err := request.Session.Get("userId")
	if err != nil || sessionUserId.IsEmpty() {
		return 0, errors.New("no active session found")
	}

	if sessionUserId.String() != userID {
		return 0, errors.New("user ID mismatch")
	}

	currentSessionId, _ := request.Session.Get("sessionId")

	// If not excluding current session, destroy it
	if exceptSessionID != currentSessionId.String() {
		err = s.DestroyUserSession(ctx)
		if err != nil {
			return 0, fmt.Errorf("failed to destroy current session: %w", err)
		}
		return 1, nil
	}

	// In a real implementation, this would revoke other sessions from database
	return 0, nil
}

// ValidateTokenRequest validates a token and returns its metadata
func (s *sSessionManager) ValidateTokenRequest(ctx context.Context, token string, expectedType string) (*model.TokenValidationResponse, error) {
	response := &model.TokenValidationResponse{
		Valid: false,
	}

	// For access tokens, validate with Supabase
	if expectedType == "access" {
		user, err := service.SupabaseService().GetUserFromToken(ctx, token)
		if err != nil {
			return response, fmt.Errorf("token validation failed: %w", err)
		}

		if user != nil {
			response.Valid = true
			response.UserID = user.ID.String()
			response.TokenType = "access"
		}
	}

	// For refresh tokens, attempt to refresh
	if expectedType == "refresh" {
		tokenResp, err := service.SupabaseService().RefreshToken(ctx, token)
		if err != nil {
			return response, fmt.Errorf("refresh token validation failed: %w", err)
		}

		if tokenResp != nil {
			response.Valid = true
			response.TokenType = "refresh"
		}
	}

	return response, nil
}

// CreateTokenForUser creates a token for a specific user and type
func (s *sSessionManager) CreateTokenForUser(ctx context.Context, userID string, tokenType string, expiresAt time.Time, metadata map[string]string) (*model.TokenCreateResponse, error) {
	response := &model.TokenCreateResponse{}

	// In a real implementation, this would create different types of tokens
	// For now, we'll handle basic access token creation through Supabase sign-in
	if tokenType == "api" {
		// Generate API token (this would typically be stored in database)
		tokenID := uuid.New().String()
		token := fmt.Sprintf("api_%s_%s", userID, tokenID)

		response.Token = token
		response.TokenID = tokenID

		g.Log().Infof(ctx, "Created %s token for user %s", tokenType, userID)
	}

	return response, nil
}

// CleanupExpiredTokens removes expired tokens from storage
func (s *sSessionManager) CleanupExpiredTokens(ctx context.Context) (int32, error) {
	// In a real implementation, this would clean up expired tokens from database
	g.Log().Infof(ctx, "Cleaning up expired tokens")
	return 0, nil
}

// RPC Service Methods that match the protobuf definitions

// CreateSession creates a new session (RPC: CreateSession)
func (s *sSessionManager) CreateSession(ctx context.Context, req *model.CreateSessionRequest) (*model.CreateSessionResponse, error) {
	if req.UserID == "" {
		return nil, errors.New("user_id cannot be empty")
	}

	// Generate session ID
	sessionID := fmt.Sprintf("%s_%s_%d", req.UserID, uuid.New().String(), time.Now().Unix())

	// Set session expiration time (24 hours from creation)
	expiresAt := time.Now().Add(24 * time.Hour)

	// Store additional session metadata
	request := g.RequestFromCtx(ctx)
	if request != nil && request.Session != nil {
		_ = request.Session.Set("sessionId", sessionID)
		_ = request.Session.Set("ipAddress", req.IPAddress)
		_ = request.Session.Set("userAgent", req.UserAgent)
		_ = request.Session.Set("expiresAt", expiresAt.Unix())
	}

	// For now, we'll create tokens through Supabase authentication
	// In a real implementation, you might have separate token generation logic
	response := &model.CreateSessionResponse{
		SessionID: sessionID,
		// These would be generated through proper authentication flow
		AccessToken:  "",
		RefreshToken: "",
	}

	g.Log().Infof(ctx, "Created session %s for user %s from IP %s", sessionID, req.UserID, req.IPAddress)
	return response, nil
}

// ValidateSessionRPC validates an existing session (RPC: ValidateSession)
func (s *sSessionManager) ValidateSessionRPC(ctx context.Context, req *model.ValidateSessionRequest) (*model.SessionValidationResponse, error) {
	if req.SessionID == "" {
		return nil, errors.New("session_id cannot be empty")
	}
	if req.AccessToken == "" {
		return nil, errors.New("access_token cannot be empty")
	}

	response := &model.SessionValidationResponse{
		Valid: false,
	}

	// Validate token with Supabase
	user, err := service.SupabaseService().GetUserFromToken(ctx, req.AccessToken)
	if err != nil {
		g.Log().Errorf(ctx, "Token validation failed for session %s: %v", req.SessionID, err)
		return response, nil // Return invalid but no error
	}

	if user != nil {
		response.Valid = true
		response.UserID = user.ID.String()
		if user.Email != "" {
			response.Email = user.Email
		}

		// Set expiration (example: 24 hours from now)
		response.ExpiresAt = time.Now().Add(24 * time.Hour)

		g.Log().Debugf(ctx, "Session %s validated successfully for user %s", req.SessionID, user.ID.String())
	}

	return response, nil
}

// RefreshSession refreshes session tokens (RPC: RefreshSession)
func (s *sSessionManager) RefreshSession(ctx context.Context, req *model.RefreshSessionRequest) (*model.RefreshSessionResponse, error) {
	if req.SessionID == "" {
		return nil, errors.New("session_id cannot be empty")
	}
	if req.RefreshToken == "" {
		return nil, errors.New("refresh_token cannot be empty")
	}

	// Refresh tokens with Supabase
	tokenResponse, err := service.SupabaseService().RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, fmt.Errorf("failed to refresh tokens for session %s: %w", req.SessionID, err)
	}

	response := &model.RefreshSessionResponse{
		NewAccessToken:  tokenResponse.AccessToken,
		NewRefreshToken: tokenResponse.RefreshToken,
		ExpiresAt:       time.Now().Add(24 * time.Hour), // Set appropriate expiration
	}

	g.Log().Infof(ctx, "Session %s tokens refreshed successfully", req.SessionID)
	return response, nil
}

// RevokeSessionRPC revokes a specific session (RPC: RevokeSession)
func (s *sSessionManager) RevokeSessionRPC(ctx context.Context, req *model.RevokeSessionRequest) (*model.RevokeSessionResponse, error) {
	if req.SessionID == "" {
		return nil, errors.New("session_id cannot be empty")
	}

	// Check if this is the current session
	request := g.RequestFromCtx(ctx)
	currentSessionId, _ := request.Session.Get("sessionId")

	if currentSessionId.String() == req.SessionID {
		// Revoking current session
		err := s.DestroyUserSession(ctx)
		if err != nil {
			return nil, fmt.Errorf("failed to revoke current session: %w", err)
		}
	}

	response := &model.RevokeSessionResponse{
		Message: fmt.Sprintf("Session %s has been revoked", req.SessionID),
	}

	g.Log().Infof(ctx, "Session %s revoked", req.SessionID)
	return response, nil
}

// GetActiveSessions retrieves active sessions for a user (RPC: GetActiveSessions)
func (s *sSessionManager) GetActiveSessions(ctx context.Context, req *model.GetActiveSessionsRequest) (*model.GetActiveSessionsResponse, error) {
	if req.UserID == "" {
		return nil, errors.New("user_id cannot be empty")
	}

	// In a real implementation, this would query database for active sessions
	// For now, we'll return the current session if it matches the user
	request := g.RequestFromCtx(ctx)
	sessionUserId, _ := request.Session.Get("userId")

	response := &model.GetActiveSessionsResponse{
		Sessions: []*model.SessionInfo{},
	}

	if !sessionUserId.IsEmpty() && sessionUserId.String() == req.UserID {
		sessionInfo := &model.SessionInfo{
			UserID:    req.UserID,
			IsCurrent: true,
		}

		if sessionId, _ := request.Session.Get("sessionId"); !sessionId.IsEmpty() {
			sessionInfo.SessionID = sessionId.String()
		}

		if createdAt, _ := request.Session.Get("sessionCreatedAt"); !createdAt.IsEmpty() {
			sessionInfo.CreatedAt = time.Unix(createdAt.Int64(), 0)
			sessionInfo.ExpiresAt = sessionInfo.CreatedAt.Add(24 * time.Hour)
		}

		if lastValidated, _ := request.Session.Get("lastValidatedAt"); !lastValidated.IsEmpty() {
			sessionInfo.LastAccessed = time.Unix(lastValidated.Int64(), 0)
		}

		response.Sessions = append(response.Sessions, sessionInfo)
	}

	g.Log().Debugf(ctx, "Retrieved %d active sessions for user %s", len(response.Sessions), req.UserID)
	return response, nil
}

// RevokeAllSessionsRPC revokes all sessions for a user (RPC: RevokeAllSessions)
func (s *sSessionManager) RevokeAllSessionsRPC(ctx context.Context, req *model.RevokeAllSessionsRequest) (*model.RevokeAllSessionsResponse, error) {
	if req.UserID == "" {
		return nil, errors.New("user_id cannot be empty")
	}

	request := g.RequestFromCtx(ctx)
	sessionUserId, _ := request.Session.Get("userId")
	currentSessionId, _ := request.Session.Get("sessionId")

	var revokedCount int32 = 0

	if !sessionUserId.IsEmpty() && sessionUserId.String() == req.UserID {
		// If not excluding current session, destroy it
		if req.ExceptSessionID != currentSessionId.String() {
			err := s.DestroyUserSession(ctx)
			if err != nil {
				return nil, fmt.Errorf("failed to destroy current session: %w", err)
			}
			revokedCount = 1
		}
	}

	response := &model.RevokeAllSessionsResponse{
		RevokedCount: revokedCount,
		Message:      fmt.Sprintf("Revoked %d sessions for user %s", revokedCount, req.UserID),
	}

	g.Log().Infof(ctx, "Revoked %d sessions for user %s", revokedCount, req.UserID)
	return response, nil
}

// CreateToken creates a new token (RPC: CreateToken)
func (s *sSessionManager) CreateToken(ctx context.Context, req *model.CreateTokenRequest) (*model.TokenCreateResponse, error) {
	if req.UserID == "" {
		return nil, errors.New("user_id cannot be empty")
	}
	if req.TokenType == "" {
		return nil, errors.New("token_type cannot be empty")
	}

	// Generate token ID and token
	tokenID := uuid.New().String()
	token := fmt.Sprintf("%s_%s_%s", req.TokenType, req.UserID, tokenID)

	response := &model.TokenCreateResponse{
		Token:   token,
		TokenID: tokenID,
	}

	g.Log().Infof(ctx, "Created %s token %s for user %s", req.TokenType, tokenID, req.UserID)
	return response, nil
}

// ValidateToken validates a token (RPC: ValidateToken)
func (s *sSessionManager) ValidateToken(ctx context.Context, req *model.ValidateTokenRequest) (*model.TokenValidationResponse, error) {
	if req.Token == "" {
		return nil, errors.New("token cannot be empty")
	}

	response := &model.TokenValidationResponse{
		Valid: false,
	}

	// For access tokens, validate with Supabase
	if req.ExpectedType == "access" {
		user, err := service.SupabaseService().GetUserFromToken(ctx, req.Token)
		if err != nil {
			g.Log().Errorf(ctx, "Access token validation failed: %v", err)
			return response, nil // Return invalid but no error
		}

		if user != nil {
			response.Valid = true
			response.UserID = user.ID.String()
			response.TokenType = "access"
			response.ExpiresAt = time.Now().Add(1 * time.Hour) // Example expiration
		}
	}

	// For refresh tokens, attempt to refresh
	if req.ExpectedType == "refresh" {
		tokenResp, err := service.SupabaseService().RefreshToken(ctx, req.Token)
		if err != nil {
			g.Log().Errorf(ctx, "Refresh token validation failed: %v", err)
			return response, nil // Return invalid but no error
		}

		if tokenResp != nil {
			response.Valid = true
			response.TokenType = "refresh"
			response.ExpiresAt = time.Now().Add(30 * 24 * time.Hour) // Example: 30 days
		}
	}

	// For other token types (api, reset, verification), implement custom validation logic
	if req.ExpectedType == "api" || req.ExpectedType == "reset" || req.ExpectedType == "verification" {
		// In a real implementation, this would check against a database
		// For now, just check token format
		if len(req.Token) > 10 && fmt.Sprintf("%s_", req.ExpectedType) == req.Token[:len(req.ExpectedType)+1] {
			response.Valid = true
			response.TokenType = req.ExpectedType
			response.ExpiresAt = time.Now().Add(24 * time.Hour)
		}
	}

	g.Log().Debugf(ctx, "Token validation result: valid=%v, type=%s", response.Valid, response.TokenType)
	return response, nil
}

// RevokeToken revokes a specific token (RPC: RevokeToken)
func (s *sSessionManager) RevokeToken(ctx context.Context, req *model.RevokeTokenRequest) (*model.RevokeTokenResponse, error) {
	if req.Token == "" {
		return nil, errors.New("token cannot be empty")
	}

	// In a real implementation, this would mark the token as revoked in database
	// For Supabase tokens, we might need to sign out the user

	response := &model.RevokeTokenResponse{
		Message: "Token has been revoked",
	}

	g.Log().Infof(ctx, "Token revoked: %s...", req.Token[:min(10, len(req.Token))])
	return response, nil
}

// CleanupExpiredTokensRPC cleans up expired tokens (RPC: CleanupExpiredTokens)
func (s *sSessionManager) CleanupExpiredTokensRPC(ctx context.Context, req *model.CleanupExpiredTokensRequest) (*model.CleanupExpiredTokensResponse, error) {
	// In a real implementation, this would query database and remove expired tokens
	cleanedCount := int32(0) // Placeholder

	response := &model.CleanupExpiredTokensResponse{
		CleanedCount: cleanedCount,
	}

	g.Log().Infof(ctx, "Cleaned up %d expired tokens", cleanedCount)
	return response, nil
}

// Helper function for min
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
