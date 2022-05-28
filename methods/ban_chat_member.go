package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

// BanChatMember Use this method to ban a user in a group, a supergroup or a channel
// In the case of supergroups and channels, the user will not be able to return to the chat on their own using invite links, etc., unless unbanned first
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights
// Returns True on success.
type BanChatMember struct {
	ChatId int64 `json:"chat_id"`
	RevokeMessages bool `json:"revoke_messages,omitempty"`
	UntilDate int64 `json:"until_date,omitempty"`
	UserId int64 `json:"user_id"`
}

func (entity *BanChatMember) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (BanChatMember) MethodName() string {
	return "banChatMember"
}

func (BanChatMember) ParseResult(response []byte) (*rawTypes.Result, error) {
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
