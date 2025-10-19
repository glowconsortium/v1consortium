package auth

import (
	"context"
	"database/sql"
	"time"
	"v1consortium/internal/dao"
	"v1consortium/internal/model"
	"v1consortium/internal/model/do"
	"v1consortium/internal/model/entity"
	"v1consortium/internal/service"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sAuth struct{}

func new() service.IAuth {
	return &sAuth{}
}

func init() {
	service.RegisterAuth(new())
}

// Add your methods for sAuth here

func (s *sAuth) Login(ctx context.Context, email, password, ipaddress, useragent string) (*model.LoginResponse, error) {

	resp, err := service.SupabaseService().SignIn(ctx, email, password)
	if err != nil {
		return nil, err
	}

	userssession, err := service.SessionManager().CreateSession(ctx, &model.CreateSessionRequest{
		UserID:    resp.User.ID.String(),
		IPAddress: ipaddress,
		UserAgent: useragent,
	})

	userssession.AccessToken = resp.AccessToken
	userssession.RefreshToken = resp.RefreshToken

	userprofile, err := s.GetUserProfileByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if userprofile == nil || userprofile.Id == "" {
		return nil, gerror.NewCode(gcode.CodeNotFound, "User profile not found")
	}

	userssessionresponse := &model.LoginResponse{
		Session:   userssession,
		User:      userprofile,
		ExpiresAt: time.Unix(resp.ExpiresAt, 0),
	}

	return userssessionresponse, nil
}

func (s *sAuth) Logout(ctx context.Context, token string) error {
	err := service.SupabaseService().SignOut(ctx, token)
	return err
}

func (s *sAuth) RegisterUser(ctx context.Context, email, password string, data map[string]interface{}) (userID string, err error) {

	userprofile, err := s.GetUserProfileByEmail(ctx, email)
	if err != nil && !gerror.HasCode(err, gcode.CodeNotFound) && err != sql.ErrNoRows {
		g.Log().Errorf(ctx, "Error checking user %s: %v", email, err)
		return "", err
	}

	if userprofile != nil && userprofile.Id != "" {
		return "", gerror.NewCode(gcode.CodeInvalidParameter, "User with this email already exists")
	}

	resp, err := service.SupabaseService().SignUp(ctx, email, password, data)
	if err != nil {
		return "", err
	}

	_, err = dao.UserProfiles.Ctx(ctx).Insert(do.UserProfiles{
		Email:          email,
		Id:             resp.User.ID.String(),
		OrganizationId: data["organization_id"].(string),
		Role:           data["role"].(string),
		FirstName:      data["first_name"].(string),
		LastName:       data["last_name"].(string),
	})
	return resp.User.ID.String(), err
}

// make sure to supply id in profileData for create
func (s *sAuth) CreateUserProfile(ctx context.Context, id string, profileData *do.UserProfiles) error {

	profileData.Id = id
	_, err := dao.UserProfiles.Ctx(ctx).Data(profileData).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (s *sAuth) RefreshToken(ctx context.Context, refreshToken string) (access string, refresh string, err error) {
	resp, err := service.SupabaseService().RefreshToken(ctx, refreshToken)
	if err != nil {
		return "", "", err
	}
	return resp.AccessToken, resp.RefreshToken, nil
}

func (s *sAuth) GetUserProfileByEmail(ctx context.Context, email string) (*entity.UserProfiles, error) {
	var userProfile entity.UserProfiles

	err := dao.UserProfiles.Ctx(ctx).Where(dao.UserProfiles.Columns().Email, email).Scan(&userProfile)
	if err != nil {

		return nil, err
	}
	if userProfile.Id == "" {
		return nil, nil
	}
	return &userProfile, nil
}

func (s *sAuth) ForgotPassword(ctx context.Context, email string) error {
	// Implement your forgot password logic here
	return nil
}

func (s *sAuth) ResetPassword(ctx context.Context, token, newPassword string) error {
	// Implement your reset password logic here
	return nil
}

func (s *sAuth) ChangePassword(ctx context.Context, userID, oldPassword, newPassword string) error {
	// Implement your change password logic here
	return nil
}

func (s *sAuth) VerifyEmail(ctx context.Context, token string) error {
	// Implement your email verification logic here
	return nil
}

func (s *sAuth) EnableMFA(ctx context.Context, userID string) (string, error) {
	// Implement your enable MFA logic here
	return "", nil
}

func (s *sAuth) DisableMFA(ctx context.Context, userID string) error {
	// Implement your disable MFA logic here
	return nil
}

func (s *sAuth) VerifyMFA(ctx context.Context, userID, code string) (string, error) {
	// Implement your verify MFA logic here
	return "", nil
}

func (s *sAuth) GetUserInfo(ctx context.Context, token string) (*entity.UserProfiles, error) {
	// Implement your get user info logic here
	// userfromtoken supabase
	userInfo, err := service.SupabaseService().GetUserFromToken(ctx, token)

	if err != nil {
		return nil, err
	}
	if userInfo == nil || userInfo.User.Email == "" {
		return nil, gerror.NewCode(gcode.CodeNotAuthorized, "Invalid token or user not found")
	}

	userProfile, err := s.GetUserProfileByEmail(ctx, userInfo.User.Email)
	if err != nil {
		return nil, err
	}
	if userProfile == nil || userProfile.Id == "" {
		return nil, gerror.NewCode(gcode.CodeNotFound, "User profile not found")
	}

	return userProfile, nil
}

func (s *sAuth) UpdateUserProfile(ctx context.Context, userID string, profileData map[string]interface{}) error {
	// Implement your update user profile logic here
	return nil
}

func (s *sAuth) CheckPermission(ctx context.Context, userID, permission string) (bool, error) {
	// Implement your check permissions logic here
	return false, nil
}

func (s *sAuth) GetUserPermissions(ctx context.Context, userID string) ([]string, error) {
	// Implement your get user permissions logic here
	return nil, nil
}
