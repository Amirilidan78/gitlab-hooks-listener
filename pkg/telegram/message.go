package telegram

type Message interface {
	SendMessage(message string) error
	SendPhoto(photo string, caption string) error
}

func (t *telegram) SendMessage(message string) error {

	res := SendMessageTelegramResponse{}

	url := baseUrl + t.token + "/sendMessage"

	body := map[string]string{
		"chat_id":    t.chat,
		"text":       message,
		"parse_mode": "html",
	}

	err := t.hc.SimpleGet(url, body, &res)

	return err
}

func (t *telegram) SendPhoto(photo string, caption string) error {

	//res := SendMessageTelegramResponse{}
	//
	//url := baseUrl + t.token + "/sendMessage?chat_id=" + t.chat + "&photo=" + photo + "&caption=" + caption + "&parse_mode=html"
	//
	//err := t.hc.SimpleGet(url, &res)

	return nil
}
