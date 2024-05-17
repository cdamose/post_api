package app

import (
	"context"

	"post_api/internal/common/config"
	"post_api/internal/post/domain"
	"post_api/internal/post/model/dto"

	"github.com/sirupsen/logrus"
)

type PingApplication struct {
	logger logrus.Entry
	config config.Config
	domain domain.PingDomain
}

func NewPingApplication(logger logrus.Entry, config config.Config, domain domain.PingDomain) PingApplication {
	return PingApplication{
		logger: logger,
		config: config,
		domain: domain,
	}
}
func (p PingApplication) Ping(ctx context.Context) (*dto.Ping, error) {
	return p.domain.Ping(ctx)

}
