package cmd

import (
	"errors"
	"log"

	"github.com/philips-labs/hmac/alerts"
	"github.com/philips-labs/hmac/router"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start a hmac server",
	Long:  `Starts a hmac server which exposes a webhook endpoint to capture alerts`,
	Run: func(cmd *cobra.Command, args []string) {
		token, _ := cmd.Flags().GetString("token")
		port, _ := cmd.Flags().GetString("port")

		if token == "" {
			token = viper.GetString("token")
		}
		if token == "" {
			log.Fatal(errors.New("you need to provide a token"))
		}
		if port == "" {
			log.Fatal(errors.New("you need to provide a port to listen on"))
		}
		storer, err := alerts.NewPGStorer()
		if err != nil {
			log.Fatal(err)
		}
		e, err := router.New(router.Config{
			Storer: storer,
			Token:  token,
		})
		if err != nil {
			log.Fatal(err)
		}
		e.Logger.Fatal(e.Start(":" + port))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringP("port", "p", "8080", "Port to listen on")
	serverCmd.Flags().StringP("token", "t", "", "Token for /webhook/alerts/:token")
}
