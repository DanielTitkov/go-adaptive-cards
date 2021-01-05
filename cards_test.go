package cards

import (
	"io/ioutil"
	"testing"
)

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
