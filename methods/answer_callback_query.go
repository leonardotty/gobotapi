package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

// AnswerCallbackQuery Use this method to send answers to callback queries sent from inline keyboards
// The answer will be displayed to the user as a notification at the top of the chat screen or as an alert
// On success, True is returned.
type AnswerCallbackQuery struct {
	CacheTime int `json:"cache_time,omitempty"`
	CallbackQueryId string `json:"callback_query_id"`
	ShowAlert bool `json:"show_alert,omitempty"`
	Text string `json:"text,omitempty"`
	Url string `json:"url,omitempty"`
}

func (entity *AnswerCallbackQuery) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (AnswerCallbackQuery) MethodName() string {
	return "answerCallbackQuery"
}

func (AnswerCallbackQuery) ParseResult(response []byte) (*rawTypes.Result, error) {
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
