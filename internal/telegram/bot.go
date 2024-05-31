package telegram

import (
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

func (b *Bot) Start() {
	b.log.Info("Authorized on account ", slog.String("username", b.bot.Self.UserName))

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := b.bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // if we got a message
			b.log.Debug(fmt.Sprintf("[%s] %s", update.Message.From.UserName, update.Message.Text))

			msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

			b.bot.Send(msg)
		}
	}
}
