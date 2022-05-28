package types

// PassportElementErrorTranslationFiles Represents an issue with the translated version of a document
// The error is considered resolved when a file with the document translation change.
type PassportElementErrorTranslationFiles struct {
	FileHashes []string `json:"file_hashes,omitempty"`
	Message string `json:"message"`
	Source string `json:"source"`
	Type string `json:"type"`
}
