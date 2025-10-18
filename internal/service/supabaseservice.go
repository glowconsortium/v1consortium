// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"v1consortium/internal/pkg/supabaseclient"

	"github.com/google/uuid"
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
)

type (
	ISupabaseService interface {
		GetAnonClient(ctx context.Context) (*supabase.Client, error)
		GetServiceClient(ctx context.Context) (*supabase.Client, error)
		GetUserAuthenticatedClient(ctx context.Context, accessToken string) (*supabase.Client, error)
		GetSupabaseConfig(ctx context.Context) *supabaseclient.SupabaseConfig
		GetUser(ctx context.Context, userID uuid.UUID) (*types.AdminGetUserResponse, error)
		CreateUser(ctx context.Context, email string, password string, userMetadata map[string]interface{}) (*types.AdminCreateUserResponse, error)
		DeleteUser(ctx context.Context, userID uuid.UUID) error
		UpdateUser(ctx context.Context, userID uuid.UUID, email *string, password *string, userMetadata map[string]interface{}) (*types.AdminUpdateUserResponse, error)
		SignIn(ctx context.Context, email string, password string) (*types.TokenResponse, error)
		SignUp(ctx context.Context, email string, password string, userMetadata map[string]interface{}) (*types.SignupResponse, error)
		// social signup
		SignUpWithProvider(ctx context.Context, provider string, redirectTo string, scopes string) (*types.AuthorizeResponse, error)
		RefreshToken(ctx context.Context, refreshToken string) (*types.TokenResponse, error)
		SignOut(ctx context.Context, accessToken string) error
		GetUserFromToken(ctx context.Context, accessToken string) (*types.UserResponse, error)
		// ValidateTokenExpiry checks if the access token is expired
		ValidateTokenExpiry(ctx context.Context, accessToken string) error
		// ValidateAndRefreshSession validates the current session and refreshes if needed
		ValidateAndRefreshSession(ctx context.Context, accessToken string, refreshToken string) (*types.TokenResponse, error)
		// ValidateSessionSecurity performs additional security checks on the session
		ValidateSessionSecurity(ctx context.Context, accessToken string, expectedUserID uuid.UUID) error
	}
)

var (
	localSupabaseService ISupabaseService
)

func SupabaseService() ISupabaseService {
	if localSupabaseService == nil {
		panic("implement not found for interface ISupabaseService, forgot register?")
	}
	return localSupabaseService
}

func RegisterSupabaseService(i ISupabaseService) {
	localSupabaseService = i
}
