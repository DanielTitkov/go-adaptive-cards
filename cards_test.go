package cards

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestInputsCard(t *testing.T) {
	inputsCardJSON := mustReadFile("./test/inputs.json")

	var choices []InputChoice
	for i, color := range []string{"Red", "Green", "Blue"} {
		choices = append(choices, InputChoice{
			Title: color,
			Value: fmt.Sprint(i + 1),
		})
	}

	c := New([]Node{
		TextBlock{
			Type:                TextBlockType,
			Size:                "Medium",
			Weight:              "Bolder",
			Text:                "Input.Text elements",
			HorizontalAlignment: "Center",
			Wrap:                TruePtr(),
		},
		TextBlock{
			Type: TextBlockType,
			Text: "Name",
			Wrap: TruePtr(),
		},
		InputText{
			Type:  InputTextType,
			Style: "text",
			ID:    "SimpleVal",
		},
		TextBlock{
			Type: TextBlockType,
			Text: "Homepage",
			Wrap: TruePtr(),
		},
		InputText{
			Type:  InputTextType,
			Style: "url",
			ID:    "UrlVal",
		},
		TextBlock{
			Type: TextBlockType,
			Text: "Email",
			Wrap: TruePtr(),
		},
		InputText{
			Type:  InputTextType,
			Style: "email",
			ID:    "EmailVal",
		},
		TextBlock{
			Type: TextBlockType,
			Text: "Phone",
			Wrap: TruePtr(),
		},
		InputText{
			Type:  InputTextType,
			Style: "tel",
			ID:    "TelVal",
		},
		TextBlock{
			Type: TextBlockType,
			Text: "Comments",
			Wrap: TruePtr(),
		},
		InputText{
			Type:        InputTextType,
			Style:       "text",
			ID:          "MultiLineVal",
			IsMultiline: TruePtr(),
		},
		TextBlock{
			Type: TextBlockType,
			Text: "Quantity",
			Wrap: TruePtr(),
		},
		InputNumber{
			Type:  InputNumberType,
			ID:    "NumVal",
			Max:   5,
			Min:   -5,
			Value: 1,
		},
		TextBlock{
			Type: TextBlockType,
			Text: "Due Date",
			Wrap: TruePtr(),
		},
		InputDate{
			Type:  InputDateType,
			ID:    "DateVal",
			Value: "2017-09-20",
		},
		TextBlock{
			Type: TextBlockType,
			Text: "Start time",
			Wrap: TruePtr(),
		},
		InputTime{
			Type:  InputTimeType,
			ID:    "TimeVal",
			Value: "16:59",
		},
		TextBlock{
			Type:                TextBlockType,
			Text:                "Input ChoiceSet",
			Size:                "Medium",
			Weight:              "Bolder",
			HorizontalAlignment: "Center",
			Wrap:                TruePtr(),
		},
		TextBlock{
			Type: TextBlockType,
			Text: "What color do you want? (compact)",
			Wrap: TruePtr(),
		},
		InputChoiceSet{
			Type:    InputChoiceSetType,
			ID:      "CompactSelectVal",
			Value:   "1",
			Choices: choices,
		},
		TextBlock{
			Type: TextBlockType,
			Text: "What color do you want? (expanded)",
			Wrap: TruePtr(),
		},
		InputChoiceSet{
			Type:    InputChoiceSetType,
			ID:      "SingleSelectVal",
			Value:   "1",
			Style:   "expanded",
			Choices: choices,
		},
		TextBlock{
			Type: TextBlockType,
			Text: "What color do you want? (multiselect)",
			Wrap: TruePtr(),
		},
		InputChoiceSet{
			Type:          InputChoiceSetType,
			ID:            "MultiSelectVal",
			Value:         "1,3",
			IsMultiSelect: TruePtr(),
			Choices:       choices,
		},
		TextBlock{
			Type:                TextBlockType,
			Text:                "Input.Toggle",
			Size:                "Medium",
			Weight:              "Bolder",
			HorizontalAlignment: "Center",
			Wrap:                TruePtr(),
		},
		InputToggle{
			Type:  InputToggleType,
			Title: "I accept the terms and conditions (True/False)",
			ID:    "AcceptsTerms",
			Wrap:  FalsePtr(),
			Value: "false",
		},
		InputToggle{
			Type:     InputToggleType,
			Title:    "Red cars are better than other cars",
			ID:       "ColorPreference",
			Wrap:     FalsePtr(),
			Value:    "NotRedCars",
			ValueOff: "RedCars",
			ValueOn:  "NotRedCars",
		},
	}, []Node{
		ActionSubmit{
			Type:  ActionSubmitType,
			Title: "Submit",
			Data: map[string]interface{}{
				"id": "1234567890",
			},
		},
		ActionShowCard{
			Type:  ActionShowCardType,
			Title: "Show Card",
			Card: NestedCard{
				Type:   AdaptiveCardType,
				Schema: DefaultSchema,
				Body: []Node{
					TextBlock{
						Type: TextBlockType,
						Text: "Enter comment",
						Wrap: TruePtr(),
					},
					InputText{
						Type:  InputTextType,
						Style: "text",
						ID:    "CommentVal",
					},
				},
				Actions: []Node{
					ActionSubmit{
						Type:  ActionSubmitType,
						Title: "OK",
					},
				},
			},
		},
	}).WithVersion(Version1).WithSchema(DefaultSchema)
	got, err := c.StringIndent("", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if got != inputsCardJSON {
		t.Errorf("expected:\n%s\nbut got:\n%s", inputsCardJSON, got)
	}
}

func TestToggleCard(t *testing.T) {
	toggleCardJSON := mustReadFile("./test/toggle.json")
	c := New([]Node{
		TextBlock{
			Type: TextBlockType,
			Text: "Press the buttons to toggle the images!",
			Wrap: TruePtr(),
		},
		TextBlock{
			Type:      TextBlockType,
			Text:      "Here are some images:",
			IsVisible: FalsePtr(),
			ID:        "textToToggle",
		},
		ColumnSet{
			Type: ColumnSetType,
			Columns: []Column{
				{
					Type: ColumnType,
					Items: []Node{
						Image{
							Type:      ImageType,
							URL:       "https://picsum.photos/100/100?image=112",
							Style:     "person",
							IsVisible: FalsePtr(),
							ID:        "imageToToggle",
							AltText:   "sample image 1",
							Size:      "medium",
						},
					},
				},
			},
		},
	}, []Node{
		ActionToggleVisibility{
			Type:  ActionToggleVisibilityType,
			Title: "Toggle!",
			TargetElements: []TargetElement{
				{
					ElementID: "textToToggle",
				},
				{
					ElementID: "imageToToggle",
				},
			},
		},
		ActionToggleVisibility{
			Type:  ActionToggleVisibilityType,
			Title: "Show!",
			TargetElements: []TargetElement{
				{
					ElementID: "textToToggle",
					IsVisible: TruePtr(),
				},
				{
					ElementID: "imageToToggle",
					IsVisible: TruePtr(),
				},
			},
		},
		ActionToggleVisibility{
			Type:  ActionToggleVisibilityType,
			Title: "Hide!",
			TargetElements: []TargetElement{
				{
					ElementID: "textToToggle",
					IsVisible: FalsePtr(),
				},
				{
					ElementID: "imageToToggle",
					IsVisible: FalsePtr(),
				},
			},
		},
		ActionToggleVisibility{
			Type:  ActionToggleVisibilityType,
			Title: "Grain!",
			TargetElements: []TargetElement{
				{
					ElementID: "textToToggle",
					IsVisible: FalsePtr(),
				},
				{
					ElementID: "imageToToggle",
					IsVisible: TruePtr(),
				},
			},
		},
	}).WithVersion(Version12)
	got, err := c.StringIndent("", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if got != toggleCardJSON {
		t.Errorf("expected:\n%s\nbut got:\n%s", toggleCardJSON, got)
	}
}

func TestExampleCard(t *testing.T) {
	exampleCardJSON := mustReadFile("./test/example.json")
	c := New([]Node{
		Container{
			Type: ContainerType,
			Items: []Node{
				TextBlock{
					Type:   TextBlockType,
					Text:   "Publish Adaptive Card schema",
					Weight: "bolder",
					Size:   "medium",
				},
				ColumnSet{
					Type: ColumnSetType,
					Columns: []Column{
						{
							Type:  ColumnType,
							Width: "auto",
							Items: []Node{
								Image{
									Type:  ImageType,
									URL:   "https://pbs.twimg.com/profile_images/3647943215/d7f12830b3c17a5a9e4afcc370e3a37e_400x400.jpeg",
									Size:  "small",
									Style: "person",
								},
							},
						},
						{
							Type:  ColumnType,
							Width: "stretch",
							Items: []Node{
								TextBlock{
									Type:   TextBlockType,
									Text:   "Matt Hidinger",
									Weight: "bolder",
									Wrap:   TruePtr(),
								},
								TextBlock{
									Type:     TextBlockType,
									Spacing:  "none",
									Text:     "Created {{DATE(2017-02-14T06:08:39Z, SHORT)}}",
									IsSubtle: TruePtr(),
									Wrap:     TruePtr(),
								},
							},
						},
					},
				},
			},
		},
		Container{
			Type: ContainerType,
			Items: []Node{
				TextBlock{
					Type: TextBlockType,
					Text: "Now that we have defined the main rules...",
					Wrap: TruePtr(),
				},
				FactSet{
					Type: FactSetType,
					Facts: []Fact{
						{
							Title: "Board:",
							Value: "Adaptive Card",
						},
						{
							Title: "List:",
							Value: "Backlog",
						},
						{
							Title: "Assigned to:",
							Value: "Matt Hidinger",
						},
						{
							Title: "Due date:",
							Value: "Not set",
						},
					},
				},
			},
		},
	}, []Node{
		ActionShowCard{
			Type:  ActionShowCardType,
			Title: "Comment",
			Card: NestedCard{
				Type: AdaptiveCardType,
				Body: []Node{
					InputText{
						Type:        InputTextType,
						ID:          "comment",
						IsMultiline: TruePtr(),
						Placeholder: "Enter your comment",
					},
				},
				Actions: []Node{
					ActionSubmit{
						Type:  ActionSubmitType,
						Title: "OK",
					},
				},
			},
		},
		ActionOpenURL{
			Type:  ActionOpenURLType,
			Title: "View",
			URL:   "https://adaptivecards.io",
		},
	}).WithVersion(Version1).WithSchema(DefaultSchema)
	got, err := c.StringIndent("", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if got != exampleCardJSON {
		t.Errorf("expected:\n%s\nbut got:\n%s", exampleCardJSON, got)
	}
}

func TestInvalidCard(t *testing.T) {
	c := Card{Body: []Node{}}
	err := c.Validate()
	if err == nil {
		t.Error("expected to have an error, got nil")
	}
}

func mustReadFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}
