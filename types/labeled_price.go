package types

// LabeledPrice Represents a portion of the price for goods or services.
type LabeledPrice struct {
	Amount int `json:"amount"`
	Label string `json:"label"`
}
