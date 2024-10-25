package cmd

import (
	"context"
	"fmt"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/bot"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/config"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/database/postgres"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/openai"
	log "github.com/sirupsen/logrus"
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
	log.SetReportCaller(true)
	ctx := context.Background()

	fmt.Println("serve called")
	if err := config.Load(); err != nil {
		panic(err)
	}

	openaiClient := openai.NewGPT(config.GlobalConfig.OpenAI.APIKey)

	postgresConn, err := postgres.NewConnectionPool(ctx, config.GlobalConfig.Postgres)
	if err != nil {
		log.Panic(err)
	}

	telegramBot, err := bot.NewBot(config.GlobalConfig.BOT.TOKEN, openaiClient, postgresConn)
	if err != nil {
		log.Panic(err)
	}

	telegramBot.Start()

	return
}
