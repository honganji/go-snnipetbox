package main

import (
	"database/sql"
	"html/template"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
	"github.com/honganji/go-snippetbox/internal/models"
	"github.com/joho/godotenv"
)

// application struct holds the application-wide dependencies
type application struct {
	logger         *slog.Logger
	snippets       *models.SnippetModel
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}

func main() {
	// set up the logger
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	// load environment variables from .env file
	loadEnv(logger)
	addr := getEnv("ADDR", ":4000")
	dsn := getEnv("DSN", "root:123@/snippetbox")

	// open database connection pool and connect to the database
	db, err := openDB(dsn + "?parseTime=true")
	if err != nil {
		logger.Error("unable to connect to database", slog.String("dsn", dsn), slog.String("error", err.Error()))
		os.Exit(1)
	}
	// ensure the database connection pool is closed before the main function exits
	defer db.Close()

	// initialize the template cache
	templateCache, err := newTemplateCache()
	if err != nil {
		logger.Error("unable to create template cache", slog.String("error", err.Error()))
		os.Exit(1)
	}

	formDecoder := form.NewDecoder()

	// initialize session manager
	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour

	// initialize a new application instance
	app := &application{
		logger: logger,
		snippets: &models.SnippetModel{
			DB: db,
		},
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
	}

	// start the HTTP server
	logger.Info("starting server", slog.String("addr", addr))
	err = http.ListenAndServe(addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

// loadEnv loads environment variables from a .env file
func loadEnv(logger *slog.Logger) {
	if err := godotenv.Load(); err != nil {
		logger.Error("error loading .env file", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

// retrieve environment variable or return default value
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// openDB opens a database connection pool and verifies the connection
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
