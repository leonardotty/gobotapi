package types

// KeyboardButton Represents one button of the reply keyboard
// For simple text buttons String can be used instead of this object to specify text of the button
// Optional fields web_app, request_contact, request_location, and request_poll are mutually exclusive.
// Note: request_contact and request_location options will only work in Telegram versions released after 9 April, 2016
// Older clients will display unsupported message.Note: request_poll option will only work in Telegram versions released after 23 January, 2020
// Older clients will display unsupported message.Note: web_app option will only work in Telegram versions released after 16 April, 2022
// Older clients will display unsupported message.
type KeyboardButton struct {
	RequestContact bool `json:"request_contact,omitempty"`
	RequestLocation bool `json:"request_location,omitempty"`
	RequestPoll *KeyboardButtonPollType `json:"request_poll,omitempty"`
	Text string `json:"text"`
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}
