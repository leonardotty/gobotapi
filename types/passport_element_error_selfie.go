// Code AutoGenerated; DO NOT EDIT.

package types

// PassportElementErrorSelfie Represents an issue with the selfie with a document
// The error is considered resolved when the file with the selfie changes.
type PassportElementErrorSelfie struct {
	FileHash string `json:"file_hash"`
	Message string `json:"message"`
	Source string `json:"source"`
	Type string `json:"type"`
}
