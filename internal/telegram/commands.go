package telegram

import (
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	replyStart   = "Started"
	replyUnknown = "Unknown command"
)

func (b *Bot) startCommandHandler(message *tgbotapi.Message) error {
	const fn = "telegram.startCommandHandler"

	msg := tgbotapi.NewMessage(message.Chat.ID, replyStart)
	_, err := b.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}

func (b *Bot) unknownCommandHandler(message *tgbotapi.Message) error {
	const fn = "telegram.unknownCommandHandler"

	msg := tgbotapi.NewMessage(message.Chat.ID, replyUnknown)
	_, err := b.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("%s: %w", fn, err)
	}

	return nil
}
