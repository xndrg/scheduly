package telegram

import (
	"errors"
	"fmt"
	"log/slog"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/xndrg/scheduly/internal/storage"
)

type Bot struct {
	bot     *tgbotapi.BotAPI
	Storage storage.Storage
	log     *slog.Logger
}

func NewBot(bot *tgbotapi.BotAPI, storage storage.Storage, log *slog.Logger) *Bot {
	return &Bot{bot: bot, Storage: storage, log: log}
}

func (b *Bot) Start() error {
	b.log.Info("Authorized on account ", slog.String("username", b.bot.Self.UserName))

	updates := b.initUpdatesChannel()
	return b.handleUpdates(updates)
}

func (b *Bot) initUpdatesChannel() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) error {
	fn := "telegram.handleUpdates"

	for update := range updates {
		if update.Message != nil { // if we got a message
			b.log.Info(fmt.Sprintf(
				"[%s] %s",
				update.Message.From.UserName,
				update.Message.Text,
			))

			if update.Message.IsCommand() {
				if err := b.handleCommand(update.Message); err != nil {
					return err
				}
				continue
			}

			if err := b.handleMessage(update.Message); err != nil {
				return err
			}
		}
	}

	return fmt.Errorf(
		"%s: execute statement: %w",
		fn,
		errors.New("end of handle updates cycle"),
	)
}
