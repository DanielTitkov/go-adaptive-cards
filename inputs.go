package cards

import (
	"errors"
)

// InputText is Input.Text type
type InputText struct {
	Type         string `json:"type"` // required
	ID           string `json:"id"`   // required
	IsMultiline  *bool  `json:"isMultiline,omitempty"`
	MaxLength    int64  `json:"maxLength,omitempty"`
	Placeholder  string `json:"placeholder,omitempty"`
	Regex        string `json:"regex,omitempty"`
	Style        string `json:"style,omitempty"`
	InlineAction Node   `json:"inlineAction,omitempty"` // FIXME
	Value        string `json:"value,omitempty"`
	// inherited
	ErrorMessage string            `json:"errorMessage,omitempty"`
	IsRequired   *bool             `json:"isRequired,omitempty"`
	Label        string            `json:"label,omitempty"`
	Height       string            `json:"height,omitempty"`
	Separator    *bool             `json:"separator,omitempty"`
	Spacing      string            `json:"spacing,omitempty"`
	IsVisible    *bool             `json:"isVisible,omitempty"`
	Requires     map[string]string `json:"requires,omitempty"`
}

func (n *InputText) prepare() error {
	n.Type = InputTextType
	if n.ID == "" {
		return errors.New("InputText id is required")
	}
	return nil
}

// InputNumber allows a user to enter a number.
type InputNumber struct {
	Type        string  `json:"type"` // required
	ID          string  `json:"id"`   // required
	Max         float64 `json:"max,omitempty"`
	Min         float64 `json:"min,omitempty"`
	Placeholder string  `json:"placeholder,omitempty"`
	Value       float64 `json:"value,omitempty"`
	// inherited
	ErrorMessage string            `json:"errorMessage,omitempty"`
	IsRequired   *bool             `json:"isRequired,omitempty"`
	Label        string            `json:"label,omitempty"`
	Fallback     []Node            `json:"fallback,omitempty"`
	Height       string            `json:"height,omitempty"`
	Separator    *bool             `json:"separator,omitempty"`
	Spacing      string            `json:"spacing,omitempty"`
	IsVisible    *bool             `json:"isVisible,omitempty"`
	Requires     map[string]string `json:"requires,omitempty"`
}

func (n *InputNumber) prepare() error {
	n.Type = InputNumberType
	if n.ID == "" {
		return errors.New("InputNumber id is required")
	}
	return nil
}

// InputTime lets a user select a time.
type InputTime struct {
	Type        string `json:"type"` // required
	ID          string `json:"id"`   // required
	Max         string `json:"max,omitempty"`
	Min         string `json:"min,omitempty"`
	Placeholder string `json:"placeholder,omitempty"`
	Value       string `json:"value,omitempty"`
	// inherited
	ErrorMessage string            `json:"errorMessage,omitempty"`
	IsRequired   *bool             `json:"isRequired,omitempty"`
	Label        string            `json:"label,omitempty"`
	Fallback     []Node            `json:"fallback,omitempty"`
	Height       string            `json:"height,omitempty"`
	Separator    *bool             `json:"separator,omitempty"`
	Spacing      string            `json:"spacing,omitempty"`
	IsVisible    *bool             `json:"isVisible,omitempty"`
	Requires     map[string]string `json:"requires,omitempty"`
}

func (n *InputTime) prepare() error {
	n.Type = InputTimeType
	if n.ID == "" {
		return errors.New("InputTime id is required")
	}
	return nil
}

// InputDate lets a user choose a date.
type InputDate struct {
	Type        string `json:"type"` // required
	ID          string `json:"id"`   // required
	Max         string `json:"max,omitempty"`
	Min         string `json:"min,omitempty"`
	Placeholder string `json:"placeholder,omitempty"`
	Value       string `json:"value,omitempty"`
	// inherited
	ErrorMessage string            `json:"errorMessage,omitempty"`
	IsRequired   *bool             `json:"isRequired,omitempty"`
	Label        string            `json:"label,omitempty"`
	Fallback     []Node            `json:"fallback,omitempty"`
	Height       string            `json:"height,omitempty"`
	Separator    *bool             `json:"separator,omitempty"`
	Spacing      string            `json:"spacing,omitempty"`
	IsVisible    *bool             `json:"isVisible,omitempty"`
	Requires     map[string]string `json:"requires,omitempty"`
}

func (n *InputDate) prepare() error {
	n.Type = InputDateType
	if n.ID == "" {
		return errors.New("InputDate id is required")
	}
	return nil
}

// InputChoiceSet allows a user to input a Choice.
type InputChoiceSet struct {
	Type          string         `json:"type"`    // required
	Choices       []*InputChoice `json:"choices"` // required
	ID            string         `json:"id"`      // required
	IsMultiSelect *bool          `json:"isMultiSelect,omitempty"`
	Style         string         `json:"style,omitempty"`
	Placeholder   string         `json:"placeholder,omitempty"`
	Value         string         `json:"value,omitempty"`
	Wrap          *bool          `json:"wrap,omitempty"`
	// inherited
	ErrorMessage string            `json:"errorMessage,omitempty"`
	IsRequired   *bool             `json:"isRequired,omitempty"`
	Label        string            `json:"label,omitempty"`
	Fallback     []Node            `json:"fallback,omitempty"`
	Height       string            `json:"height,omitempty"`
	Separator    *bool             `json:"separator,omitempty"`
	Spacing      string            `json:"spacing,omitempty"`
	IsVisible    *bool             `json:"isVisible,omitempty"`
	Requires     map[string]string `json:"requires,omitempty"`
}

func (n *InputChoiceSet) prepare() error {
	n.Type = InputChoiceSetType
	if n.ID == "" {
		return errors.New("InputChoiceSet id is required")
	}
	if len(n.Choices) < 1 {
		return errors.New("InputChoiceSet must have choices")
	}
	for _, c := range n.Choices {
		if err := c.prepare(); err != nil {
			return err
		}
	}
	return nil
}

// InputChoice describes a choice for use in a ChoiceSet.
// Warning: do not use comma is the value.
type InputChoice struct {
	Title string `json:"title"` // required
	Value string `json:"value"` // required
}

func (c *InputChoice) prepare() error {
	if c.Title == "" {
		return errors.New("Choice must have title")
	}
	if c.Value == "" {
		return errors.New("Choice must have value")
	}
	// TODO validate no comma in the value
	return nil
}

// InputToggle lets a user choose between two options.
type InputToggle struct {
	Type     string `json:"type"`  // required
	Title    string `json:"title"` // required
	ID       string `json:"id"`    // required
	Value    string `json:"value,omitempty"`
	ValueOff string `json:"valueOff,omitempty"`
	ValueOn  string `json:"valueOn,omitempty"`
	Wrap     *bool  `json:"wrap,omitempty"`
	// inherited
	ErrorMessage string            `json:"errorMessage,omitempty"`
	IsRequired   *bool             `json:"isRequired,omitempty"`
	Label        string            `json:"label,omitempty"`
	Fallback     []Node            `json:"fallback,omitempty"`
	Height       string            `json:"height,omitempty"`
	Separator    *bool             `json:"separator,omitempty"`
	Spacing      string            `json:"spacing,omitempty"`
	IsVisible    *bool             `json:"isVisible,omitempty"`
	Requires     map[string]string `json:"requires,omitempty"`
}

func (n *InputToggle) prepare() error {
	n.Type = InputToggleType
	if n.ID == "" {
		return errors.New("InputToggle id is required")
	}
	if n.Title == "" {
		return errors.New("InputToggle must have title")
	}
	return nil
}
