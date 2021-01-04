package cards

import (
	"errors"
	"fmt"
)

// Container groups items together.
type Container struct {
	Type  string `json:"type"`  // required
	Items []Node `json:"items"` // required
}

func (n Container) validate() error {
	if n.Type != ContainerType {
		return fmt.Errorf("type must be %s", ContainerType)
	}
	if len(n.Items) < 1 {
		return errors.New("container must have elements")
	}
	for _, node := range n.Items {
		if err := node.validate(); err != nil {
			return err
		}
	}
	return nil
}

// ColumnSet divides a region into Columns,
// allowing elements to sit side-by-side.
type ColumnSet struct {
	Type                string   `json:"type"`              // required
	Columns             []Column `json:"columns,omitempty"` // TODO: maybe make it a Node
	SelectAction        []Node   `json:"selectAction,omitempty"`
	Style               []Node   `json:"style,omitempty"` // FIXME
	Bleed               bool     `json:"bleed,omitempty"`
	MinHeight           string   `json:"minHeight,omitempty"`
	HorizontalAlignment string   `json:"horizontalAlignment,omitempty"`
	// inherited
	Fallback  []Node            `json:"fallback,omitempty"`
	Height    string            `json:"height,omitempty"`
	Spacing   string            `json:"spacing,omitempty"`
	ID        string            `json:"id,omitempty"`
	IsVisible bool              `json:"isVisible,omitempty"`
	Requires  map[string]string `json:"requires,omitempty"`
}

func (n ColumnSet) validate() error {
	if n.Type != ColumnSetType {
		return fmt.Errorf("ColumnSet type must be %s", ColumnSetType)
	}
	for _, c := range n.Columns {
		if err := c.validate(); err != nil {
			return err
		}
	}
	return nil
}

// Column defines a container that is part of a ColumnSet.
type Column struct {
	Type                     string `json:"type"` // required - it is not stated in a.c. docs but actually has to be "Column"
	Items                    []Node `json:"columns,omitempty"`
	BackgroundImage          string `json:"backgroundImage,omitempty"`
	Bleed                    bool   `json:"bleed,omitempty"`
	Fallback                 Node   `json:"fallback,omitempty"`
	MinHeight                string `json:"minHeight,omitempty"`
	Separator                bool   `json:"separator,omitempty"`
	Spacing                  string `json:"spacing,omitempty"`
	SelectAction             []Node `json:"selectAction,omitempty"`
	Style                    []Node `json:"style,omitempty"` // FIXME
	VerticalContentAlignment string `json:"verticalContentAlignment,omitempty"`
	Width                    string `json:"width,omitempty"`
	// inherited
	ID        string            `json:"id,omitempty"`
	IsVisible bool              `json:"isVisible,omitempty"`
	Requires  map[string]string `json:"requires,omitempty"`
}

func (c Column) validate() error {
	if c.Type != ColumnType {
		return fmt.Errorf("Column type must be %s", ColumnType)
	}
	for _, node := range c.Items {
		if err := node.validate(); err != nil {
			return err
		}
	}
	return nil
}
