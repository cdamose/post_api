package repository

import (
	"context"
	"post_api/internal/post/model/dao"
)

type PostRepository interface {
	CreatePost(ctx context.Context, post dao.Post) (*dao.Post, error)
	ReadPost(ctx context.Context, post_id string) (*dao.Post, error)
	UpdatePost(ctx context.Context, post dao.Post) (bool, error)
	DeletePost(ctx context.Context, post_id string) (bool, error)
}
