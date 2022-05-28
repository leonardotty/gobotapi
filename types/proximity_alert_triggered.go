package types

// ProximityAlertTriggered Represents the content of a service message, sent whenever a user in the chat triggers a proximity alert set by another user.
type ProximityAlertTriggered struct {
	Distance int `json:"distance"`
	Traveler User `json:"traveler"`
	Watcher User `json:"watcher"`
}
