package main

import (
	_ "v1consortium/internal/packed"

	_ "v1consortium/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"

	"v1consortium/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
