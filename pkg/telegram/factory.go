package telegram

type MessageFactory interface {
	NewLine() string
	Text(text string) string
	Bold(text string) string
	Italic(text string) string
	Strike(text string) string
	Underline(text string) string
	Spoiler(text string) string
	Link(text string, url string) string
	Code(text string) string
	Copy(text string) string
	Divider() string
}

func (t *telegram) NewLine() string {
	return " \n"
}

func (t *telegram) Text(text string) string {
	return text + " "
}

func (t *telegram) Bold(text string) string {
	return "<b>" + text + "</b> "
}

func (t *telegram) Italic(text string) string {
	return "<i>" + text + "</i> "
}

func (t *telegram) Strike(text string) string {
	return "<s>" + text + "</s> "
}

func (t *telegram) Underline(text string) string {
	return "<u>" + text + "</u> "
}

func (t *telegram) Spoiler(text string) string {
	return "<tg-spoiler>" + text + "</tg-spoiler> "
}

func (t *telegram) Link(text string, url string) string {
	return "<a href='" + url + "'>" + text + "</a> "
}

func (t *telegram) Code(text string) string {
	return "<pre'>" + text + "</pre> "
}

func (t *telegram) Copy(text string) string {
	return "<code>" + text + "</code> "
}

func (t *telegram) Divider() string {
	return "<pre>-------------------------------</pre>"
}
