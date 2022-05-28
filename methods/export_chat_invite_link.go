package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

// ExportChatInviteLink Use this method to generate a new primary invite link for a chat; any previously generated primary link is revoked
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights
// Returns the new invite link as String on success.
type ExportChatInviteLink struct {
	ChatId int64 `json:"chat_id"`
}

func (entity *ExportChatInviteLink) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (ExportChatInviteLink) MethodName() string {
	return "exportChatInviteLink"
}

func (ExportChatInviteLink) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result string `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeString,
		Result: x1.Result,
	}
	return &result, nil
}
