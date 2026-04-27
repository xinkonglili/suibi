package main

import (
	"goframeP/frame/internal/cmd"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {

	cmd.Main.Run(gctx.GetInitCtx())
}
