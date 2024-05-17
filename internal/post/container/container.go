package container

import (
	"post_api/internal/common/config"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/gommon/log"
)

func InitApplication(config config.Config, db *sqlx.DB) (Application, error) {
	ping_app, err := InitializePingApplication(config, db)
	log.Info(err)
	post_app, err := InitializePostApplication(config, db)
	log.Info(err)
	return Application{
		PingApplication: ping_app,
		PostApplication: post_app,
	}, err

}
