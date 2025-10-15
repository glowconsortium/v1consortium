package main

import (
	_ "v1consortium/internal/packed"

	_ "v1consortium/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"

	"v1consortium/internal/cmd"
)

func main() {
	err := cmd.Main.AddCommand(&cmd.Combined, &cmd.DBOSWorker)
	if err != nil {
		panic(err)
	}
	cmd.Main.Run(gctx.GetInitCtx())
}
