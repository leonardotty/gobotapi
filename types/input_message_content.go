package types

// InputMessageContent Represents the content of a message to be sent as a result of an inline query
// Telegram clients currently support the following 5 types:
//  - InputTextMessageContent
//  - InputLocationMessageContent
//  - InputVenueMessageContent
//  - InputContactMessageContent
//  - InputInvoiceMessageContent
type InputMessageContent interface{}
