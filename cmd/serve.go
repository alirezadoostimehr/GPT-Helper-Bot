package cmd

import (
	"fmt"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/bot"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/config"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/openai"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start serving",
	Run:   serve,
}

func init() {
	rootCmd.AddCommand(serveCmd)
}

func serve(cmd *cobra.Command, args []string) {
	fmt.Println("serve called")
	if err := config.Load(); err != nil {
		panic(err)
	}

	openaiClient := openai.NewGPT(config.GlobalConfig.OpenAI.APIKey)

	tgbot, err := bot.NewBot(config.GlobalConfig.BOT.TOKEN, openaiClient)
	if err != nil {
		panic(err)
	}

	tgbot.Start()

	return
}
