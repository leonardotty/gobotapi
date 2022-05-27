package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

type BanChatSenderChat struct {
	ChatId int64 `json:"chat_id"`
	SenderChatId int64 `json:"sender_chat_id"`
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