package main

import (
	"context"
	"github.com/weflux/fluxworks/app"
)

func main() {
	a := app.App{}
	_ = a.Start(context.Background())
}
