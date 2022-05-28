package types

// ChatJoinRequest Represents a join request sent to a chat.
type ChatJoinRequest struct {
	Bio string `json:"bio,omitempty"`
	Chat Chat `json:"chat"`
	Date int64 `json:"date"`
	From User `json:"from"`
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
}
