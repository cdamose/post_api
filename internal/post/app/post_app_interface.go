package app

import (
	"context"
	postv1 "post_api/internal/common/genproto/post/api/protobuf"
	"post_api/internal/post/model/dto"
)

type PostApp interface {
	CreatePost(ctx context.Context, req *postv1.CreatePostRequest) (*dto.Post, error)
	UpdatePost(ctx context.Context, req *postv1.UpdatePostRequest) (bool, error)
	// DeletePost(ctx context.Context, post_id string) (*dao.Post, error)
	// ReadPost(ctx context.Context, post_id string) (*dao.Post, error)
}
