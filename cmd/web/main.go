package main

import (
	"database/sql"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/honganji/go-snippetbox/internal/models"
	"github.com/joho/godotenv"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	if err := loadEnv(); err != nil {
		logger.Error("error loading .env file", slog.String("error", err.Error()))
		os.Exit(1)
	}

	addr := getEnv("ADDR", ":4000")
	dsn := getEnv("DSN", "root:123@/snippetbox")

	db, err := openDB(dsn + "?parseTime=true")
	if err != nil {
		logger.Error("unable to connect to database", slog.String("dsn", dsn), slog.String("error", err.Error()))
		os.Exit(1)
	}
	defer db.Close()

	app := &application{
		logger: logger,
		snippets: &models.SnippetModel{
			DB: db,
		},
	}

	logger.Info("starting server", slog.String("addr", addr))
	err = http.ListenAndServe(addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func loadEnv() error {
	return godotenv.Load()
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
