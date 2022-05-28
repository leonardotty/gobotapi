package types

// PassportElementErrorTranslationFile Represents an issue with one of the files that constitute the translation of a document
// The error is considered resolved when the file changes.
type PassportElementErrorTranslationFile struct {
	FileHash string `json:"file_hash"`
	Message string `json:"message"`
	Source string `json:"source"`
	Type string `json:"type"`
}
