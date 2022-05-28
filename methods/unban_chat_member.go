package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

// UnbanChatMember Use this method to unban a previously banned user in a supergroup or channel
// The user will not return to the group or channel automatically, but will be able to join via link, etc
// The bot must be an administrator for this to work
// By default, this method guarantees that after the call the user is not a member of the chat, but will be able to join it
// So if the user is a member of the chat they will also be removed from the chat
// If you don't want this, use the parameter only_if_banned
// Returns True on success.
type UnbanChatMember struct {
	ChatId int64 `json:"chat_id"`
	OnlyIfBanned bool `json:"only_if_banned,omitempty"`
	UserId int64 `json:"user_id"`
}

func (entity *UnbanChatMember) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (UnbanChatMember) MethodName() string {
	return "unbanChatMember"
}

func (UnbanChatMember) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result bool `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeBoolean,
		Result: x1.Result,
	}
	return &result, nil
}
