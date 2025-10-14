package authconnect

import (
	"context"
	"fmt"
	v1 "v1consortium/api/auth/v1"
	"v1consortium/internal/controller/auth"

	"connectrpc.com/connect"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
)

type AuthConnectService struct {
	authController *auth.Controller
}

func NewAuthConnectService(ctx context.Context) *AuthConnectService {
	return &AuthConnectService{
		authController: &auth.Controller{},
	}
}

// type Controller struct {
// 	v1.UnimplementedAuthServiceServer
// }

// func Register(s *grpcx.GrpcServer) {
// 	v1.RegisterAuthServiceServer(s.Server, &Controller{})
// }

func (s *AuthConnectService) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (res *connect.Response[v1.LoginResponse], err error) {
	resp, err := s.authController.Login(ctx, req.Msg)
	if err != nil {
		return nil, err
	}

	// Create response
	response := connect.NewResponse(resp)

	// Set cookies directly in the response for login
	// Use access token as session identifier for the cookie
	if resp.SessionId != "" {
		sessionCookie := buildSessionCookie("session_id", resp.SessionId)
		response.Header().Add("Set-Cookie", sessionCookie)
	}

	if resp.User != nil && resp.User.OrganizationId != "" {
		orgCookie := buildSessionCookie("organization_id", resp.User.OrganizationId)
		response.Header().Add("Set-Cookie", orgCookie)
	}

	// Also set temporary headers for the interceptor to pick up (as backup)
	response.Header().Set("X-Session-ID", resp.SessionId)
	if resp.User != nil {
		response.Header().Set("X-Organization-ID", resp.User.OrganizationId)
	}

	return response, nil
}

// buildSessionCookie creates a session cookie string
func buildSessionCookie(name, value string) string {
	// Build cookie with secure defaults
	cookie := fmt.Sprintf("%s=%s; Path=/; HttpOnly; Secure; SameSite=Strict; Max-Age=86400", name, value)
	return cookie
}

func (s *AuthConnectService) RegisterUser(ctx context.Context, req *connect.Request[v1.RegisterRequest]) (res *connect.Response[v1.RegisterResponse], err error) {
	resp, err := s.authController.RegisterUser(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *AuthConnectService) RefreshToken(ctx context.Context, req *connect.Request[v1.RefreshTokenRequest]) (res *connect.Response[v1.RefreshTokenResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) Logout(ctx context.Context, req *connect.Request[v1.LogoutRequest]) (res *connect.Response[v1.LogoutResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) ForgotPassword(ctx context.Context, req *connect.Request[v1.ForgotPasswordRequest]) (res *connect.Response[v1.ForgotPasswordResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) ResetPassword(ctx context.Context, req *connect.Request[v1.ResetPasswordRequest]) (res *connect.Response[v1.ResetPasswordResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) ChangePassword(ctx context.Context, req *connect.Request[v1.ChangePasswordRequest]) (res *connect.Response[v1.ChangePasswordResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) VerifyEmail(ctx context.Context, req *connect.Request[v1.VerifyEmailRequest]) (res *connect.Response[v1.VerifyEmailResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) EnableMFA(ctx context.Context, req *connect.Request[v1.EnableMFARequest]) (res *connect.Response[v1.EnableMFAResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) VerifyMFA(ctx context.Context, req *connect.Request[v1.VerifyMFARequest]) (res *connect.Response[v1.VerifyMFAResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) DisableMFA(ctx context.Context, req *connect.Request[v1.DisableMFARequest]) (res *connect.Response[v1.DisableMFAResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) GetUser(ctx context.Context, req *connect.Request[v1.GetUserRequest]) (res *connect.Response[v1.GetUserResponse], err error) {
	resp, err := s.authController.GetUser(ctx, req.Msg)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(resp), nil
}

func (s *AuthConnectService) UpdateUser(ctx context.Context, req *connect.Request[v1.UpdateUserRequest]) (res *connect.Response[v1.UpdateUserResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) CheckPermission(ctx context.Context, req *connect.Request[v1.CheckPermissionRequest]) (res *connect.Response[v1.CheckPermissionResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) GetUserPermissions(ctx context.Context, req *connect.Request[v1.GetUserPermissionsRequest]) (res *connect.Response[v1.GetUserPermissionsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) HasPermission(ctx context.Context, req *connect.Request[v1.HasPermissionRequest]) (res *connect.Response[v1.HasPermissionResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) AssignRole(ctx context.Context, req *connect.Request[v1.AssignRoleRequest]) (res *connect.Response[v1.AssignRoleResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) RemoveRole(ctx context.Context, req *connect.Request[v1.RemoveRoleRequest]) (res *connect.Response[v1.RemoveRoleResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) GetUserRoles(ctx context.Context, req *connect.Request[v1.GetUserRolesRequest]) (res *connect.Response[v1.GetUserRolesResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) GetRolePermissions(ctx context.Context, req *connect.Request[v1.GetRolePermissionsRequest]) (res *connect.Response[v1.GetRolePermissionsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) CreateSession(ctx context.Context, req *connect.Request[v1.CreateSessionRequest]) (res *connect.Response[v1.CreateSessionResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) ValidateSession(ctx context.Context, req *connect.Request[v1.ValidateSessionRequest]) (res *connect.Response[v1.ValidateSessionResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) RefreshSession(ctx context.Context, req *connect.Request[v1.RefreshSessionRequest]) (res *connect.Response[v1.RefreshSessionResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) RevokeSession(ctx context.Context, req *connect.Request[v1.RevokeSessionRequest]) (res *connect.Response[v1.RevokeSessionResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) GetActiveSessions(ctx context.Context, req *connect.Request[v1.GetActiveSessionsRequest]) (res *connect.Response[v1.GetActiveSessionsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) RevokeAllSessions(ctx context.Context, req *connect.Request[v1.RevokeAllSessionsRequest]) (res *connect.Response[v1.RevokeAllSessionsResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) CreateToken(ctx context.Context, req *connect.Request[v1.CreateTokenRequest]) (res *connect.Response[v1.CreateTokenResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) ValidateToken(ctx context.Context, req *connect.Request[v1.ValidateTokenRequest]) (res *connect.Response[v1.ValidateTokenResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) RevokeToken(ctx context.Context, req *connect.Request[v1.RevokeTokenRequest]) (res *connect.Response[v1.RevokeTokenResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (s *AuthConnectService) CleanupExpiredTokens(ctx context.Context, req *connect.Request[v1.CleanupExpiredTokensRequest]) (res *connect.Response[v1.CleanupExpiredTokensResponse], err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
