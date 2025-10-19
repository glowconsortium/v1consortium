package auth

import (
	"context"
	"fmt"
	"time"
	v1 "v1consortium/api/auth/v1"
	"v1consortium/internal/consts"
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
		ExpiresAt: timestamppb.New(resp.ExpiresAt),
	}, nil
}

func (*Controller) RefreshToken(ctx context.Context, req *v1.RefreshTokenRequest) (res *v1.RefreshTokenResponse, err error) {
	accessToken, refreshToken, err := service.Auth().RefreshToken(ctx, req.RefreshToken)
	if err != nil {
		return nil, err
	}

	return &v1.RefreshTokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpiresAt:    timestamppb.New(gtime.Now().Add(time.Hour).Time),
	}, nil
}

func (*Controller) Logout(ctx context.Context, req *v1.LogoutRequest) (res *v1.LogoutResponse, err error) {
	err = service.Auth().Logout(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}

	return &v1.LogoutResponse{
		Message: "Successfully logged out",
	}, nil
}

func (*Controller) ForgotPassword(ctx context.Context, req *v1.ForgotPasswordRequest) (res *v1.ForgotPasswordResponse, err error) {
	err = service.Auth().ForgotPassword(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	return &v1.ForgotPasswordResponse{
		Message: "Password reset email sent successfully",
	}, nil
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
	//session := service.SessionManager().GetSessionInfo(ctx)

	// if session == nil || session["access_token"] == nil {
	// 	return nil, gerror.NewCode(gcode.CodeNotAuthorized, "no valid session found")
	// }
	//resp, err := service.Auth().GetUserInfo(ctx, session["access_token"].(string))

	user, err := service.BizCtx().GetSupabaseUser(ctx)
	if err != nil {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "no valid user info found in context")
	}

	resp, err := service.Auth().GetUserProfileByEmail(ctx, user.User.Email)
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

func (*Controller) Signup(ctx context.Context, req *v1.SignupRequest) (res *v1.SignupResponse, err error) {
	// Check if user already exists
	user, err := service.Auth().GetUserProfileByEmail(ctx, req.Email)
	if err == nil && user != nil {
		return nil, gerror.NewCode(gcode.CodeBusinessValidationFailed, "user with this email already exists")
	}

	// Prepare workflow input data
	workflowInput := map[string]interface{}{
		"email":             req.Email,
		"password":          req.Password,
		"first_name":        req.FirstName,
		"last_name":         req.LastName,
		"organization_name": req.CompanyName,
		"is_dot_company":    req.IsDotCompany,
		"dot_number":        req.DotNumber,
		"role":              string(consts.RoleClientAdmin),
	}

	// Start the user signup workflow using the bridge (includes deduplication)
	workflowId, err := service.WorkflowBridge().StartUserSignupWorkflow(ctx, workflowInput, "", "")
	if err != nil {
		g.Log().Errorf(ctx, "Failed to start user signup workflow: %v", err)
		return nil, fmt.Errorf("failed to start signup workflow: %w", err)
	}

	g.Log().Info(ctx, "Started user signup workflow", g.Map{
		"workflow_id": workflowId,
		"email":       req.Email,
	})

	return &v1.SignupResponse{
		WorkflowId:                workflowId,
		Message:                   "user signup initiated successfully",
		RequiresEmailVerification: true,
	}, nil
}

func (*Controller) SocialSignup(ctx context.Context, req *v1.SocialSignupRequest) (res *v1.SocialSignupResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) CompleteRegistration(ctx context.Context, req *v1.CompleteRegistrationRequest) (res *v1.CompleteRegistrationResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}

func (*Controller) GetSignupStatus(ctx context.Context, req *v1.GetSignupStatusRequest) (res *v1.GetSignupStatusResponse, err error) {
	execution, err := service.WorkflowBridge().GetWorkflowStatus(ctx, req.WorkflowId)
	if err != nil {
		return nil, err
	}

	var emailVerified bool
	var subscriptionCompleted bool
	var email string
	var createdAt time.Time
	var status string

	if execution != nil {
		status = execution.Status
		// Safely handle execution.Context which is expected to be map[string]interface{}
		if ctxMap := execution.Context; ctxMap != nil {
			if v, ok := ctxMap["email_verified"]; ok && v != nil {
				switch t := v.(type) {
				case bool:
					emailVerified = t
				case *bool:
					if t != nil {
						emailVerified = *t
					}
				case string:
					emailVerified = (t == "true")
				}
			}
			if v, ok := ctxMap["subscription_completed"]; ok && v != nil {
				switch t := v.(type) {
				case bool:
					subscriptionCompleted = t
				case *bool:
					if t != nil {
						subscriptionCompleted = *t
					}
				case string:
					subscriptionCompleted = (t == "true")
				}
			}
			if v, ok := ctxMap["email"]; ok && v != nil {
				if s, ok := v.(string); ok {
					email = s
				}
			}
			if v, ok := ctxMap["created_at"]; ok && v != nil {
				switch t := v.(type) {
				case time.Time:
					createdAt = t
				case *time.Time:
					if t != nil {
						createdAt = *t
					}
				case string:
					if tt, err := time.Parse(time.RFC3339, t); err == nil {
						createdAt = tt
					}
				}
			}
		}
		// If not found in context, fallback to execution.CreatedAt if available
		if createdAt.IsZero() && !execution.CreatedAt.IsZero() {
			createdAt = execution.CreatedAt
		}
	}

	return &v1.GetSignupStatusResponse{
		WorkflowId:            req.WorkflowId,
		Status:                status,
		EmailVerified:         emailVerified,
		SubscriptionCompleted: subscriptionCompleted,
		Email:                 email,
		CreatedAt:             timestamppb.New(createdAt),
	}, nil
}

func (*Controller) ResendVerification(ctx context.Context, req *v1.ResendVerificationRequest) (res *v1.ResendVerificationResponse, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
