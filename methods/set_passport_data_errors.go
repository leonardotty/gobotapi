package methods

import "github.com/Squirrel-Network/gobotapi/types"
import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "encoding/json"

// SetPassportDataErrors Informs a user that some of the Telegram Passport elements they provided contains errors
// The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change)
// Returns True on success.
// Use this if the data submitted by the user doesn't satisfy the standards your service requires for any reason
// For example, if a birthday date seems invalid, a submitted document is blurry, a scan shows evidence of tampering, etc
// Supply some details in the error message to make sure the user knows how to correct the issues.
type SetPassportDataErrors struct {
	Errors []types.PassportElementError `json:"errors,omitempty"`
	UserId int64 `json:"user_id"`
}

func (entity *SetPassportDataErrors) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (SetPassportDataErrors) MethodName() string {
	return "setPassportDataErrors"
}

func (SetPassportDataErrors) ParseResult(response []byte) (*rawTypes.Result, error) {
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
