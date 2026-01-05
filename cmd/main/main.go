// cmd/restaurant/main.go
package main

import (
	"UserCrud/internal/bootstrap/di"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		di.ConfigModule,
		di.ValidatorModule,
		di.LoggerModule,
		di.PostgresModule,
		di.FiberModule,
		di.MiddlewareModule,
		di.RepositoryModule,
		di.ServiceModule,
		di.HandlerModule,
		di.RouterModule,
	).Run()
}
