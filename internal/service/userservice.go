// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"v1consortium/internal/pkg/auth0client"
)

type (
	IUserService interface {
		GetUserInfo(ctx context.Context)
		GetAuth0Config(ctx context.Context) (auth0client.Config, error)
	}
)

var (
	localUserService IUserService
)

func UserService() IUserService {
	if localUserService == nil {
		panic("implement not found for interface IUserService, forgot register?")
	}
	return localUserService
}

func RegisterUserService(i IUserService) {
	localUserService = i
}
