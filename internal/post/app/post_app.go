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

// func (app AuthApplication) SignUp(ctx context.Context, phone_number string) (*dto.UserProfile, error) {
// 	domain_obj, err := app.domain.CreateUserProfile(ctx, phone_number)
// 	if err != nil {
// 		return nil, err
// 	}
// 	dto_obj := conversion.ConvertToUpdatedUserProfile(*domain_obj)
// 	return &dto_obj, nil
// }

// func (app AuthApplication) VerifyAccount(ctx context.Context, user_id string, otp string) (*dto.VerifiedAccountResp, error) {
// 	var resp = &dto.VerifiedAccountResp{}
// 	domain_obj, err := app.domain.VerifyAccount(ctx, user_id, otp)

// 	if err != nil {
// 		resp.Message = "Account Not Able to verified , pls try again"
// 	}
// 	if domain_obj {
// 		resp.Message = "Account Verified Successfully..!"
// 	}
// 	return resp, err

// }

// func (app AuthApplication) GenerateOTP(ctx context.Context, phone_number string) (*dto.CommonResponse, error) {
// 	var resp = &dto.CommonResponse{}
// 	domain_obj, err := app.domain.GenerateOTP(ctx, phone_number)
// 	if err != nil {
// 		resp.Message = "Something went wrong as of now , try again later "
// 	}
// 	if domain_obj {
// 		resp.Message = "OTP Sent Successfully..!"
// 	} else {
// 		resp.Message = "Seems to be given mobile number not register with our system"
// 	}
// 	return resp, err

// }

// func (app AuthApplication) Login(ctx context.Context, phone_number string) (*dto.LoginResponse, error) {
// 	var resp = &dto.LoginResponse{}
// 	domain_obj, err := app.domain.Login(ctx, phone_number)
// 	if err != nil {
// 		resp.Message = "Something went wrong,please try again"
// 	}
// 	resp.UserID = *domain_obj
// 	return resp, err

// }
