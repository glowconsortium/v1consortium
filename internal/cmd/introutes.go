package cmd

import (
	"context"
	"v1consortium/internal/cmd/internalsetup"
	"v1consortium/internal/controller/auth"
	"v1consortium/internal/controller/gateway"
	"v1consortium/internal/controller/services"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/net/ghttp"
	"google.golang.org/grpc"
)

func SetupInternalRoutes(ctx context.Context, s *ghttp.Server) {

	internalSetup := internalsetup.NewInternalSetupConfig(ctx)

	if internalSetup.Enabled {
		internalSetup.SetupOrganizations()
		internalSetup.SetupUsers()
	}
}

func setupRpcRoutes(ctx context.Context) *grpcx.GrpcServer {
	c := grpcx.Server.NewConfig()
	c.Options = append(c.Options, []grpc.ServerOption{
		grpcx.Server.ChainUnary(
			grpcx.Server.UnaryValidate,
		)}...,
	)
	s := grpcx.Server.New(c)
	auth.Register(s)
	services.Register(s)
	gateway.Register(s)
	return s
}
