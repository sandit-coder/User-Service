// cmd/restaurant/main.go
package main

import (
	"UserCrud/internal/bootstrap/di"

	"go.uber.org/fx"
)

func main() {
	fx.New(
		di.LoggerModule,
		di.ConfigModule,
		di.MiddlewareModule,

		di.ValidatorModule,
		di.PostgresModule,
		di.FiberModule,

		di.RepositoryModule,
		di.ServiceModule,
		di.HandlerModule,
		di.RouterModule,
	).Run()
}
