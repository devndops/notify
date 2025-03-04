package models

import "time"

// Theme defines styling for emails.
type Theme struct {
	Name        string `json:"name"`
	CSS         string `json:"css"`
	ButtonColor string `json:"buttonColor"`
}

// Preset themes.
func DefaultTheme() Theme {
	return Theme{
		Name:        "default",
		CSS:         "body { font-family: Arial, sans-serif; }",
		ButtonColor: "#3869D4",
	}
}

func ModernTheme() Theme {
	return Theme{
		Name:        "modern",
		CSS:         "body { font-family: 'Segoe UI', sans-serif; }",
		ButtonColor: "#e74c3c",
	}
}

// BaseEmailData contains common fields for all emails.
type BaseEmailData struct {
	AppName        string `json:"appName"`        // e.g. "MyApp"
	AppLink        string `json:"appLink"`        // e.g. "https://myapp.com"
	Logo           string `json:"logo"`           // Optional logo URL
	Year           int    `json:"year"`           // Defaults to current year
	RecipientEmail string `json:"recipientEmail"` // Must be provided
	RecipientName  string `json:"recipientName"`  // Must be provided
	Theme          Theme  `json:"theme"`          // If omitted or ButtonColor is empty, a preset is used.
}

// FillBaseData fills in missing common fields.
func FillBaseData(data *BaseEmailData) {
	if data.Year == 0 {
		data.Year = time.Now().Year()
	}
	if data.Theme.Name == "" {
		data.Theme = DefaultTheme()
	} else if data.Theme.ButtonColor == "" {
		switch data.Theme.Name {
		case "modern":
			data.Theme.ButtonColor = ModernTheme().ButtonColor
		default:
			data.Theme.ButtonColor = DefaultTheme().ButtonColor
		}
	}
}
