module github.com/weflux/fluxworks/contrib/kratos

go 1.20
replace (
	github.com/weflux/fluxworks => ../../
)

require (
	github.com/weflux/fluxworks v0.0.0-20230724074707-4330c2482c5e
	github.com/go-kratos/kratos/contrib/log/zap/v2 v2.0.0-20230706115902-bffc1a0989a6
	github.com/go-kratos/kratos/v2 v2.6.3
	go.uber.org/zap v1.24.0
	google.golang.org/protobuf v1.31.0
)

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
)
