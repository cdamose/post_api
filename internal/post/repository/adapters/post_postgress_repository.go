package adapters

import (
	"context"
	"fmt"

	"post_api/internal/common/config"
	"post_api/internal/post/model/dao"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/sirupsen/logrus"
)

type PostgresPostRepository struct {
	db     *sqlx.DB
	logger logrus.Entry
	config config.Config
}

func NewPostgresPostRepository(db *sqlx.DB, logger logrus.Entry, config config.Config) *PostgresPostRepository {
	if db == nil {
		panic("missing db")
	}
	return &PostgresPostRepository{db: db, logger: logger, config: config}
}

func (m PostgresPostRepository) CreatePost(ctx context.Context, post dao.Post) (*dao.Post, error) {

	query := `INSERT INTO posts (post_id, title, content, author, publication_date, tags)
    VALUES (uuid_generate_v4(), :title, :content, :author, :publication_date, :tags) RETURNING *`
	param := map[string]interface{}{
		"title":            post.Title,
		"content":          post.Content,
		"author":           post.Author,
		"publication_date": post.PublicationDate,
		"tags":             post.Tags,
	}
	var res dao.Post
	result, err := m.db.NamedQueryContext(ctx, query, param)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	if !result.Next() {
		return nil, errors.New("no rows returned after insert")
	}
	fmt.Println(result.Rows)
	err = result.StructScan(&res)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan result into struct")
	}

	return &res, nil
}

func (m PostgresPostRepository) ReadPost(ctx context.Context, post_id string) (*dao.Post, error) {
	query := `select * from posts where post_id = :post_id`
	param := map[string]interface{}{
		"post_id": post_id,
	}
	var post dao.Post
	result, err := m.db.NamedQueryContext(ctx, query, param)

	if err != nil {
		return nil, err
	}
	defer result.Close()
	if !result.Next() {
		return nil, errors.New("no rows returned after insert")
	}
	err = result.StructScan(&post)
	if err != nil {
		return nil, errors.Wrap(err, "failed to scan result into struct")
	}
	return &post, nil

}
func (m PostgresPostRepository) UpdatePost(ctx context.Context, post dao.Post) (bool, error) {
	query := ` UPDATE posts
    SET title = :title,
        content = :content,
        author = :author,
        tags = :tags
    WHERE post_id = :post_id `

	param := map[string]interface{}{
		"title":   post.Title,
		"content": post.Content,
		"author":  post.Author,
		"tags":    post.Tags,
		"post_id": post.PostID,
	}
	_, err := m.db.NamedQueryContext(ctx, query, param)
	fmt.Println(err)
	if err != nil {
		return false, err
	}
	return true, nil
}
func (m PostgresPostRepository) DeletePost(ctx context.Context, post_id string) (bool, error) {
	query := `DELETE FROM posts
    WHERE post_id = :post_id`
	param := map[string]interface{}{
		"post_id": post_id,
	}
	_, err := m.db.NamedQueryContext(ctx, query, param)

	if err != nil {
		return false, err
	}
	return true, nil
}
