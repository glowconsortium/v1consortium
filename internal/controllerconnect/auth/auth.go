package authconnect

import (
	"context"
	"errors"
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

// HandleError converts GoFrame errors to Connect errors with appropriate error codes
func HandleError(err error) error {
	if err == nil {
		return nil
	}

	// Check if it's already a GoFrame error with a code
	if gErr := gerror.Cause(err); gErr != nil {
		if code := gerror.Code(gErr); code != gcode.CodeNil {
			return mapGoFrameToConnectError(code, gErr.Error())
		}
	}

	// Default to internal error for unknown errors
	return connect.NewError(connect.CodeInternal, err)
}

// mapGoFrameToConnectError maps GoFrame error codes to Connect error codes
func mapGoFrameToConnectError(gfCode gcode.Code, message string) error {
	var connectCode connect.Code

	switch gfCode {
	// Success cases
	case gcode.CodeOK:
		return nil

	// Client errors (4xx equivalent)
	case gcode.CodeInvalidParameter, gcode.CodeMissingParameter, gcode.CodeInvalidRequest:
		connectCode = connect.CodeInvalidArgument
	case gcode.CodeValidationFailed, gcode.CodeBusinessValidationFailed:
		connectCode = connect.CodeInvalidArgument
	case gcode.CodeNotAuthorized:
		connectCode = connect.CodeUnauthenticated
	case gcode.CodeNotFound:
		connectCode = connect.CodeNotFound
	case gcode.CodeInvalidOperation:
		connectCode = connect.CodeFailedPrecondition
	case gcode.CodeNotSupported:
		connectCode = connect.CodeUnimplemented
	case gcode.CodeNotImplemented:
		connectCode = connect.CodeUnimplemented

	// Server errors (5xx equivalent)
	case gcode.CodeInternalError, gcode.CodeInternalPanic:
		connectCode = connect.CodeInternal
	case gcode.CodeDbOperationError, gcode.CodeOperationFailed:
		connectCode = connect.CodeInternal
	case gcode.CodeServerBusy:
		connectCode = connect.CodeUnavailable
	case gcode.CodeInvalidConfiguration, gcode.CodeMissingConfiguration:
		connectCode = connect.CodeInternal
	case gcode.CodeNecessaryPackageNotImport:
		connectCode = connect.CodeInternal

	// Fallback for unknown codes
	case gcode.CodeUnknown:
		connectCode = connect.CodeUnknown
	default:
		connectCode = connect.CodeInternal
	}

	return connect.NewError(connectCode, errors.New(message))
}

func (s *AuthConnectService) Login(ctx context.Context, req *connect.Request[v1.LoginRequest]) (res *connect.Response[v1.LoginResponse], err error) {
	resp, err := s.authController.Login(ctx, req.Msg)
	if err != nil {
		return nil, HandleError(err)
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

func (s *AuthConnectService) RefreshToken(ctx context.Context, req *connect.Request[v1.RefreshTokenRequest]) (res *connect.Response[v1.RefreshTokenResponse], err error) {
	resp, err := s.authController.RefreshToken(ctx, req.Msg)
	if err != nil {
		return nil, HandleError(err)
	}
	return connect.NewResponse(resp), nil
}

func (s *AuthConnectService) Logout(ctx context.Context, req *connect.Request[v1.LogoutRequest]) (res *connect.Response[v1.LogoutResponse], err error) {
	resp, err := s.authController.Logout(ctx, req.Msg)
	if err != nil {
		return nil, HandleError(err)
	}
	return connect.NewResponse(resp), nil
}

func (s *AuthConnectService) ForgotPassword(ctx context.Context, req *connect.Request[v1.ForgotPasswordRequest]) (res *connect.Response[v1.ForgotPasswordResponse], err error) {
	resp, err := s.authController.ForgotPassword(ctx, req.Msg)
	if err != nil {
		return nil, HandleError(err)
	}
	return connect.NewResponse(resp), nil
}

func (s *AuthConnectService) ResetPassword(ctx context.Context, req *connect.Request[v1.ResetPasswordRequest]) (res *connect.Response[v1.ResetPasswordResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("reset password not implemented"))
}

func (s *AuthConnectService) ChangePassword(ctx context.Context, req *connect.Request[v1.ChangePasswordRequest]) (res *connect.Response[v1.ChangePasswordResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("change password not implemented"))
}

func (s *AuthConnectService) VerifyEmail(ctx context.Context, req *connect.Request[v1.VerifyEmailRequest]) (res *connect.Response[v1.VerifyEmailResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("verify email not implemented"))
}

func (s *AuthConnectService) EnableMFA(ctx context.Context, req *connect.Request[v1.EnableMFARequest]) (res *connect.Response[v1.EnableMFAResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("enable MFA not implemented"))
}

func (s *AuthConnectService) VerifyMFA(ctx context.Context, req *connect.Request[v1.VerifyMFARequest]) (res *connect.Response[v1.VerifyMFAResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("verify MFA not implemented"))
}

func (s *AuthConnectService) DisableMFA(ctx context.Context, req *connect.Request[v1.DisableMFARequest]) (res *connect.Response[v1.DisableMFAResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("disable MFA not implemented"))
}

func (s *AuthConnectService) GetUser(ctx context.Context, req *connect.Request[v1.GetUserRequest]) (res *connect.Response[v1.GetUserResponse], err error) {
	resp, err := s.authController.GetUser(ctx, req.Msg)
	if err != nil {
		return nil, HandleError(err)
	}
	return connect.NewResponse(resp), nil
}

func (s *AuthConnectService) UpdateUser(ctx context.Context, req *connect.Request[v1.UpdateUserRequest]) (res *connect.Response[v1.UpdateUserResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("update user not implemented"))
}

func (s *AuthConnectService) CheckPermission(ctx context.Context, req *connect.Request[v1.CheckPermissionRequest]) (res *connect.Response[v1.CheckPermissionResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("check permission not implemented"))
}

func (s *AuthConnectService) GetUserPermissions(ctx context.Context, req *connect.Request[v1.GetUserPermissionsRequest]) (res *connect.Response[v1.GetUserPermissionsResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("get user permissions not implemented"))
}

func (s *AuthConnectService) HasPermission(ctx context.Context, req *connect.Request[v1.HasPermissionRequest]) (res *connect.Response[v1.HasPermissionResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("has permission not implemented"))
}

func (s *AuthConnectService) AssignRole(ctx context.Context, req *connect.Request[v1.AssignRoleRequest]) (res *connect.Response[v1.AssignRoleResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("assign role not implemented"))
}

func (s *AuthConnectService) RemoveRole(ctx context.Context, req *connect.Request[v1.RemoveRoleRequest]) (res *connect.Response[v1.RemoveRoleResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("remove role not implemented"))
}

func (s *AuthConnectService) GetUserRoles(ctx context.Context, req *connect.Request[v1.GetUserRolesRequest]) (res *connect.Response[v1.GetUserRolesResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("get user roles not implemented"))
}

func (s *AuthConnectService) GetRolePermissions(ctx context.Context, req *connect.Request[v1.GetRolePermissionsRequest]) (res *connect.Response[v1.GetRolePermissionsResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("get role permissions not implemented"))
}

func (s *AuthConnectService) CreateSession(ctx context.Context, req *connect.Request[v1.CreateSessionRequest]) (res *connect.Response[v1.CreateSessionResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("create session not implemented"))
}

func (s *AuthConnectService) ValidateSession(ctx context.Context, req *connect.Request[v1.ValidateSessionRequest]) (res *connect.Response[v1.ValidateSessionResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("validate session not implemented"))
}

func (s *AuthConnectService) RefreshSession(ctx context.Context, req *connect.Request[v1.RefreshSessionRequest]) (res *connect.Response[v1.RefreshSessionResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("refresh session not implemented"))
}

func (s *AuthConnectService) RevokeSession(ctx context.Context, req *connect.Request[v1.RevokeSessionRequest]) (res *connect.Response[v1.RevokeSessionResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("revoke session not implemented"))
}

func (s *AuthConnectService) GetActiveSessions(ctx context.Context, req *connect.Request[v1.GetActiveSessionsRequest]) (res *connect.Response[v1.GetActiveSessionsResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("get active sessions not implemented"))
}

func (s *AuthConnectService) RevokeAllSessions(ctx context.Context, req *connect.Request[v1.RevokeAllSessionsRequest]) (res *connect.Response[v1.RevokeAllSessionsResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("revoke all sessions not implemented"))
}

func (s *AuthConnectService) CreateToken(ctx context.Context, req *connect.Request[v1.CreateTokenRequest]) (res *connect.Response[v1.CreateTokenResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("create token not implemented"))
}

func (s *AuthConnectService) ValidateToken(ctx context.Context, req *connect.Request[v1.ValidateTokenRequest]) (res *connect.Response[v1.ValidateTokenResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("validate token not implemented"))
}

func (s *AuthConnectService) RevokeToken(ctx context.Context, req *connect.Request[v1.RevokeTokenRequest]) (res *connect.Response[v1.RevokeTokenResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("revoke token not implemented"))
}

func (s *AuthConnectService) CleanupExpiredTokens(ctx context.Context, req *connect.Request[v1.CleanupExpiredTokensRequest]) (res *connect.Response[v1.CleanupExpiredTokensResponse], err error) {
	return nil, connect.NewError(connect.CodeUnimplemented, errors.New("cleanup expired tokens not implemented"))
}

// CompleteRegistration handles completing a user's registration; stubbed as unimplemented
func (s *AuthConnectService) CompleteRegistration(ctx context.Context, req *connect.Request[v1.CompleteRegistrationRequest]) (res *connect.Response[v1.CompleteRegistrationResponse], err error) {
	resp, err := s.authController.CompleteRegistration(ctx, req.Msg)
	if err != nil {
		return nil, HandleError(err)
	}
	return connect.NewResponse(resp), nil
}

// Signup handles user signup/registration
func (s *AuthConnectService) Signup(ctx context.Context, req *connect.Request[v1.SignupRequest]) (res *connect.Response[v1.SignupResponse], err error) {
	resp, err := s.authController.Signup(ctx, req.Msg)
	if err != nil {
		return nil, HandleError(err)
	}
	return connect.NewResponse(resp), nil
}

// SocialSignup handles social media signup
func (s *AuthConnectService) SocialSignup(ctx context.Context, req *connect.Request[v1.SocialSignupRequest]) (res *connect.Response[v1.SocialSignupResponse], err error) {
	resp, err := s.authController.SocialSignup(ctx, req.Msg)
	if err != nil {
		return nil, HandleError(err)
	}
	return connect.NewResponse(resp), nil
}

// GetSignupStatus gets the status of a signup workflow
func (s *AuthConnectService) GetSignupStatus(ctx context.Context, req *connect.Request[v1.GetSignupStatusRequest]) (res *connect.Response[v1.GetSignupStatusResponse], err error) {
	resp, err := s.authController.GetSignupStatus(ctx, req.Msg)
	if err != nil {
		return nil, HandleError(err)
	}
	return connect.NewResponse(resp), nil
}

// ResendVerification resends verification email
func (s *AuthConnectService) ResendVerification(ctx context.Context, req *connect.Request[v1.ResendVerificationRequest]) (res *connect.Response[v1.ResendVerificationResponse], err error) {
	resp, err := s.authController.ResendVerification(ctx, req.Msg)
	if err != nil {
		return nil, HandleError(err)
	}
	return connect.NewResponse(resp), nil
}
