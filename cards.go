package cards

import (
	"encoding/json"
	"errors"
)

const (
	// DefaultSchema is card schema value by default
	DefaultSchema = "http://adaptivecards.io/schemas/adaptive-card.json"
	// Version1 is cards verstion 1.0
	Version1 = "1.0"
	// Version11 is cards verstion 1.1
	Version11 = "1.1"
	// Version12 is cards verstion 1.2
	Version12 = "1.2"
	// Version13 is cards verstion 1.3.
	// Warning maybe not supported by bot framework!
	Version13 = "1.3"

	// Types

	// AdaptiveCardType is general card type
	AdaptiveCardType = "AdaptiveCard"
	// TextBlockType is type for text block
	TextBlockType = "TextBlock"
	// ImageType is type for Image
	ImageType = "Image"
	// MediaType is type for Media
	MediaType = "Media"
	// RichTextBlockType is type for RichTextBlock
	RichTextBlockType = "RichTextBlock"
	// TextRunType is type for TextRun
	TextRunType = "TextRun"
	// ActionSetType is type for ActionSet
	ActionSetType = "ActionSet"
	// ContainerType is type for container
	ContainerType = "Container"
	// ColumnSetType is type for ColumnSet
	ColumnSetType = "ColumnSet"
	// ColumnType is type for Column
	ColumnType = "Column"
	// FactSetType is type for FactSet
	FactSetType = "FactSet"
	// ImageSetType is type for ImageSet
	ImageSetType = "ImageSet"
	// ActionShowCardType is type for Action.ShowCard
	ActionShowCardType = "Action.ShowCard"
	// ActionSubmitType is type for Action.Submit
	ActionSubmitType = "Action.Submit"
	// ActionOpenURLType is type for Action.OpenUrl
	ActionOpenURLType = "Action.OpenUrl"
	// ActionToggleVisibilityType is type for Action.ToggleVisibility
	ActionToggleVisibilityType = "Action.ToggleVisibility"
	// InputTextType is type for Input.Text
	InputTextType = "Input.Text"
	// InputNumberType is type for Input.Number
	InputNumberType = "Input.Number"
	// InputTimeType is type for Imput.Time
	InputTimeType = "Input.Time"
	// InputDateType is type for Input.Date
	InputDateType = "Input.Date"
	// InputChoiceSetType is type for Input.ChoiceSet
	InputChoiceSetType = "Input.ChoiceSet"
	// InputToggleType is type for Input.Toggle
	InputToggleType = "Input.Toggle"
)

// Node is card element
type Node interface {
	prepare() error
}

// Card is basic adaptive cards type.
type Card struct {
	Type                     string           `json:"type"`    // required
	Version                  string           `json:"version"` // required
	Schema                   string           `json:"$schema,omitempty"`
	Body                     []Node           `json:"body,omitempty"`
	Actions                  []Node           `json:"actions,omitempty"`
	SelectAction             Node             `json:"selectAction,omitempty"`
	FallbackText             string           `json:"fallbackText,omitempty"`
	BackgroundImage          *BackgroundImage `json:"backgroundImage,omitempty"`
	MinHeight                string           `json:"minHeight,omitempty"`
	Speak                    string           `json:"speak,omitempty"`
	Lang                     string           `json:"lang,omitempty"`
	VerticalContentAlignment string           `json:"verticalContentAlignment,omitempty"`
}

// New returns a card with provided body and default schema
func New(body []Node, actions []Node) *Card {
	return &Card{
		Type:    AdaptiveCardType,
		Version: Version13,
		Body:    body,
		Actions: actions,
	}
}

// WithVersion allows to set card version
func (c *Card) WithVersion(v string) *Card {
	c.Version = v
	return c
}

// WithSchema allows to set card schema
func (c *Card) WithSchema(s string) *Card {
	c.Schema = s
	return c
}

// WithBackgroundImage allows to set card image
func (c *Card) WithBackgroundImage(i BackgroundImage) *Card {
	c.BackgroundImage = &i
	return c
}

// WithMinHeight allows to set card min height
func (c *Card) WithMinHeight(h string) *Card {
	c.MinHeight = h
	return c
}

// Prepare validates card (required fields etc) and sets relevant types
func (c *Card) Prepare() error {
	c.Type = AdaptiveCardType
	if c.Version == "" {
		return errors.New("card version is required")
	}
	for _, node := range c.Body {
		if err := node.prepare(); err != nil {
			return err
		}
	}
	for _, node := range c.Actions {
		if err := node.prepare(); err != nil {
			return err
		}
	}
	if c.BackgroundImage != nil {
		if err := c.BackgroundImage.prepare(); err != nil {
			return err
		}
	}
	return nil
}

// Bytes returns adaptive card JSON as bytes
func (c *Card) Bytes() ([]byte, error) {
	if err := c.Prepare(); err != nil {
		return []byte{}, err
	}
	return json.Marshal(c)
}

// String returns adaptive card JSON as string
func (c *Card) String() (string, error) {
	cardJSON, err := c.Bytes()
	if err != nil {
		return "", err
	}
	return string(cardJSON), nil
}

// BytesIndent returns adaptive card JSON as bytes with indentation
func (c *Card) BytesIndent(prefix string, indent string) ([]byte, error) {
	if err := c.Prepare(); err != nil {
		return []byte{}, err
	}
	return json.MarshalIndent(c, prefix, indent)
}

// StringIndent returns adaptive card JSON as string with indentation
func (c *Card) StringIndent(prefix string, indent string) (string, error) {
	cardJSON, err := c.BytesIndent(prefix, indent)
	if err != nil {
		return "", err
	}
	return string(cardJSON), nil
}

// NestedCard is similar to adaptive card but doesn't require version and schema
type NestedCard struct {
	Type                     string           `json:"type"` // required
	Version                  string           `json:"version,omitempty"`
	Schema                   string           `json:"$schema,omitempty"`
	Body                     []Node           `json:"body,omitempty"`
	Actions                  []Node           `json:"actions,omitempty"`
	SelectAction             Node             `json:"selectAction,omitempty"`
	FallbackText             string           `json:"fallbackText,omitempty"`
	BackgroundImage          *BackgroundImage `json:"backgroundImage,omitempty"`
	MinHeight                string           `json:"minHeight,omitempty"`
	Speak                    string           `json:"speak,omitempty"`
	Lang                     string           `json:"lang,omitempty"`
	VerticalContentAlignment string           `json:"verticalContentAlignment,omitempty"`
}

func (n *NestedCard) prepare() error {
	n.Type = AdaptiveCardType
	for _, node := range n.Body {
		if err := node.prepare(); err != nil {
			return err
		}
	}
	for _, node := range n.Actions {
		if err := node.prepare(); err != nil {
			return err
		}
	}
	if n.BackgroundImage != nil {
		if err := n.BackgroundImage.prepare(); err != nil {
			return err
		}
	}
	return nil
}

// BackgroundImage specifies a background image. Acceptable formats are PNG, JPEG, and GIF.
type BackgroundImage struct {
	URL                 string `json:"url"` // required
	FillMode            string `json:"fillMode,omitempty"`
	HorizontalAlignment string `json:"horizontalAlignment,omitempty"`
	VerticalAlignment   string `json:"verticalAlignment,omitempty"`
}

func (b *BackgroundImage) prepare() error {
	if b.URL == "" {
		return errors.New("BackgroundImage must have url")
	}
	return nil
}
