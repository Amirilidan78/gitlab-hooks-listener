package telegram

import "gitlab-hooks-listener/pkg/http_client"

const baseUrl = "https://api.telegram.org/"

type Telegram interface {
	SendMessage(message string) error
}

type telegram struct {
	hc    http_client.HttpClient
	token string
	chat  string
}

func (t *telegram) SendMessage(message string) error {

	res := SendMessageTelegramResponse{}

	url := baseUrl + t.token + "/sendMessage?chat_id=" + t.chat + "&text=" + message + "&parse_mode=html"

	err := t.hc.SimpleGet(url, &res)

	return err
}
