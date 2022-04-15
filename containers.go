package cards

import (
	"errors"
)

// ActionSet displays a set of actions.
type ActionSet struct {
	Type    string `json:"type"` // required
	Actions []Node `json:"actions,omitempty"`
	// inherited
	Fallback  []Node            `json:"fallback,omitempty"`
	Height    string            `json:"height,omitempty"`
	Separator *bool             `json:"separator,omitempty"`
	Spacing   string            `json:"spacing,omitempty"`
	ID        string            `json:"id,omitempty"`
	IsVisible *bool             `json:"isVisible,omitempty"`
	Requires  map[string]string `json:"requires,omitempty"`
}

func (n *ActionSet) prepare() error {
	n.Type = ActionSetType
	if len(n.Actions) < 1 {
		return errors.New("ActionSet must have elements")
	}
	for _, node := range n.Actions {
		if err := node.prepare(); err != nil {
			return err
		}
	}
	return nil
}

// Container groups items together.
type Container struct {
	Type            string           `json:"type"`  // required
	Items           []Node           `json:"items"` // required
	SelectAction    Node             `json:"selectAction,omitempty"`
	Style           string           `json:"style,omitempty"`
	Bleed           *bool            `json:"bleed,omitempty"`
	BackgroundImage *BackgroundImage `json:"backgroundImage,omitempty"`
	MinHeight       string           `json:"minHeight,omitempty"`
	// inherited
	Fallback  []Node            `json:"fallback,omitempty"`
	Height    string            `json:"height,omitempty"`
	Separator *bool             `json:"separator,omitempty"`
	Spacing   string            `json:"spacing,omitempty"`
	ID        string            `json:"id,omitempty"`
	IsVisible *bool             `json:"isVisible,omitempty"`
	Requires  map[string]string `json:"requires,omitempty"`
}

func (n *Container) prepare() error {
	n.Type = ContainerType
	if len(n.Items) < 1 {
		return errors.New("container must have elements")
	}
	for _, node := range n.Items {
		if err := node.prepare(); err != nil {
			return err
		}
	}
	return nil
}

// ColumnSet divides a region into Columns,
// allowing elements to sit side-by-side.
type ColumnSet struct {
	Type                string    `json:"type"` // required
	Columns             []*Column `json:"columns,omitempty"`
	SelectAction        Node      `json:"selectAction,omitempty"`
	Style               string    `json:"style,omitempty"`
	Bleed               *bool     `json:"bleed,omitempty"`
	MinHeight           string    `json:"minHeight,omitempty"`
	HorizontalAlignment string    `json:"horizontalAlignment,omitempty"`
	// inherited
	Fallback  []Node            `json:"fallback,omitempty"`
	Height    string            `json:"height,omitempty"`
	Spacing   string            `json:"spacing,omitempty"`
	ID        string            `json:"id,omitempty"`
	IsVisible *bool             `json:"isVisible,omitempty"`
	Requires  map[string]string `json:"requires,omitempty"`
}

func (n *ColumnSet) prepare() error {
	n.Type = ColumnSetType
	for _, c := range n.Columns {
		if err := c.prepare(); err != nil {
			return err
		}
	}
	return nil
}

// Column defines a container that is part of a ColumnSet.
type Column struct {
	Type                     string           `json:"type"` // required - it is not stated in a.c. docs but actually has to be "Column"
	Items                    []Node           `json:"items,omitempty"`
	BackgroundImage          *BackgroundImage `json:"backgroundImage,omitempty"`
	Bleed                    *bool            `json:"bleed,omitempty"`
	Fallback                 Node             `json:"fallback,omitempty"`
	MinHeight                string           `json:"minHeight,omitempty"`
	Separator                *bool            `json:"separator,omitempty"`
	Spacing                  string           `json:"spacing,omitempty"`
	SelectAction             Node             `json:"selectAction,omitempty"`
	Style                    string           `json:"style,omitempty"`
	VerticalContentAlignment string           `json:"verticalContentAlignment,omitempty"`
	Width                    string           `json:"width,omitempty"`
	// inherited
	ID        string            `json:"id,omitempty"`
	IsVisible *bool             `json:"isVisible,omitempty"`
	Requires  map[string]string `json:"requires,omitempty"`
}

func (c *Column) prepare() error {
	c.Type = ColumnType
	for _, node := range c.Items {
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

// FactSet element displays a series of facts (i.e. name/value pairs) in a tabular form.
type FactSet struct {
	Type  string  `json:"type"`  // required - must be "FactSet"
	Facts []*Fact `json:"facts"` // required
	// inherited
	Fallback  []Node            `json:"fallback,omitempty"`
	Height    string            `json:"height,omitempty"`
	Separator *bool             `json:"separator,omitempty"`
	Spacing   string            `json:"spacing,omitempty"`
	ID        string            `json:"id,omitempty"`
	IsVisible *bool             `json:"isVisible,omitempty"`
	Requires  map[string]string `json:"requires,omitempty"`
}

func (n *FactSet) prepare() error {
	n.Type = FactSetType
	if len(n.Facts) < 1 {
		return errors.New("FactSet must have facts")
	}
	for _, f := range n.Facts {
		if err := f.prepare(); err != nil {
			return err
		}
	}
	return nil
}

// Fact describes a Fact in a FactSet as a key/value pair.
type Fact struct {
	Title string `json:"title"` // required
	Value string `json:"value"` // required
}

func (f *Fact) prepare() error {
	if f.Title == "" {
		return errors.New("Fact must have title")
	}
	if f.Value == "" {
		return errors.New("Fact must have value")
	}
	return nil
}

// ImageSet displays a collection of Images similar to a gallery. Acceptable formats are PNG, JPEG, and GIF.
type ImageSet struct {
	Type      string   `json:"type"`   // required
	Images    []*Image `json:"images"` // required
	ImageSize string   `json:"imageSize,omitempty"`
	// inherited
	Fallback  []Node            `json:"fallback,omitempty"`
	Height    string            `json:"height,omitempty"`
	Separator *bool             `json:"separator,omitempty"`
	Spacing   string            `json:"spacing,omitempty"`
	ID        string            `json:"id,omitempty"`
	IsVisible *bool             `json:"isVisible,omitempty"`
	Requires  map[string]string `json:"requires,omitempty"`
}

func (n *ImageSet) prepare() error {
	n.Type = ImageSetType
	if len(n.Images) < 1 {
		return errors.New("ImageSet must have images")
	}
	for _, f := range n.Images {
		if err := f.prepare(); err != nil {
			return err
		}
	}
	return nil
}
