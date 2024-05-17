package conversion

import (
	"post_api/internal/post/model/dao"
	"post_api/internal/post/model/dto"
)

func ConvertToDto(post dao.Post) dto.Post {
	return dto.Post{
		ID:              post.PostID,
		Title:           post.Title,
		Content:         post.Content,
		Author:          post.Author,
		PublicationDate: post.PublicationDate,
		Tags:            post.Tags,
	}
}
