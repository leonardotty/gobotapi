package types

import rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
import "fmt"
import "encoding/json"

type InputMediaAudio struct {
	Caption string `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Duration int `json:"duration,omitempty"`
	Media rawTypes.InputFile `json:"media,omitempty"`
	ParseMode string `json:"parse_mode,omitempty"`
	Performer string `json:"performer,omitempty"`
	Thumb rawTypes.InputFile `json:"thumb,omitempty"`
	Title string `json:"title,omitempty"`
}

func (entity *InputMediaAudio) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Media.(type) {
		case InputFile:
			files["audio"] = entity.Media
	}
	switch entity.Thumb.(type) {
		case InputFile:
			files["thumb"] = entity.Thumb
	}
	return files
}

func (entity *InputMediaAudio) SetAttachment(attach string) {
	entity.Media = InputURL(fmt.Sprintf("attach://%s", attach))
}

func (entity *InputMediaAudio) SetAttachmentThumb(attach string) {
	entity.Thumb = InputURL(fmt.Sprintf("attach://%s", attach))
}

func (entity InputMediaAudio) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type string `json:"type"`
		Media rawTypes.InputFile `json:"media,omitempty"`
		Thumb rawTypes.InputFile `json:"thumb,omitempty"`
		Caption string `json:"caption,omitempty"`
		ParseMode string `json:"parse_mode,omitempty"`
		CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
		Duration int `json:"duration,omitempty"`
		Performer string `json:"performer,omitempty"`
		Title string `json:"title,omitempty"`
	} {
		Type: "audio",
		Media: entity.Media,
		Thumb: entity.Thumb,
		Caption: entity.Caption,
		ParseMode: entity.ParseMode,
		CaptionEntities: entity.CaptionEntities,
		Duration: entity.Duration,
		Performer: entity.Performer,
		Title: entity.Title,
	}
	return json.Marshal(alias)
}