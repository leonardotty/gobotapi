// Code AutoGenerated; DO NOT EDIT.

package types

import (
	"encoding/json"
	"fmt"
	rawTypes "github.com/Squirrel-Network/gobotapi/types/raw"
)

// InputMediaAnimation Represents an animation file (GIF or H.264/MPEG-4 AVC video without sound) to be sent.
type InputMediaAnimation struct {
	Caption string `json:"caption,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	Duration int `json:"duration,omitempty"`
	Height int `json:"height,omitempty"`
	Media rawTypes.InputFile `json:"media,omitempty"`
	ParseMode string `json:"parse_mode,omitempty"`
	Thumb rawTypes.InputFile `json:"thumb,omitempty"`
	Width int64 `json:"width,omitempty"`
}

func (entity *InputMediaAnimation) Files() map[string]rawTypes.InputFile {
	files := make(map[string]rawTypes.InputFile)
	switch entity.Media.(type) {
		case InputFile:
			files["animation"] = entity.Media
	}
	switch entity.Thumb.(type) {
		case InputFile:
			files["thumb"] = entity.Thumb
	}
	return files
}

func (entity *InputMediaAnimation) SetAttachment(attach string) {
	entity.Media = InputPath(fmt.Sprintf("attach://%s", attach))
}

func (entity *InputMediaAnimation) SetAttachmentThumb(attach string) {
	entity.Thumb = InputPath(fmt.Sprintf("attach://%s", attach))
}

func (entity InputMediaAnimation) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type string `json:"type"`
		Media rawTypes.InputFile `json:"media,omitempty"`
		Thumb rawTypes.InputFile `json:"thumb,omitempty"`
		Caption string `json:"caption,omitempty"`
		ParseMode string `json:"parse_mode,omitempty"`
		CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
		Width int64 `json:"width,omitempty"`
		Height int `json:"height,omitempty"`
		Duration int `json:"duration,omitempty"`
	} {
		Type: "animation",
		Media: entity.Media,
		Thumb: entity.Thumb,
		Caption: entity.Caption,
		ParseMode: entity.ParseMode,
		CaptionEntities: entity.CaptionEntities,
		Width: entity.Width,
		Height: entity.Height,
		Duration: entity.Duration,
	}
	return json.Marshal(alias)
}
