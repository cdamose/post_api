package domain

import (
	"context"
	"fmt"
	"post_api/internal/common/config"
	"post_api/internal/post/model/dto"
	"post_api/internal/post/repository"

	"github.com/sirupsen/logrus"
)

type PingDomain struct {
	logger     logrus.Entry
	config     config.Config
	repository repository.Repository
}

func NewPingDomain(logger logrus.Entry, config config.Config, repository repository.Repository) PingDomain {
	return PingDomain{
		logger:     logger,
		config:     config,
		repository: repository,
	}
}

func (ud *PingDomain) Ping(ctx context.Context) (*dto.Ping, error) {
	ud.logger.Info("in ==> domain layer ==> Ping()")
	data_obj, err := ud.repository.Ping(ctx)
	dto_obj := &dto.Ping{}
	if err == nil {
		fmt.Println(data_obj)
		dto_obj.Message = data_obj.Message
		// dto_obj.ID = data_obj.ID
		// dto_obj.Name = data_obj.Name
	} else {
		return nil, err
	}
	return dto_obj, nil
}
