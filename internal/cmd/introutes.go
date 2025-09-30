package cmd

import (
	"context"
	"v1consortium/internal/cmd/internalsetup"

	"github.com/gogf/gf/v2/net/ghttp"
)

func SetupInternalRoutes(ctx context.Context, s *ghttp.Server) {

	internalSetup := internalsetup.NewInternalSetupConfig(ctx)

	if internalSetup.Enabled {
		internalSetup.SetupOrganizations()
		internalSetup.SetupUsers()
	}
}
