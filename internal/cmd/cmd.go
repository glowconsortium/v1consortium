package cmd

import (
	"context"

	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// s := g.Server()

			// // Setup internal start data.
			SetupInternalStartupData(ctx)

			// s.Group("/", func(group *ghttp.RouterGroup) {
			// 	group.Middleware(ghttp.MiddlewareHandlerResponse)
			// 	group.Bind(
			// 		hello.NewV1(),
			// 	)
			// })
			s := setupRpcRoutes(ctx)
			s.Run()
			return nil
		},
	}
)
