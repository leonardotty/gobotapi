package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

// GetStickerSet Use this method to get a sticker set
// On success, a StickerSet object is returned.
type GetStickerSet struct {
	Name string `json:"name"`
}

func (entity *GetStickerSet) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (GetStickerSet) MethodName() string {
	return "getStickerSet"
}

func (GetStickerSet) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.StickerSet `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeStickerSet,
		Result: x1.Result,
	}
	return &result, nil
}
