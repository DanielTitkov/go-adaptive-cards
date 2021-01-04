package cards

import (
	"errors"
	"fmt"
)

// InputText is Input.Text type
type InputText struct {
	Type         string `json:"type"` // required
	ID           string `json:"id"`   // required
	IsMultiline  bool   `json:"isMultiline,omitempty"`
	MaxLength    int64  `json:"maxLength,omitempty"`
	Placeholder  string `json:"placeholder,omitempty"`
	Regex        string `json:"regex,omitempty"`
	Style        string `json:"style,omitempty"`
	InlineAction Node   `json:"inlineAction,omitempty"` // FIXME
	Value        string `json:"value,omitempty"`
	// inherited
	ErrorMessage string            `json:"errorMessage,omitempty"`
	IsRequired   bool              `json:"isRequired,omitempty"`
	Label        string            `json:"label,omitempty"`
	Height       string            `json:"height,omitempty"`
	Separator    bool              `json:"separator,omitempty"`
	Spacing      string            `json:"spacing,omitempty"`
	IsVisible    bool              `json:"isVisible,omitempty"`
	Requires     map[string]string `json:"requires,omitempty"`
}

func (n InputText) validate() error {
	if n.Type != InputTextType {
		return fmt.Errorf("type must be %s", InputTextType)
	}
	if n.ID == "" {
		return errors.New("id must be present")
	}
	return nil
}
