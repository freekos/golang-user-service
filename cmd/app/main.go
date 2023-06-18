package main

import (
	"fmt"
	migrations "golang-postgres"
	"golang-postgres/pkg/repository"
	"golang-postgres/pkg/repository/database"
	"net/http"
	"os"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

type Config struct {
	repo repository.Repository
}

func main() {
	db, err := database.NewPostgres(database.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  os.Getenv("POSTGRES_SSL_MODE"),
	})
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Bad practice, but it works.
	if err := RunMigrate(db); err != nil {
		panic(err)
	}

	app := Config{
		repo: *repository.NewRepository(db),
	}

	srv := http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("PORT")),
		Handler: app.Routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func RunMigrate(db *sqlx.DB) error {
	goose.SetBaseFS(migrations.EmbedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		return err
	}

	if err := goose.Up(db.DB, "migrations"); err != nil {
		return err
	}

	return nil
}
