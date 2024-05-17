package domain

import (
	"context"
	"fmt"
	"post_api/internal/common/config"
	"post_api/internal/post/model/dao"
	"post_api/internal/post/repository"

	"github.com/sirupsen/logrus"
)

type PostDomain struct {
	logger     logrus.Entry
	config     config.Config
	repository repository.PostRepository
}

func NewPostDomain(logger logrus.Entry, config config.Config, repository repository.PostRepository) PostDomain {
	return PostDomain{
		logger:     logger,
		config:     config,
		repository: repository,
	}
}

func (ad *PostDomain) CreatePost(ctx context.Context, post dao.Post) (*dao.Post, error) {
	fmt.Println("debug 12")
	result, err := ad.repository.CreatePost(ctx, post)
	fmt.Println("debug 13")
	fmt.Println(err)
	if err != nil {
		return nil, err
	}

	return result, err

}
func (ad *PostDomain) UpdatePost(ctx context.Context, post dao.Post) (bool, error) {
	result, err := ad.repository.UpdatePost(ctx, post)
	if err != nil {
		return false, err
	}
	return result, nil

}

func (ad *PostDomain) DeletePost(ctx context.Context, post_id string) (bool, error) {
	result, err := ad.repository.DeletePost(ctx, post_id)
	if err != nil {
		return false, err
	}
	return result, nil
}

func (ad *PostDomain) ReadPost(ctx context.Context, post_id string) (*dao.Post, error) {
	result, err := ad.repository.ReadPost(ctx, post_id)
	if err != nil {
		return nil, err
	}
	return result, nil
}
