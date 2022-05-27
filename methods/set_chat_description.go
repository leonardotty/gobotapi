package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

type SetChatDescription struct {
	ChatId int64 `json:"chat_id"`
	Description string `json:"description,omitempty"`
}

func (entity *SetChatDescription) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (SetChatDescription) MethodName() string {
	return "setChatDescription"
}

func (SetChatDescription) ParseResult(response []byte) (*rawTypes.Result, error) {
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