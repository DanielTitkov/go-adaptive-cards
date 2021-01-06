package cards

import (
	"errors"
)

// ActionShowCard defines an AdaptiveCard which is shown to the user when the button or link is clicked.
type ActionShowCard struct {
	Type string     `json:"type"` // required
	Card NestedCard `json:"card,omitempty"`
	// inherited
	Title    string            `json:"title,omitempty"`
	IconURL  string            `json:"iconUrl,omitempty"`
	Style    string            `json:"style,omitempty"`
	Fallback []Node            `json:"fallback,omitempty"`
	Requires map[string]string `json:"requires,omitempty"`
}

func (n *ActionShowCard) prepare() error {
	n.Type = ActionShowCardType
	if err := n.Card.prepare(); err != nil {
		return err
	}
	return nil
}

// ActionSubmit gathers input fields, merges with optional data field,
// and sends an event to the client.
// It is up to the client to determine how this data is processed.
// See https://docs.microsoft.com/en-us/adaptive-cards/authoring-cards/input-validation for more details.
type ActionSubmit struct {
	Type             string                 `json:"type"` // required
	Data             map[string]interface{} `json:"data,omitempty"`
	AssociatedInputs string                 `json:"associatedInputs,omitempty"`
	// inherited
	Title    string            `json:"title,omitempty"`
	IconURL  string            `json:"iconUrl,omitempty"`
	Style    string            `json:"style,omitempty"`
	Fallback []Node            `json:"fallback,omitempty"`
	Requires map[string]string `json:"requires,omitempty"`
}

func (n *ActionSubmit) prepare() error {
	n.Type = ActionSubmitType
	return nil
}

// ActionOpenURL when invoked, show the given url
// either by launching it in an external web browser or showing within an embedded web browser.
type ActionOpenURL struct {
	Type string `json:"type"` // required
	URL  string `json:"url"`  // required
	// inherited
	Title    string            `json:"title,omitempty"`
	IconURL  string            `json:"iconUrl,omitempty"`
	Style    string            `json:"style,omitempty"`
	Fallback []Node            `json:"fallback,omitempty"`
	Requires map[string]string `json:"requires,omitempty"`
}

func (n *ActionOpenURL) prepare() error {
	n.Type = ActionOpenURLType
	return nil
}

// ActionToggleVisibility toggles the visibility of associated card elements.
type ActionToggleVisibility struct {
	Type           string          `json:"type"` // required
	TargetElements []TargetElement `json:"targetElements,omitempty"`
	// inherited
	Title    string            `json:"title,omitempty"`
	IconURL  string            `json:"iconUrl,omitempty"`
	Style    string            `json:"style,omitempty"`
	Fallback []Node            `json:"fallback,omitempty"`
	Requires map[string]string `json:"requires,omitempty"`
}

func (n *ActionToggleVisibility) prepare() error {
	n.Type = ActionToggleVisibilityType
	for _, e := range n.TargetElements {
		if err := e.prepare(); err != nil {
			return err
		}
	}
	return nil
}

// TargetElement represents an entry for Action.ToggleVisibility's targetElements property
type TargetElement struct {
	ElementID string `json:"elementId"` // required
	IsVisible *bool  `json:"isVisible,omitempty"`
}

func (t *TargetElement) prepare() error {
	if t.ElementID == "" {
		return errors.New("TargetElement element id is required")
	}
	return nil
}

// ActionInheritedFields holds common inherited fields for actions.
// Not really used for the brewety of API.
type ActionInheritedFields struct {
	Title    string            `json:"title,omitempty"`
	IconURL  string            `json:"iconUrl,omitempty"`
	Style    string            `json:"style,omitempty"`
	Fallback []Node            `json:"fallback,omitempty"`
	Requires map[string]string `json:"requires,omitempty"`
}

func (n ActionInheritedFields) validateInherited() error {
	return nil
}
