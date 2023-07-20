package app

import (
	"context"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

type App struct {
	cqrs.Facade
	Config *Config
}

func NewApp() *App {
	return &App{}
}

func (a *App) Start(ctx context.Context) error {
	return nil
}

func (a *App) Stop(ctx context.Context) error {
	return nil
}
