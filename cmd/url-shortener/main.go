package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"

	"url-shortener/internal/config"
	"url-shortener/internal/repository"
	"url-shortener/internal/repository/in_memory"
	"url-shortener/internal/repository/postgres"
	urlService "url-shortener/internal/service/url"
	"url-shortener/internal/util/logger/handlers/slogpretty"

	_ "github.com/jackc/pgx/v5/stdlib"

	getHandler "url-shortener/internal/http-server/handlers/url/get"
	saveHandler "url-shortener/internal/http-server/handlers/url/save"
	mwLogger "url-shortener/internal/http-server/middleware/logger"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log.Println(cfg)

	log := setupLogger(cfg.Env)

	log.Info("starting url-shortener", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	var urlInfoRepository repository.URLInfoRepository

	switch getInMemoryFlag() {
	case true:
		urlInfoRepository = in_memory.NewRepository()
	default:
		db, err := sqlx.Connect("pgx", cfg.PGDSN)
		if err != nil {
			log.Error(fmt.Errorf("failed to connect to pg: %w", err).Error())
		}

		urlInfoRepository = postgres.NewRepository(db)
	}

	urlService := urlService.NewService(
		urlInfoRepository,
	)

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Post("/url", saveHandler.New(log, urlService))
	router.Get("/{alias}", getHandler.New(log, urlService))

	log.Info("starting server", slog.String("address", cfg.Address))

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Error("failed to start server")
	}

	log.Error("server stopped")

}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = setupPrettySlog()
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)

	default: // If env config is invalid, set prod settings by default due to security
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}

func setupPrettySlog() *slog.Logger {
	opts := slogpretty.PrettyHandlerOptions{
		SlogOpts: &slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	}

	handler := opts.NewPrettyHandler(os.Stdout)

	return slog.New(handler)
}

func getInMemoryFlag() bool {
	if len(os.Args) < 2 {
		return false
	}

	if os.Args[1] == "--in_memory" {
		return true
	}

	return false
}
