package repository

import (
	"context"
	"post_api/internal/post/model/dao"
)

type Repository interface {
	Ping(ctx context.Context) (*dao.Ping, error)
}
