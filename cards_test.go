package cards

import (
	"io/ioutil"
	"testing"
)

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
									Wrap:   true,
								},
								TextBlock{
									Type:     TextBlockType,
									Spacing:  "none",
									Text:     "Created {{DATE(2017-02-14T06:08:39Z, SHORT)}}",
									IsSubtle: true,
									Wrap:     true,
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
					Wrap: true,
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
						IsMultiline: true,
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
	}).WithVersion(Version1)
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
