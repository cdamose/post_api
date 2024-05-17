package app

import (
	"context"
	"post_api/internal/common/config"
	postv1 "post_api/internal/common/genproto/post/api/protobuf"
	"post_api/internal/post/domain"
	"post_api/internal/post/model/conversion"
	"post_api/internal/post/model/dao"
	"post_api/internal/post/model/dto"
	"time"

	"github.com/sirupsen/logrus"
)

type PostApplication struct {
	logger logrus.Entry
	config config.Config
	domain domain.PostDomain
}

func NewPostApplication(logger logrus.Entry, config config.Config, domain domain.PostDomain) PostApplication {
	return PostApplication{
		logger: logger,
		config: config,
		domain: domain,
	}
}

func (app PostApplication) CreatePost(ctx context.Context, req *postv1.CreatePostRequest) (*dto.Post, error) {
	pubDate, err := time.Parse("2006-01-02", req.PublicationDate)
	if err != nil {
		return nil, err
	}
	req_obj := dao.Post{Title: req.Title, Content: req.Content, Author: req.Author, PublicationDate: pubDate, Tags: req.Tags}
	domain_obj, err := app.domain.CreatePost(ctx, req_obj)
	if err != nil {
		return nil, err
	}
	dto_obj := conversion.ConvertToDto(*domain_obj)
	return &dto_obj, nil
}

func (app PostApplication) UpdatePost(ctx context.Context, req *postv1.UpdatePostRequest) (bool, error) {
	req_obj := dao.Post{Title: req.Title, Content: req.Content, Author: req.Author, Tags: req.Tags, PostID: req.PostId}
	domain_obj, err := app.domain.UpdatePost(ctx, req_obj)
	if err != nil {
		return false, err
	}
	return domain_obj, nil

}

func (app PostApplication) DeletePost(ctx context.Context, post_id string) (bool, error) {
	domain_obj, err := app.domain.DeletePost(ctx, post_id)
	if err != nil {
		return false, err
	}
	return domain_obj, nil

}

func (app PostApplication) ReadPost(ctx context.Context, post_id string) (*dao.Post, error) {
	domain_obj, err := app.domain.ReadPost(ctx, post_id)
	if err != nil {
		return nil, err
	}
	return domain_obj, nil

}
