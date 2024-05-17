package container

import (
	"post_api/internal/post/app"
)

type Application struct {
	PingApplication app.PingApp
	PostApplication app.PostApp
}
