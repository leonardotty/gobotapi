package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

// GetChatAdministrators Use this method to get a list of administrators in a chat
// On success, returns an Array of ChatMember objects that contains information about all chat administrators except other bots
// If the chat is a group or a supergroup and no administrators were appointed, only the creator will be returned.
type GetChatAdministrators struct {
	ChatId int64 `json:"chat_id"`
}

func (entity *GetChatAdministrators) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (GetChatAdministrators) MethodName() string {
	return "getChatAdministrators"
}

func (GetChatAdministrators) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result []types.ChatMember `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeArrayOfChatMember,
		Result: x1.Result,
	}
	return &result, nil
}
