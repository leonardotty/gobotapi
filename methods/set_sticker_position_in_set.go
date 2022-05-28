package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

// SetStickerPositionInSet Use this method to move a sticker in a set created by the bot to a specific position
// Returns True on success.
type SetStickerPositionInSet struct {
	Position int `json:"position"`
	Sticker string `json:"sticker"`
}

func (entity *SetStickerPositionInSet) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (SetStickerPositionInSet) MethodName() string {
	return "setStickerPositionInSet"
}

func (SetStickerPositionInSet) ParseResult(response []byte) (*rawTypes.Result, error) {
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
