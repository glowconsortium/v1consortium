// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"v1consortium/internal/model"
	"v1consortium/internal/model/entity"
)

type (
	IAuth interface {
		Login(ctx context.Context, email string, password string, ipaddress string, useragent string) (*model.LoginResponse, error)
		Logout(ctx context.Context, token string) error
		RegisterUser(ctx context.Context, email string, password string, data map[string]interface{}) (userID string, err error)
		RefreshToken(ctx context.Context, refreshToken string) (access string, refresh string, err error)
		GetUserProfileByEmail(ctx context.Context, email string) (*entity.UserProfiles, error)
		ForgotPassword(ctx context.Context, email string) error
		ResetPassword(ctx context.Context, token string, newPassword string) error
		ChangePassword(ctx context.Context, userID string, oldPassword string, newPassword string) error
		VerifyEmail(ctx context.Context, token string) error
		EnableMFA(ctx context.Context, userID string) (string, error)
		DisableMFA(ctx context.Context, userID string) error
		VerifyMFA(ctx context.Context, userID string, code string) (string, error)
		GetUserInfo(ctx context.Context, token string) (*entity.UserProfiles, error)
		UpdateUserProfile(ctx context.Context, userID string, profileData map[string]interface{}) error
		CheckPermission(ctx context.Context, userID string, permission string) (bool, error)
		GetUserPermissions(ctx context.Context, userID string) ([]string, error)
	}
)

var (
	localAuth IAuth
)

func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot register?")
	}
	return localAuth
}

func RegisterAuth(i IAuth) {
	localAuth = i
}
