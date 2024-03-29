// Code AutoGenerated; DO NOT EDIT.

package types

// PassportElementErrorFiles Represents an issue with a list of scans
// The error is considered resolved when the list of files containing the scans changes.
type PassportElementErrorFiles struct {
	FileHashes []string `json:"file_hashes,omitempty"`
	Message string `json:"message"`
	Source string `json:"source"`
	Type string `json:"type"`
}
