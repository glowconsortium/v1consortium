package auth

import (
	"context"
	v1 "v1consortium/api/auth/v1"
	"v1consortium/internal/service"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Controller struct {
	v1.UnimplementedAuthServiceServer
}

func Register(s *grpcx.GrpcServer) {
	v1.RegisterAuthServiceServer(s.Server, &Controller{})
}

func (*Controller) Login(ctx context.Context, req *v1.LoginRequest) (res *v1.LoginResponse, err error) {
	//return nil, gerror.NewCode(gcode.CodeNotImplemented)

	reqinct := g.RequestFromCtx(ctx)

	if reqinct == nil {
		return nil, gerror.NewCode(gcode.CodeInternalError, "request not found in context")
	}

	useragent := reqinct.UserAgent()
	ipaddress := reqinct.GetClientIp()

	resp, err := service.Auth().Login(ctx, req.Email, req.Password, ipaddress, useragent)
	if err != nil {
		return nil, err
	}

	//reqinct.Cookie.SetCookie("session_id", resp.Session.SessionID, "", "/", time.Duration(3600)*time.Second)

	// Update context with session and organization information
	ctx = service.BizCtx().SetCurrentSessionID(ctx, resp.Session.SessionID)
	ctx = service.BizCtx().SetCurrentOrganizationID(ctx, resp.User.OrganizationId)

	return &v1.LoginResponse{
		AccessToken:  resp.Session.AccessToken,
		RefreshToken: resp.Session.RefreshToken,
		User: &v1.UserSession{
			UserId:         resp.User.Id,
			Email:          resp.User.Email,
			FirstName:      resp.User.FirstName,
			LastName:       resp.User.LastName,
			OrganizationId: resp.User.OrganizationId,
			LastLogin:      timestamppb.New(gtime.Now().Time),
		},
		SessionId: resp.Session.SessionID,
	}, nil
}

func (*Controller) RegisterUser(ctx context.Context, req *v1.RegisterRequest) (res *v1.RegisterResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) RefreshToken(ctx context.Context, req *v1.RefreshTokenRequest) (res *v1.RefreshTokenResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) Logout(ctx context.Context, req *v1.LogoutRequest) (res *v1.LogoutResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ForgotPassword(ctx context.Context, req *v1.ForgotPasswordRequest) (res *v1.ForgotPasswordResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ResetPassword(ctx context.Context, req *v1.ResetPasswordRequest) (res *v1.ResetPasswordResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ChangePassword(ctx context.Context, req *v1.ChangePasswordRequest) (res *v1.ChangePasswordResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) VerifyEmail(ctx context.Context, req *v1.VerifyEmailRequest) (res *v1.VerifyEmailResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) EnableMFA(ctx context.Context, req *v1.EnableMFARequest) (res *v1.EnableMFAResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) VerifyMFA(ctx context.Context, req *v1.VerifyMFARequest) (res *v1.VerifyMFAResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) DisableMFA(ctx context.Context, req *v1.DisableMFARequest) (res *v1.DisableMFAResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetUser(ctx context.Context, req *v1.GetUserRequest) (res *v1.GetUserResponse, err error) {
	session := service.SessionManager().GetSessionInfo(ctx)
	resp, err := service.Auth().GetUserInfo(ctx, session["access_token"].(string))
	if err != nil {
		g.Log().Errorf(ctx, "Failed to get user info: %v", err)
		return nil, err
	}
	return &v1.GetUserResponse{
		User: &v1.UserSession{
			UserId:         resp.Id,
			Email:          resp.Email,
			FirstName:      resp.FirstName,
			LastName:       resp.LastName,
			OrganizationId: resp.OrganizationId,
			Role:           resp.Role,
		},
	}, nil
}

func (*Controller) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (res *v1.UpdateUserResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) CheckPermission(ctx context.Context, req *v1.CheckPermissionRequest) (res *v1.CheckPermissionResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetUserPermissions(ctx context.Context, req *v1.GetUserPermissionsRequest) (res *v1.GetUserPermissionsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) HasPermission(ctx context.Context, req *v1.HasPermissionRequest) (res *v1.HasPermissionResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) AssignRole(ctx context.Context, req *v1.AssignRoleRequest) (res *v1.AssignRoleResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) RemoveRole(ctx context.Context, req *v1.RemoveRoleRequest) (res *v1.RemoveRoleResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetUserRoles(ctx context.Context, req *v1.GetUserRolesRequest) (res *v1.GetUserRolesResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetRolePermissions(ctx context.Context, req *v1.GetRolePermissionsRequest) (res *v1.GetRolePermissionsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) CreateSession(ctx context.Context, req *v1.CreateSessionRequest) (res *v1.CreateSessionResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ValidateSession(ctx context.Context, req *v1.ValidateSessionRequest) (res *v1.ValidateSessionResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) RefreshSession(ctx context.Context, req *v1.RefreshSessionRequest) (res *v1.RefreshSessionResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) RevokeSession(ctx context.Context, req *v1.RevokeSessionRequest) (res *v1.RevokeSessionResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetActiveSessions(ctx context.Context, req *v1.GetActiveSessionsRequest) (res *v1.GetActiveSessionsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) RevokeAllSessions(ctx context.Context, req *v1.RevokeAllSessionsRequest) (res *v1.RevokeAllSessionsResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) CreateToken(ctx context.Context, req *v1.CreateTokenRequest) (res *v1.CreateTokenResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) ValidateToken(ctx context.Context, req *v1.ValidateTokenRequest) (res *v1.ValidateTokenResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) RevokeToken(ctx context.Context, req *v1.RevokeTokenRequest) (res *v1.RevokeTokenResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) CleanupExpiredTokens(ctx context.Context, req *v1.CleanupExpiredTokensRequest) (res *v1.CleanupExpiredTokensResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
