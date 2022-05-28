package types

// ChatAdministratorRights Represents the rights of an administrator in a chat.
type ChatAdministratorRights struct {
	CanChangeInfo bool `json:"can_change_info"`
	CanDeleteMessages bool `json:"can_delete_messages"`
	CanEditMessages bool `json:"can_edit_messages,omitempty"`
	CanInviteUsers bool `json:"can_invite_users"`
	CanManageChat bool `json:"can_manage_chat"`
	CanManageVideoChats bool `json:"can_manage_video_chats"`
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
	CanPostMessages bool `json:"can_post_messages,omitempty"`
	CanPromoteMembers bool `json:"can_promote_members"`
	CanRestrictMembers bool `json:"can_restrict_members"`
	IsAnonymous bool `json:"is_anonymous"`
}
