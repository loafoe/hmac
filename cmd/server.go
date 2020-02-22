/*
Copyright © 2020 Andy Lo-A-Foe <andy.lo-a-foe@philips.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"errors"
	"log"
	"os"

	"github.com/philips-labs/hmac/alerts"
	"github.com/philips-labs/hmac/router"
	"github.com/spf13/cobra"
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
			log.Fatal(errors.New("you need to provide a token"))
		}
		if port == "" {
			log.Fatal(errors.New("you need to provide a port to listen on"))
		}
		storer, err := alerts.NewPGStorer()
		if err != nil {
			log.Fatal(err)
		}
		e := router.New(router.Config{
			Storer: storer,
			Token:  token,
		})
		e.Logger.Fatal(e.Start(":" + port))
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
	serverCmd.Flags().StringP("port", "p", "8080", "Port to listen on")
	serverCmd.Flags().StringP("token", "t", os.Getenv("TOKEN"), "Token for /web/hook/alerts/:token")
}
