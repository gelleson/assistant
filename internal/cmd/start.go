/*
 * MIT License
 *
 * Copyright (c) 2020 gelleson
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package cmd

import (
	"github.com/gelleson/assistant/internal/ascii"
	"github.com/gelleson/assistant/internal/controller"
	"github.com/gelleson/assistant/internal/host"
	"github.com/gelleson/assistant/internal/notification"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"gopkg.in/macaron.v1"
)

const (
	// TELEGRAM_TOKEN Telegram token to grant access to the bot
	TELEGRAM_TOKEN = "tg-token"
	// DISABLE_NOTIFY disable send message with address after start the server
	DISABLE_NOTIFY = "disable-notify"
	// SERVER_PORT port of the server
	SERVER_PORT = "port"
	// PROD if want to run in production mode
	PROD = "prod"
)

var (
	provider *notification.Telegram
	l        *zap.Logger
)

func init() {
	start.Flags().Bool(DISABLE_NOTIFY, false, "disable-notify")
	start.Flags().String(TELEGRAM_TOKEN, "", "XXXXXXXX:XXXXXXXXXXXXXXXXXXX")
	start.Flags().IntP(SERVER_PORT, "p", 4000, "5000")
	start.Flags().Bool(PROD, false, "true")
}

var start = &cobra.Command{
	Use:   "start",
	Short: "command to run web server",
	PreRun: func(cmd *cobra.Command, args []string) {
		disableNotify, err := cmd.Flags().GetBool(DISABLE_NOTIFY)
		l, _ = zap.NewDevelopment(zap.Development())

		token, err := cmd.Flags().GetString(TELEGRAM_TOKEN)

		if err != nil {
			l.Error(err.Error())
		}

		provider = notification.NewTelegramProvider(token)

		if !disableNotify {
			err = provider.Notify("301990443", host.GetIPAddress())

			if err != nil {
				l.Error(err.Error())
			}
		}

		isProd, err := cmd.Flags().GetBool(PROD)

		if err != nil {
			l.Error(err.Error())
		}

		if isProd {
			macaron.Env = macaron.PROD
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		banner := ascii.NewBannerWithColor(ascii.BANNER_COLOR_RED)
		banner.Print("ASSISTANT", "START")

		m := macaron.Classic()

		m.Map(l.Named("server"))

		m.Get("/getMyAddress", func(l *zap.Logger) string {
			return host.GetIPAddress()
		})

		m.Group("/monitoring", func() {
			m.Get("/metric", controller.GetSnapshot)
		})

		port, err := cmd.Flags().GetInt(SERVER_PORT)

		if err != nil {
			l.Error(err.Error())
		}

		m.Run(port)
	},
}
