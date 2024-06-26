package main

import (
	"log/slog"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xndrg/scheduly/internal/config"
	"github.com/xndrg/scheduly/internal/lib/logger/sl"
	"github.com/xndrg/scheduly/internal/storage/sqlite"
	"github.com/xndrg/scheduly/internal/telegram"
	"github.com/xndrg/scheduly/pkg/scraper/mau"
)

const (
	envLocal = "local"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	log := setupLogger(cfg.Env)
	log.Info("starting scheduly", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")

	storage, err := sqlite.New(cfg.StoragePath)
	if err != nil {
		log.Error("failed to init storage", sl.Err(err))
		os.Exit(1)
	}

	scraper := mau.New()

	bot, err := tgbotapi.NewBotAPI(cfg.TelegramToken)
	if err != nil {
		log.Error("failed to init bot", sl.Err(err))
		os.Exit(1)
	}
	tgBot := telegram.NewBot(bot, storage, log, scraper)
	log.Info("starting telegram-bot")
	if err := tgBot.Start(); err != nil {
		log.Error("bot runtime error", sl.Err(err))
	}
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
