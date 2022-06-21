package telegram

import (
	"gitlab-hooks-listener/config"
	"gitlab-hooks-listener/pkg/http_client"
	"sync"
)

var lock = &sync.Mutex{}

var singleton *telegram

func GetTelegramClient() Telegram {

	if singleton == nil {

		lock.Lock()
		defer lock.Unlock()

		c := config.NewConfig()
		hc := http_client.NewHttpClient()

		token := c.GetString("telegram.token")
		chat := c.GetString("telegram.chat")

		singleton = &telegram{hc: hc, token: token, chat: chat}
	}

	return singleton
}
