package types

// SentWebAppMessage Contains information about an inline message sent by a Web App on behalf of a user.
type SentWebAppMessage struct {
	InlineMessageId string `json:"inline_message_id,omitempty"`
}
