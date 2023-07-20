package app

import "github.com/ThreeDotsLabs/watermill/components/cqrs"

type App struct {
	cqrs.Facade
	Config *Config
}
