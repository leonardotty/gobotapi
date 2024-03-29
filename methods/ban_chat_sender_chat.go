// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

// BanChatSenderChat Use this method to ban a channel chat in a supergroup or a channel
// Until the chat is unbanned, the owner of the banned chat won't be able to send messages on behalf of any of their channels
// The bot must be an administrator in the supergroup or channel for this to work and must have the appropriate administrator rights
// Returns True on success.
type BanChatSenderChat struct {
	ChatID int64 `json:"chat_id"`
	SenderChatID int64 `json:"sender_chat_id"`
}

func (entity *BanChatSenderChat) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (BanChatSenderChat) MethodName() string {
	return "banChatSenderChat"
}

func (BanChatSenderChat) ParseResult(response []byte) (*rawTypes.Result, error) {
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
