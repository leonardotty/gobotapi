package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "encoding/json"
import "fmt"
import "github.com/Squirrel-Network/gobotapi/types"

type SendLocation struct {
	AllowSendingWithoutReply bool `json:"allow_sending_without_reply,omitempty"`
	ChatId int64 `json:"chat_id"`
	DisableNotification bool `json:"disable_notification,omitempty"`
	Heading int `json:"heading,omitempty"`
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
	Latitude float64 `json:"latitude"`
	LivePeriod int `json:"live_period,omitempty"`
	Longitude float64 `json:"longitude"`
	ProtectContent bool `json:"protect_content,omitempty"`
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
	ReplyToMessageId int64 `json:"reply_to_message_id,omitempty"`
}

func (entity *SendLocation) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (entity SendLocation) MarshalJSON() ([]byte, error) {
	if entity.ReplyMarkup != nil {
		switch entity.ReplyMarkup.(type) {
			case *types.InlineKeyboardMarkup, *types.ReplyKeyboardMarkup, *types.ReplyKeyboardRemove, *types.ForceReply:
				break
			default:
				return nil, fmt.Errorf("reply_markup: unknown type: %T", entity.ReplyMarkup)
		}
	}
	type x0 SendLocation
	return json.Marshal((x0)(entity))
}

func (SendLocation) MethodName() string {
	return "sendLocation"
}

func (SendLocation) ParseResult(response []byte) (*rawTypes.Result, error) {
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