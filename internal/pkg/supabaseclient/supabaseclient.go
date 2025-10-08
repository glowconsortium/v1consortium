package supabaseclient

import "github.com/supabase-community/supabase-go"

type SupabaseConfig struct {
	Url          string
	PublicApiKey string
	SecretApiKey string
}

type SupabaseClient struct {
	anonClient    *supabase.Client
	serviceClient *supabase.Client
}

func NewSupabaseClient(config SupabaseConfig) (*SupabaseClient, error) {
	anonClient, err := supabase.NewClient(config.Url, config.PublicApiKey, nil)
	if err != nil {
		return nil, err
	}
	serviceClient, err := supabase.NewClient(config.Url, config.SecretApiKey, nil)
	if err != nil {
		return nil, err
	}
	return &SupabaseClient{
		anonClient:    anonClient,
		serviceClient: serviceClient,
	}, nil
}

func (s *SupabaseClient) GetAnonClient() *supabase.Client {
	return s.anonClient
}

func (s *SupabaseClient) GetServiceClient() *supabase.Client {
	return s.serviceClient
}

// func (s *SupabaseClient) GetUserAuthenticatedClient(accessToken string) *supabase.Client {
// 	return
// }
