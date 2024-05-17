package app

import (
	"context"
	"post_api/internal/post/model/dto"
)

type PingApp interface {
	Ping(ctx context.Context) (*dto.Ping, error)
}
