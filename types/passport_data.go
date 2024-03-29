// Code AutoGenerated; DO NOT EDIT.

package types

// PassportData Contains information about Telegram Passport data shared with the bot by the user.
type PassportData struct {
	Credentials EncryptedCredentials `json:"credentials"`
	Data []EncryptedPassportElement `json:"data,omitempty"`
}
