package cmd

import (
	"alfred/pkg/ascii"
	"alfred/pkg/host"
	"alfred/pkg/notification"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"gopkg.in/macaron.v1"
)

const (
	TELEGRAM_TOKEN = "tg-token"
	DISABLE_NOTIFY = "disable-notify"
	SERVER_PORT    = "port"
	PROD           = "prod"
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
		banner.HeadBanner("ASSISTANT", "START")

		m := macaron.Classic()

		m.Map(l.Named("server"))

		m.Get("/getMyAddress", func(l *zap.Logger) string {
			return host.GetIPAddress()
		})

		port, err := cmd.Flags().GetInt(SERVER_PORT)

		if err != nil {
			l.Error(err.Error())
		}

		m.Run(port)
	},
}
