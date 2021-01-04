package cards

import (
	"encoding/json"
	"errors"
	"fmt"
)

const (
	// DefaultSchema is card schema value by default
	DefaultSchema = "http://adaptivecards.io/schemas/adaptive-card.json"
	// Version1 is cards verstion 1.0
	Version1 = "1.0"
	// Version12 is cards verstion 1.2
	Version12 = "1.2"
	// Version13 is cards verstion 1.3
	Version13 = "1.3"

	// Types

	// AdaptiveCardType is general card type
	AdaptiveCardType = "AdaptiveCard"
	// ContainerType is type for container
	ContainerType = "Container"
	// TextBlockType is type for text block
	TextBlockType = "TextBlock"
	// InputTextType is type for Input.Text
	InputTextType = "Input.Text"
	// ColumnSetType is type for ColumnSet
	ColumnSetType = "ColumnSet"
	// ColumnType is type for Column
	ColumnType = "Column"
)

// Card is basic adaptive cards type.
type Card struct {
	Type    string `json:"type"`    // required
	Version string `json:"version"` // required
	Schema  string `json:"$schema"`
	Body    []Node `json:"body,omitempty"`
}

// New returns a card with provided body and default schema
func New(body []Node) *Card {
	return &Card{
		Schema:  DefaultSchema,
		Type:    AdaptiveCardType,
		Version: Version13,
		Body:    body,
	}
}

// WithVersion allows to set card version
func (c *Card) WithVersion(v string) *Card {
	c.Version = v
	return c
}

// Node is card element
type Node interface {
	validate() error
}

// Validate validates card (required fields etc)
func (c *Card) Validate() error {
	if c.Type != AdaptiveCardType {
		return fmt.Errorf("type must be %s", AdaptiveCardType)
	}
	if c.Version == "" {
		return errors.New("card version is required")
	}
	for _, node := range c.Body {
		if err := node.validate(); err != nil {
			return err
		}
	}
	return nil
}

// Bytes returns adaptive card JSON as bytes
func (c *Card) Bytes() ([]byte, error) {
	if err := c.Validate(); err != nil {
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
	if err := c.Validate(); err != nil {
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
