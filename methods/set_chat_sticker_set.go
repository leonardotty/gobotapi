package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

// SetChatStickerSet Use this method to set a new group sticker set for a supergroup
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights
// Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method
// Returns True on success.
type SetChatStickerSet struct {
	ChatId int64 `json:"chat_id"`
	StickerSetName string `json:"sticker_set_name"`
}

func (entity *SetChatStickerSet) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (SetChatStickerSet) MethodName() string {
	return "setChatStickerSet"
}

func (SetChatStickerSet) ParseResult(response []byte) (*rawTypes.Result, error) {
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
