package bot

import (
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/bot/handler"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/openai"
	tb "gopkg.in/telebot.v3"
)

type Bot struct {
	*tb.Bot
}

func NewBot(token string, openaiClient *openai.Client) (*Bot, error) {
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
			handler.NewText(openaiClient),
		})
	return bot, nil
}

func (b *Bot) registerCommands(commands []handler.Command) {
	for _, h := range commands {
		b.Handle(h.Command(), h.Handle, h.Middleware()...)
	}
}
