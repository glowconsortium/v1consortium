package userservice

import (
	"context"
	"v1consortium/internal/pkg/auth0client"
	"v1consortium/internal/service"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

type sUserService struct{}

func init() {
	service.RegisterUserService(new())
}
func new() service.IUserService {
	return &sUserService{}
}

func (s *sUserService) GetUserInfo(ctx context.Context) {}

func (s *sUserService) GetAuth0Config(ctx context.Context) (auth0client.Config, error) {

	Domain := g.Cfg().MustGet(ctx, "auth0.domain").String()
	ClientID := g.Cfg().MustGet(ctx, "auth0.clientId").String()
	ClientSecret := g.Cfg().MustGet(ctx, "auth0.clientSecret").String()
	Audience := g.Cfg().MustGet(ctx, "auth0.audience").String()
	Connection := g.Cfg().MustGet(ctx, "auth0.connection").String()

	if Domain == "" || ClientID == "" || ClientSecret == "" || Audience == "" || Connection == "" {
		return auth0client.Config{}, gerror.New("auth0 config is not set")
	}
	return auth0client.Config{
		Domain:       Domain,
		ClientID:     ClientID,
		ClientSecret: ClientSecret,
		Audience:     Audience,
		Connection:   Connection,
	}, nil
}
