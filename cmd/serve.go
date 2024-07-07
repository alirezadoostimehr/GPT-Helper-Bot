package cmd

import (
	"fmt"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/bot"
	"github.com/alirezadoostimehr/GPT-Helper-Bot/internal/config"

	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start serving",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serve called")
		if err := config.Load(); err != nil {
			panic(err)
		}

		bot, err := bot.NewBot(config.GlobalConfig.BOT.TOKEN)
		if err != nil {
			panic(err)
		}

		bot.Start()

		return
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
