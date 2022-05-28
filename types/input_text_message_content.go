package types

// InputTextMessageContent Represents the content of a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
	Entities []MessageEntity `json:"entities,omitempty"`
	MessageText string `json:"message_text"`
	ParseMode string `json:"parse_mode,omitempty"`
}
