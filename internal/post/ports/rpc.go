package ports

import (
	"context"
	"fmt"
	postv1 "post_api/internal/common/genproto/post/api/protobuf"
	"post_api/internal/post/container"

	"connectrpc.com/connect"
)

type PostServer struct {
	Application container.Application
}

func NewPostServer(application container.Application) *PostServer {
	return &PostServer{Application: application}

}

func (p *PostServer) CreatePost(ctx context.Context, re *connect.Request[postv1.CreatePostRequest]) (*connect.Response[postv1.CreatePostResponse], error) {

	dto_obj, err := p.Application.PostApplication.CreatePost(ctx, re.Msg)
	if err != nil {
		res := connect.NewResponse(&postv1.CreatePostResponse{
			ErrorMessage: "Something went wrong... please try again later",
		})
		return res, nil
	}
	res := connect.NewResponse(&postv1.CreatePostResponse{
		PostId:          dto_obj.ID,
		Title:           dto_obj.Title,
		Content:         dto_obj.Content,
		Author:          dto_obj.Author,
		PublicationDate: dto_obj.PublicationDate.String(),
		Tags:            dto_obj.Tags,
	})
	return res, nil

}
func (p *PostServer) ReadPost(context.Context, *connect.Request[postv1.ReadPostRequest]) (*connect.Response[postv1.ReadPostResponse], error) {
	return nil, nil
}
func (p *PostServer) UpdatePost(ctx context.Context, re *connect.Request[postv1.UpdatePostRequest]) (*connect.Response[postv1.UpdatePostResponse], error) {
	fmt.Println(re.Msg.PostId)
	dto_obj, err := p.Application.PostApplication.UpdatePost(ctx, re.Msg)

	if err != nil {
		res := connect.NewResponse(&postv1.UpdatePostResponse{
			ErrorMessage: "Something went wrong... please try again later",
		})
		return res, nil
	}
	res := connect.NewResponse(&postv1.UpdatePostResponse{})
	if dto_obj {
		res.Msg.Message = "Successfully updated..!"
	} else {
		res.Msg.ErrorMessage = "Something went wrong... please try again later"
	}
	return res, nil
}
func (p *PostServer) DeletePost(context.Context, *connect.Request[postv1.DeletePostRequest]) (*connect.Response[postv1.DeletePostResponse], error) {
	return nil, nil
}

// func (av *AuthServer) SignupWithPhoneNumber(ctx context.Context, re *connect.Request[authv1.PhoneNumber]) (*connect.Response[authv1.SignUpResponse], error) {
// 	dto_obj, err := av.Application.AuthApplication.SignUp(ctx, re.Msg.Number)
// 	if err != nil {
// 		res := connect.NewResponse(&authv1.SignUpResponse{
// 			Error: &authv1.Error{
// 				Message: "Something went wrong",
// 				Code:    "10002",
// 			},
// 		})
// 		return res, nil
// 	}
// 	res := connect.NewResponse(&authv1.SignUpResponse{
// 		UserId:    dto_obj.UserId,
// 		IsVerfied: dto_obj.IsVerified,
// 		CreatedAt: dto_obj.CreatedAt,
// 	})
// 	return res, nil
// }
// func (av *AuthServer) VerifyAccount(ctx context.Context, re *connect.Request[authv1.VerifyAccountRequest]) (*connect.Response[authv1.VerifyAccountResponse], error) {
// 	dto_obj, _ := av.Application.AuthApplication.VerifyAccount(ctx, re.Msg.UserId, re.Msg.Code)
// 	res := connect.NewResponse(&authv1.VerifyAccountResponse{
// 		Message: dto_obj.Message,
// 	})
// 	return res, nil
// }
// func (av *AuthServer) Login(ctx context.Context, re *connect.Request[authv1.LoginRequest]) (*connect.Response[authv1.LoginResponse], error) {
// 	dto_obj, _ := av.Application.AuthApplication.Login(ctx, re.Msg.PhoneNumber.Number)

// 	res := connect.NewResponse(&authv1.LoginResponse{
// 		UserId:  dto_obj.UserID,
// 		Message: dto_obj.Message,
// 	})
// 	return res, nil
// }
// func (av *AuthServer) OTPGenerate(ctx context.Context, re *connect.Request[authv1.PhoneNumber]) (*connect.Response[authv1.Response], error) {
// 	dto_obj, err := av.Application.AuthApplication.GenerateOTP(ctx, re.Msg.Number)
// 	if err != nil {
// 		return nil, connect.NewError(connect.CodeUnknown, err)
// 	}
// 	res := connect.NewResponse(&authv1.Response{
// 		Message: dto_obj.Message,
// 	})
// 	return res, nil
// }
// func (av *AuthServer) GetProfile(ctx context.Context, re *connect.Request[authv1.PhoneNumber]) (*connect.Response[authv1.ProfileResponse], error) {
// 	dto_obj, err := av.Application.AuthApplication.GetUserProfile(ctx, re.Msg.Number)
// 	if err != nil {
// 		res := connect.NewResponse(&authv1.ProfileResponse{
// 			Error: &authv1.Error{
// 				Message: "Something went wrong",
// 				Code:    "10002",
// 			},
// 		})
// 		return res, nil
// 	}
// 	res := connect.NewResponse(&authv1.ProfileResponse{
// 		IsVerfied:   dto_obj.IsVerified,
// 		CreatedAt:   dto_obj.CreatedAt,
// 		PhoneNumber: dto_obj.PhoneNumber,
// 		VerfiedAt:   dto_obj.VerfiedAt,
// 	})

// 	return res, nil
// }
