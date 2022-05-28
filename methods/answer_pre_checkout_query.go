package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

// AnswerPreCheckoutQuery Once the user has confirmed their payment and shipping details, the Bot API sends the final confirmation in the form of an Update with the field pre_checkout_query
// Use this method to respond to such pre-checkout queries
// On success, True is returned
// Note: The Bot API must receive an answer within 10 seconds after the pre-checkout query was sent.
type AnswerPreCheckoutQuery struct {
	ErrorMessage string `json:"error_message,omitempty"`
	Ok bool `json:"ok"`
	PreCheckoutQueryId string `json:"pre_checkout_query_id"`
}

func (entity *AnswerPreCheckoutQuery) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (AnswerPreCheckoutQuery) MethodName() string {
	return "answerPreCheckoutQuery"
}

func (AnswerPreCheckoutQuery) ParseResult(response []byte) (*rawTypes.Result, error) {
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
