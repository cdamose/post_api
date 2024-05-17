package main

import (
	"fmt"
	"log"

	"post_api/internal/post/container"
	"post_api/internal/post/ports"
	"post_api/internal/post/repository/adapters"

	//"post_api/internal/post/ports"
	"post_api/internal/common/config"
	"post_api/internal/common/database/migrations"

	//"post_api/internal/common/server"
	"net/http"

	"post_api/internal/common/genproto/post/api/protobuf/postv1connect"

	//"github.com/go-chi/chi/v5"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func main() {
	config := config.InitConfig()
	db, err := adapters.NewPostgreSQLConnection()
	if err != nil {
		log.Fatalln(err)
	}
	application, err := container.InitApplication(config, db)
	fmt.Println(err)
	migrations.ExecMigration(config.MYSQLUser, config.MYSQLPassword, config.MYSQLHost, config.MYSQLDatabase, "file://./db/migrations/auth")
	auther := ports.NewPostServer(application)
	mux := http.NewServeMux()
	path, handler := postv1connect.NewPostServiceHandler(auther)
	mux.Handle(path, handler)
	http.ListenAndServe(
		":8080",
		h2c.NewHandler(mux, &http2.Server{}),
	)

}
