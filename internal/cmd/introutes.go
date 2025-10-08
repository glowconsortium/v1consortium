package cmd

import (
	"context"
	"v1consortium/internal/cmd/internalsetup"
	"v1consortium/internal/controller/auth"
	"v1consortium/internal/controller/gateway"
	"v1consortium/internal/controller/services"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"google.golang.org/grpc"
)

func SetupInternalStartupData(ctx context.Context) {

	internalSetup := internalsetup.NewInternalSetupConfig(ctx)

	if internalSetup.Enabled {
		err := internalSetup.SetupOrganizations()
		if err != nil {
			g.Log().Errorf(ctx, "Failed to setup organizations: %v", err)
		}
		err = internalSetup.SetupUsers()
		if err != nil {
			g.Log().Errorf(ctx, "Failed to setup users: %v", err)
		}
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
