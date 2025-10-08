package supabaseservice

import (
	"context"
	"v1consortium/internal/pkg/supabaseclient"
	"v1consortium/internal/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/uuid"
	"github.com/supabase-community/gotrue-go/types"
	"github.com/supabase-community/supabase-go"
)

func new() service.ISupabaseService {
	return &sSupabaseService{}
}

func init() {
	service.RegisterSupabaseService(new())
}

type sSupabaseService struct{}

func (s *sSupabaseService) GetAnonClient(ctx context.Context) (*supabase.Client, error) {
	config := s.GetSupabaseConfig(ctx)

	client, err := supabaseclient.NewSupabaseClient(*config)
	if err != nil {
		return nil, err
	}
	return client.GetAnonClient(), nil
}

func (s *sSupabaseService) GetServiceClient(ctx context.Context) (*supabase.Client, error) {
	config := s.GetSupabaseConfig(ctx)

	client, err := supabaseclient.NewSupabaseClient(*config)
	if err != nil {
		return nil, err
	}
	return client.GetServiceClient(), nil
}

func (s *sSupabaseService) GetUserAuthenticatedClient(ctx context.Context, accessToken string) (*supabase.Client, error) {

	anonclient, err := s.GetAnonClient(ctx)
	if err != nil {
		return nil, err
	}
	anonclient.UpdateAuthSession(types.Session{
		AccessToken: accessToken,
	})

	return anonclient, nil
}

func (s *sSupabaseService) GetSupabaseConfig(ctx context.Context) *supabaseclient.SupabaseConfig {

	Url := g.Cfg().MustGet(ctx, "supabase.url").String()
	PublicApiKey := g.Cfg().MustGet(ctx, "supabase.publicApiKey").String()
	SecretApiKey := g.Cfg().MustGet(ctx, "supabase.secretApiKey").String()

	g.Log().Debugf(ctx, "Supabase Config - URL: %s, PublicApiKey: %s, SecretApiKey: %s", Url, PublicApiKey, SecretApiKey)

	return &supabaseclient.SupabaseConfig{
		Url:          Url,
		PublicApiKey: PublicApiKey,
		SecretApiKey: SecretApiKey,
	}
}

func (s *sSupabaseService) GetUser(ctx context.Context, userID uuid.UUID) (*types.AdminGetUserResponse, error) {
	client, err := s.GetServiceClient(ctx)
	if err != nil {
		return nil, err
	}
	user, err := client.Auth.AdminGetUser(types.AdminGetUserRequest{
		UserID: userID,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *sSupabaseService) CreateUser(ctx context.Context, email, password string, userMetadata map[string]interface{}) (*types.AdminCreateUserResponse, error) {
	client, err := s.GetServiceClient(ctx)
	if err != nil {
		return nil, err
	}
	user, err := client.Auth.AdminCreateUser(types.AdminCreateUserRequest{
		Email:        email,
		Password:     &password,
		UserMetadata: userMetadata,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *sSupabaseService) DeleteUser(ctx context.Context, userID uuid.UUID) error {
	client, err := s.GetServiceClient(ctx)
	if err != nil {
		return err
	}
	err = client.Auth.AdminDeleteUser(types.AdminDeleteUserRequest{
		UserID: userID,
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *sSupabaseService) UpdateUser(ctx context.Context, userID uuid.UUID, email *string, password *string, userMetadata map[string]interface{}) (*types.AdminUpdateUserResponse, error) {
	client, err := s.GetServiceClient(ctx)
	if err != nil {
		return nil, err
	}
	user, err := client.Auth.AdminUpdateUser(types.AdminUpdateUserRequest{
		UserID:       userID,
		Email:        *email,
		Password:     *password,
		UserMetadata: userMetadata,
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *sSupabaseService) SignIn(ctx context.Context, email, password string) (*types.TokenResponse, error) {
	client, err := s.GetAnonClient(ctx)
	if err != nil {
		return nil, err
	}
	session, err := client.Auth.SignInWithEmailPassword(email, password)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *sSupabaseService) SignUp(ctx context.Context, email, password string, userMetadata map[string]interface{}) (*types.SignupResponse, error) {
	client, err := s.GetAnonClient(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Auth.Signup(types.SignupRequest{
		Email:    email,
		Password: password,
		Data:     userMetadata,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (s *sSupabaseService) RefreshToken(ctx context.Context, refreshToken string) (*types.TokenResponse, error) {
	client, err := s.GetAnonClient(ctx)
	if err != nil {
		return nil, err
	}
	session, err := client.Auth.RefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (s *sSupabaseService) SignOut(ctx context.Context, accessToken string) error {
	client, err := s.GetUserAuthenticatedClient(ctx, accessToken)
	if err != nil {
		return err
	}
	err = client.Auth.Logout()
	if err != nil {
		return err
	}
	return nil
}
