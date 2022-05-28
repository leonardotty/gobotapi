package types

// InlineQuery Represents an incoming inline query
// When the user sends an empty query, your bot could return some default or trending results.
type InlineQuery struct {
	ChatType string `json:"chat_type,omitempty"`
	From User `json:"from"`
	Id string `json:"id"`
	Location *Location `json:"location,omitempty"`
	Offset string `json:"offset"`
	Query string `json:"query"`
}
