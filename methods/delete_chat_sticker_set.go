package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

// DeleteChatStickerSet Use this method to delete a group sticker set from a supergroup
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights
// Use the field can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method
// Returns True on success.
type DeleteChatStickerSet struct {
	ChatId int64 `json:"chat_id"`
}

func (entity *DeleteChatStickerSet) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (DeleteChatStickerSet) MethodName() string {
	return "deleteChatStickerSet"
}

func (DeleteChatStickerSet) ParseResult(response []byte) (*rawTypes.Result, error) {
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
