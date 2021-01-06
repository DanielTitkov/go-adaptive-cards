package cards

import (
	"errors"
)

// TextBlock is textblock element
type TextBlock struct {
	Type                string `json:"type"` // required
	Text                string `json:"text"` // required
	Color               string `json:"color,omitempty"`
	FontType            string `json:"fontType,omitempty"` // FIXME this is a special type in a.c.
	HorizontalAlignment string `json:"horizontalAlignment,omitempty"`
	IsSubtle            *bool  `json:"isSubtle,omitempty"`
	MaxLines            int64  `json:"maxLines,omitempty"`
	Size                string `json:"size,omitempty"`
	Weight              string `json:"weight,omitempty"`
	Wrap                *bool  `json:"wrap,omitempty"`
	// inherited
	Fallback  []Node            `json:"fallback,omitempty"`
	Height    string            `json:"height,omitempty"`
	Separator *bool             `json:"separator,omitempty"`
	Spacing   string            `json:"spacing,omitempty"`
	ID        string            `json:"id,omitempty"`
	IsVisible *bool             `json:"isVisible,omitempty"`
	Requires  map[string]string `json:"requires,omitempty"`
}

func (n *TextBlock) prepare() error {
	n.Type = TextBlockType
	if n.Text == "" {
		return errors.New("TextBlock text is required")
	}
	return nil
}

// Image is Image element.
type Image struct {
	Type                string `json:"type"` // required, must be "Image"
	URL                 string `json:"url"`  // required
	AltText             string `json:"altText,omitempty"`
	BackgroundColor     string `json:"backgroundColor,omitempty"`
	Height              string `json:"height,omitempty"` // default "auto"
	HorizontalAlignment string `json:"horizontalAlignment,omitempty"`
	SelectAction        []Node `json:"selectAction,omitempty"`
	Size                string `json:"size,omitempty"`
	Style               string `json:"style,omitempty"` // "default" or "person"
	Width               string `json:"width,omitempty"`
	// inherited
	Fallback  []Node            `json:"fallback,omitempty"`
	Separator *bool             `json:"separator,omitempty"`
	Spacing   string            `json:"spacing,omitempty"`
	ID        string            `json:"id,omitempty"`
	IsVisible *bool             `json:"isVisible,omitempty"`
	Requires  map[string]string `json:"requires,omitempty"`
}

func (n *Image) prepare() error {
	n.Type = ImageType
	if n.URL == "" {
		return errors.New("Image url is required")
	}
	return nil
}

// Media displays a media player for audio or video content.
type Media struct {
	Type    string         `json:"type"`    // required, must be Media
	Sources []*MediaSource `json:"sources"` // required
	Poster  string         `json:"poster,omitempty"`
	AltText string         `json:"altText,omitempty"`
	// inherited
	Fallback  []Node            `json:"fallback,omitempty"`
	Height    string            `json:"height,omitempty"`
	Separator *bool             `json:"separator,omitempty"`
	Spacing   string            `json:"spacing,omitempty"`
	ID        string            `json:"id,omitempty"`
	IsVisible *bool             `json:"isVisible,omitempty"`
	Requires  map[string]string `json:"requires,omitempty"`
}

func (n *Media) prepare() error {
	n.Type = MediaType
	if len(n.Sources) < 1 {
		return errors.New("Media must have sources")
	}
	for _, s := range n.Sources {
		if err := s.prepare(); err != nil {
			return err
		}
	}
	return nil
}

// MediaSource defines a source for a Media element.
type MediaSource struct {
	MimeType string `json:"mimeType"` // required
	URL      string `json:"url"`      // required
}

func (m *MediaSource) prepare() error {
	if m.MimeType == "" {
		return errors.New("MediaSource must have mime type")
	}
	if m.URL == "" {
		return errors.New("MediaSource must have url")
	}
	return nil
}
