package types

// ChosenInlineResult Represents a result of an inline query that was chosen by the user and sent to their chat partner.
// Note: It is necessary to enable inline feedback via @Botfather in order to receive these objects in updates.
type ChosenInlineResult struct {
	From User `json:"from"`
	InlineMessageId string `json:"inline_message_id,omitempty"`
	Location *Location `json:"location,omitempty"`
	Query string `json:"query"`
	ResultId string `json:"result_id"`
}
