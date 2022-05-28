package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

// GetChatMember Use this method to get information about a member of a chat
// Returns a ChatMember object on success.
type GetChatMember struct {
	ChatId int64 `json:"chat_id"`
	UserId int64 `json:"user_id"`
}

func (entity *GetChatMember) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (GetChatMember) MethodName() string {
	return "getChatMember"
}

func (GetChatMember) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.ChatMember `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeChatMember,
		Result: x1.Result,
	}
	return &result, nil
}
