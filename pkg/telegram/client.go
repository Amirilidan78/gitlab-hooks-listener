package telegram

import (
	"gitlab-hooks-listener/pkg/http_client"
)

const baseUrl = "https://api.telegram.org/"

type Telegram interface {
	MessageFactory
	Message
}

type telegram struct {
	hc    http_client.HttpClient
	token string
	chat  string
}

func GetTelegramClient(token string, chat string) Telegram {

	return &telegram{hc: http_client.NewHttpClient(), token: token, chat: chat}
}
