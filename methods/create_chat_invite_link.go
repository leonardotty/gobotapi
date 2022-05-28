package methods

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "github.com/Squirrel-Network/gobotapi/types"
import "encoding/json"

// CreateChatInviteLink Use this method to create an additional invite link for a chat
// The bot must be an administrator in the chat for this to work and must have the appropriate administrator rights
// The link can be revoked using the method revokeChatInviteLink
// Returns the new invite link as ChatInviteLink object.
type CreateChatInviteLink struct {
	ChatId int64 `json:"chat_id"`
	CreatesJoinRequest bool `json:"creates_join_request,omitempty"`
	ExpireDate int64 `json:"expire_date,omitempty"`
	MemberLimit int `json:"member_limit,omitempty"`
	Name string `json:"name,omitempty"`
}

func (entity *CreateChatInviteLink) Files() map[string]rawTypes.InputFile {
	return map[string]rawTypes.InputFile{}
}

func (CreateChatInviteLink) MethodName() string {
	return "createChatInviteLink"
}

func (CreateChatInviteLink) ParseResult(response []byte) (*rawTypes.Result, error) {
	var x1 struct {
		Result types.ChatInviteLink `json:"result"`
	}
	err := json.Unmarshal(response, &x1)
	if err != nil {
		return nil, err
	}
	result := rawTypes.Result {
		Kind: types.TypeChatInviteLink,
		Result: x1.Result,
	}
	return &result, nil
}
