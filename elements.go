package cards

import (
	"errors"
	"fmt"
)

// TextBlock is textblock element
type TextBlock struct {
	Type                string `json:"type"` // required
	Text                string `json:"text"` // required
	Color               string `json:"color,omitempty"`
	FontType            string `json:"fontType,omitempty"` // FIXME this is a special type in a.c.
	HorizontalAlignment string `json:"horizontalAlignment,omitempty"`
	IsSublte            bool   `json:"isSublte,omitempty"`
	MaxLines            int64  `json:"maxLines,omitempty"`
	Size                string `json:"size,omitempty"`
	Weight              string `json:"weight,omitempty"`
	Wrap                bool   `json:"wrap,omitempty"`
	// TODO add inherited
}

func (n TextBlock) validate() error {
	if n.Type != TextBlockType {
		return fmt.Errorf("type must be %s", TextBlockType)
	}
	if n.Text == "" {
		return errors.New("text is required")
	}
	return nil
}
