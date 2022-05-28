package types

// MessageEntity Represents one special entity in a text message
// For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	Language string `json:"language,omitempty"`
	Length int `json:"length"`
	Offset int `json:"offset"`
	Type string `json:"type"`
	Url string `json:"url,omitempty"`
	User *User `json:"user,omitempty"`
}
