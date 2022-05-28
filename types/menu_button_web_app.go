package types

// MenuButtonWebApp Represents a menu button, which launches a Web App.
type MenuButtonWebApp struct {
	Text string `json:"text"`
	Type string `json:"type"`
	WebApp WebAppInfo `json:"web_app"`
}
