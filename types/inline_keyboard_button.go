package types

// InlineKeyboardButton Represents one button of an inline keyboard
// You must use exactly one of the optional fields.
type InlineKeyboardButton struct {
	CallbackData string `json:"callback_data,omitempty"`
	CallbackGame *CallbackGame `json:"callback_game,omitempty"`
	LoginUrl *LoginUrl `json:"login_url,omitempty"`
	Pay bool `json:"pay,omitempty"`
	SwitchInlineQuery string `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`
	Text string `json:"text"`
	Url string `json:"url,omitempty"`
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}
