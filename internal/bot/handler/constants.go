package handler

import tb "gopkg.in/telebot.v3"

const (
	GroupNotRegisteredErrorMessage = "This group is not registered! Please register the group by removing and adding the bot again."
	InternalErrorMessage           = "Internal error occurred. Please try again later and inform admins."
	NotRegisteredErrorMessage      = "You are not registered! Please start the bot by typing /start in the chat with the bot."
	GroupRegisteredSuccessMessage  = "This group has been successfully registered!"
	GroupAlreadyRegisteredMessage  = "This group is already registered!"
	UserRegisteredSuccessMessage   = "You are successfully registered!"
	UserAlreadyRegisteredMessage   = "You are already registered!"
	TopicCreationSuccessMessage    = "This topic has been successfully created!\nThe model which is using is %v"
)

const (
	DefaultOpenAIModel = "gpt-3.5-turbo"
)

const (
	MessageMaxLength = 4096
)

var (
	ReactionSuccess = tb.Reaction{
		Emoji: "👍",
		Type:  "emoji",
	}
)
