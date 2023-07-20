package main

import (
	"context"
	"github.com/fluxstack/fluxworks/app"
)

func main() {
	a := app.App{}
	_ = a.Start(context.Background())
}
