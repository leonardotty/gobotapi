package types

import "encoding/json"
import "fmt"

// InlineQueryResultCachedSticker Represents a link to a sticker stored on the Telegram servers
// By default, this sticker will be sent by the user
// Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.
// Note: This will only work in Telegram versions released after 9 April, 2016 for static stickers and after 06 July, 2019 for animated stickers
// Older clients will ignore them.
type InlineQueryResultCachedSticker struct {
	Id string `json:"id"`
	InputMessageContent interface{} `json:"input_message_content,omitempty"`
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	StickerFileId string `json:"sticker_file_id"`
}

func (entity InlineQueryResultCachedSticker) MarshalJSON() ([]byte, error) {
	alias := struct {
		Type string `json:"type"`
		Id string `json:"id"`
		StickerFileId string `json:"sticker_file_id"`
		ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
		InputMessageContent interface{} `json:"input_message_content,omitempty"`
	} {
		Type: "sticker",
		Id: entity.Id,
		StickerFileId: entity.StickerFileId,
		ReplyMarkup: entity.ReplyMarkup,
		InputMessageContent: entity.InputMessageContent,
	}
	if entity.InputMessageContent != nil {
		switch entity.InputMessageContent.(type) {
			case InputTextMessageContent, InputLocationMessageContent, InputVenueMessageContent, InputContactMessageContent, InputInvoiceMessageContent:
				break
			default:
				return nil, fmt.Errorf("input_message_content: unknown type: %T", entity.InputMessageContent)
		}
	}
	return json.Marshal(alias)
}
