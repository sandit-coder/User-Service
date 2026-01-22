package di

import (
	middleware "UserCrud/internal/User/adapters/middlwaries"
	handlers "UserCrud/internal/User/adapters/primary/http-adapter/handlers/user"
	"UserCrud/internal/User/adapters/primary/http-adapter/router"
	repositories "UserCrud/internal/User/adapters/secondary/repositories/user"
	"UserCrud/internal/User/application/services/user"
	"UserCrud/internal/User/config"
	"UserCrud/internal/User/infrastructure/fiber"
	"UserCrud/internal/User/infrastructure/logger"
	"UserCrud/internal/User/infrastructure/postgres"
	"UserCrud/internal/User/infrastructure/validator"

	"go.uber.org/fx"
)

var ConfigModule = fx.Module("config",
	fx.Provide(
		config.NewFiberConfig,
		config.NewDBConfig),
	fx.Invoke(config.LoadConfig))

var RepositoryModule = fx.Module("repository",
	fx.Provide(
		repositories.NewUserRepository,
	))

var ServiceModule = fx.Module("service",
	fx.Provide(
		user.NewService,
	))

var HandlerModule = fx.Module("handler",
	fx.Provide(
		handlers.NewUserHandler))

var PostgresModule = fx.Module("db",
	fx.Provide(
		postgres.NewDb,
		postgres.NewTransactionManager))

var FiberModule = fx.Module("server",
	fx.Provide(
		server.NewFiberApp),
	fx.Invoke(server.RunFiberServer))

var LoggerModule = fx.Module("logger",
	fx.Provide(
		logger.NewLogger,
	))

var ValidatorModule = fx.Module("validator",
	fx.Provide(
		validator.NewValidator,
		validator.NewFiberValidator,
	))

var MiddlewareModule = fx.Module(
	"middleware",
	fx.Provide(
		middleware.NewSlogMiddleware,
	),
	fx.Invoke(middleware.Register))

var RouterModule = fx.Module("router",
	fx.Invoke(
		router.AppendUserRoutes))
