module github.com/fluxstack/fluxworks/event

go 1.20

require github.com/fluxstack/fluxworks/core v0.1.0

require (
	github.com/ThreeDotsLabs/watermill v1.2.0 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/lithammer/shortuuid/v3 v3.0.7 // indirect
	github.com/oklog/ulid v1.3.1 // indirect
)

replace github.com/fluxstack/fluxworks/core => ../core
