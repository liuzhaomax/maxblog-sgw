package main

import (
	"context"
	"github.com/liuzhaomax/maxblog-sgw/internal/app"
	"github.com/liuzhaomax/maxblog-sgw/internal/core"
)

func main() {
	app.Launch(
		context.Background(),
		app.SetConfigFile(core.LoadEnv()),
		app.SetWWWDir("www"),
	)
}
