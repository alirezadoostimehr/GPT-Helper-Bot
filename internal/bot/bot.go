package bot

import (
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/bot/handler"
	tb "gopkg.in/telebot.v3"
)

type Bot struct {
	*tb.Bot
}

func NewBot(token string) (*Bot, error) {
	settings := tb.Settings{
		Token: token,
	}
	tgBot, err := tb.NewBot(settings)
	if err != nil {
		return nil, err
	}

	bot := &Bot{Bot: tgBot}
	bot.registerCommands(
		[]handler.Command{
			handler.NewStart(),
		})
	return bot, nil
}

func (b *Bot) registerCommands(commands []handler.Command) {
	for _, h := range commands {
		b.Handle(h.Command(), h.Handle)
	}
}
