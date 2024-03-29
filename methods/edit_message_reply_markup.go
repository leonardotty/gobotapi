// Code AutoGenerated; DO NOT EDIT.

package methods

import (
	"encoding/json"
	"github.com/Squirrel-Network/gobotapi/types"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

// EditMessageReplyMarkup Use this method to edit only the reply markup of messages
// On success, if the edited message is not an inline message, the edited Message is returned, otherwise True is returned.
type EditMessageReplyMarkup struct {
	ChatID int64 `json:"chat_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
	MessageID int64 `json:"message_id,omitempty"`
	ReplyMarkup *types.InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (entity *EditMessageReplyMarkup) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (EditMessageReplyMarkup) MethodName() string {
	return "editMessageReplyMarkup"
}

func (EditMessageReplyMarkup) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x0 struct {
		Result bool `json:"result"`
	}
	_ = json.Unmarshal(response, &x0)
	if x0.Result {
		result := rawTypes.Result {
			Kind: types.TypeBoolean,
			Result: true,
		}
		return &result, nil
	} else {
		var x1 struct {
			Result types.Message `json:"result"`
		}
		err := json.Unmarshal(response, &x1)
		if err != nil {
			return nil, err
		}
		result := rawTypes.Result {
			Kind: types.TypeMessage,
			Result: x1.Result,
		}
		return &result, nil
	}
}
