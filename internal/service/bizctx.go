// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"

	"github.com/supabase-community/gotrue-go/types"
)

type (
	IBizCtx interface {
		Init(ctx context.Context) error
		SetSupabaseUser(ctx context.Context, userInfo *types.UserResponse) context.Context
		GetSupabaseUser(ctx context.Context) (*types.UserResponse, error)
		SetCurrentSessionID(ctx context.Context, sessionID string) context.Context
		SetCurrentOrganizationID(ctx context.Context, organizationID string) context.Context
		GetCurrentOrganizationID(ctx context.Context) (string, error)
		GetCurrentSessionID(ctx context.Context) (string, error)
	}
)

var (
	localBizCtx IBizCtx
)

func BizCtx() IBizCtx {
	if localBizCtx == nil {
		panic("implement not found for interface IBizCtx, forgot register?")
	}
	return localBizCtx
}

func RegisterBizCtx(i IBizCtx) {
	localBizCtx = i
}
