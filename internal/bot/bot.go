package bot

import (
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/bot/button"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/bot/handler"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/database/postgres"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/openai"
	tb "gopkg.in/telebot.v3"
)

type Bot struct {
	*tb.Bot
}

func NewBot(token string, openaiClient *openai.Client, postgresConn *postgres.ConnectionPool) (*Bot, error) {
	settings := tb.Settings{
		Token: token,
	}
	tgBot, err := tb.NewBot(settings)
	if err != nil {
		return nil, err
	}

	bot := &Bot{Bot: tgBot}

	userRepo := postgres.NewUserRepo(postgresConn)
	groupRepo := postgres.NewGroupRepo(postgresConn)
	topicRepo := postgres.NewTopicRepo(postgresConn)
	messageRepo := postgres.NewMessageRepo(postgresConn)

	bot.registerCommands([]handler.Command{
		handler.NewStart(userRepo),
		handler.NewGroupAddition(userRepo, groupRepo),
		handler.NewTopicCreation(groupRepo, topicRepo),
		handler.NewTopicCreated(topicRepo),
		handler.NewText(openaiClient, topicRepo, messageRepo),
		handler.NewSetOpenAIModel(),
		handler.NewCloseTopic(topicRepo),
	})

	buttons := make([]button.ButtonHandler, 0)
	for model := range openai.GptModels {
		buttons = append(buttons, button.NewSetModel(groupRepo, model))
	}

	bot.registerButtons(buttons)
	return bot, nil
}

func (b *Bot) registerCommands(commands []handler.Command) {
	for _, h := range commands {
		b.Handle(h.Command(), h.Handle, h.Middleware()...)
	}
}

func (b *Bot) registerButtons(buttons []button.ButtonHandler) {
	for _, h := range buttons {
		b.Handle(h.CallbackUnique(), h.Handle, h.Middleware()...)
	}
}
