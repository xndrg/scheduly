package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleMessage(message *tgbotapi.Message) error {
	const fn = "telegram.handleMessage"

	msg := tgbotapi.NewMessage(message.Chat.ID, message.Text)
	_, err := b.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (b *Bot) handleCommand(message *tgbotapi.Message) error {
	switch message.Command() {
	case "start":
		return b.startCommandHandler(message)
	default:
		return b.unknownCommandHandler(message)
	}
}
