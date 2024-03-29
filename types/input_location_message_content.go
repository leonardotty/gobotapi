// Code AutoGenerated; DO NOT EDIT.

package types

// InputLocationMessageContent Represents the content of a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {
	Heading int `json:"heading,omitempty"`
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
	Latitude float64 `json:"latitude"`
	LivePeriod int `json:"live_period,omitempty"`
	Longitude float64 `json:"longitude"`
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
}
