package telegram

type SendMessageTelegramResponse struct {
	Ok   string `json:"ok"`
	Date string `json:"date"`
	Text string `json:"text"`
}

type SendPhotoTelegramResponse struct {
	Ok   string `json:"ok"`
	Date string `json:"date"`
	Text string `json:"text"`
}
