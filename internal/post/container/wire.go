//go:build wireinject
// +build wireinject

package container

import (
	"post_api/internal/common/config"
	"post_api/internal/common/logs"
	"post_api/internal/post/app"
	"post_api/internal/post/domain"
	"post_api/internal/post/repository"
	"post_api/internal/post/repository/adapters"

	"github.com/google/wire"
	"github.com/jmoiron/sqlx"
)

func InitializeDomain(config config.Config, db *sqlx.DB) (app.PingApp, error) {
	wire.Build(app.NewPingApplication, domain.NewPingDomain, adapters.NewPostgressPingRepository, logs.Init,
		wire.Bind(new(app.PingApp), new(app.PingApplication)),
		wire.Bind(new(repository.Repository), new(*adapters.PostgressPingRepository)))
	//wire.Build(config.InitConfig)
	return app.PingApplication{}, nil
}

func InitializePingApplication(config config.Config, db *sqlx.DB) (app.PingApp, error) {
	wire.Build(app.NewPingApplication, domain.NewPingDomain, adapters.NewPostgressPingRepository, logs.Init,
		wire.Bind(new(app.PingApp), new(app.PingApplication)),
		wire.Bind(new(repository.Repository), new(*adapters.PostgressPingRepository)))
	//wire.Build(config.InitConfig)
	return app.PingApplication{}, nil
}

func InitializePostApplication(config config.Config, db *sqlx.DB) (app.PostApp, error) {
	wire.Build(app.NewPostApplication, domain.NewPostDomain, adapters.NewPostgresPostRepository, logs.Init,
		wire.Bind(new(app.PostApp), new(app.PostApplication)),
		wire.Bind(new(repository.PostRepository), new(*adapters.PostgresPostRepository)))
	//wire.Build(config.InitConfig)
	return app.PostApplication{}, nil

}
