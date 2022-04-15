package cards

import (
	"fmt"
	"io/ioutil"
	"testing"
)

func TestInputsCard(t *testing.T) {
	inputsCardJSON := mustReadFile("./test/inputs.json")

	var choices []*InputChoice
	for i, color := range []string{"Red", "Green", "Blue"} {
		choices = append(choices, &InputChoice{
			Title: color,
			Value: fmt.Sprint(i + 1),
		})
	}

	c := New([]Node{
		&TextBlock{
			Size:                "Medium",
			Weight:              "Bolder",
			Text:                "Input.Text elements",
			HorizontalAlignment: "Center",
			Wrap:                TruePtr(),
		},
		&TextBlock{
			Text: "Name",
			Wrap: TruePtr(),
		},
		&InputText{
			Style: "text",
			ID:    "SimpleVal",
		},
		&TextBlock{
			Text: "Homepage",
			Wrap: TruePtr(),
		},
		&InputText{
			Style: "url",
			ID:    "UrlVal",
		},
		&TextBlock{
			Text: "Email",
			Wrap: TruePtr(),
		},
		&InputText{
			Style: "email",
			ID:    "EmailVal",
		},
		&TextBlock{
			Text: "Phone",
			Wrap: TruePtr(),
		},
		&InputText{
			Style: "tel",
			ID:    "TelVal",
		},
		&TextBlock{
			Text: "Comments",
			Wrap: TruePtr(),
		},
		&InputText{
			Style:       "text",
			ID:          "MultiLineVal",
			IsMultiline: TruePtr(),
		},
		&TextBlock{
			Text: "Quantity",
			Wrap: TruePtr(),
		},
		&InputNumber{
			ID:    "NumVal",
			Max:   5,
			Min:   -5,
			Value: 1,
		},
		&TextBlock{
			Text: "Due Date",
			Wrap: TruePtr(),
		},
		&InputDate{
			ID:    "DateVal",
			Value: "2017-09-20",
		},
		&TextBlock{
			Text: "Start time",
			Wrap: TruePtr(),
		},
		&InputTime{
			ID:    "TimeVal",
			Value: "16:59",
		},
		&TextBlock{
			Text:                "Input ChoiceSet",
			Size:                "Medium",
			Weight:              "Bolder",
			HorizontalAlignment: "Center",
			Wrap:                TruePtr(),
		},
		&TextBlock{
			Text: "What color do you want? (compact)",
			Wrap: TruePtr(),
		},
		&InputChoiceSet{
			ID:      "CompactSelectVal",
			Value:   "1",
			Choices: choices,
		},
		&TextBlock{
			Text: "What color do you want? (expanded)",
			Wrap: TruePtr(),
		},
		&InputChoiceSet{
			ID:      "SingleSelectVal",
			Value:   "1",
			Style:   "expanded",
			Choices: choices,
		},
		&TextBlock{
			Text: "What color do you want? (multiselect)",
			Wrap: TruePtr(),
		},
		&InputChoiceSet{
			ID:            "MultiSelectVal",
			Value:         "1,3",
			IsMultiSelect: TruePtr(),
			Choices:       choices,
		},
		&TextBlock{
			Text:                "Input.Toggle",
			Size:                "Medium",
			Weight:              "Bolder",
			HorizontalAlignment: "Center",
			Wrap:                TruePtr(),
		},
		&InputToggle{
			Title: "I accept the terms and conditions (True/False)",
			ID:    "AcceptsTerms",
			Wrap:  FalsePtr(),
			Value: "false",
		},
		&InputToggle{
			Title:    "Red cars are better than other cars",
			ID:       "ColorPreference",
			Wrap:     FalsePtr(),
			Value:    "NotRedCars",
			ValueOff: "RedCars",
			ValueOn:  "NotRedCars",
		},
	}, []Node{
		&ActionSubmit{
			Title: "Submit",
			Data: map[string]interface{}{
				"id": "1234567890",
			},
		},
		&ActionShowCard{
			Title: "Show Card",
			Card: NestedCard{
				Schema: DefaultSchema,
				Body: []Node{
					&TextBlock{
						Text: "Enter comment",
						Wrap: TruePtr(),
					},
					&InputText{
						Style: "text",
						ID:    "CommentVal",
					},
				},
				Actions: []Node{
					&ActionSubmit{
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

func TestBackGroundImage(t *testing.T) {
	backgroundCardJSON := mustReadFile("./test/background.json")
	c := New([]Node{
		&TextBlock{
			Text:   "Here is something about a cat",
			Weight: "Bolder",
			Wrap:   TruePtr(),
		},
		&TextBlock{
			Text: "Cat is good, cat is better, cat is really really greate",
			Wrap: TruePtr(),
		},
	}, []Node{}).
		WithVersion(Version12).
		WithBackgroundImage(
			BackgroundImage{
				URL:      "https://adaptivecards.io/content/cats/1.png",
				FillMode: "cover",
			},
		).
		WithMinHeight("500px")
	got, err := c.StringIndent("", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if got != backgroundCardJSON {
		t.Errorf("expected:\n%s\nbut got:\n%s", backgroundCardJSON, got)
	}
}

func TestToggleCard(t *testing.T) {
	toggleCardJSON := mustReadFile("./test/toggle.json")
	c := New([]Node{
		&TextBlock{
			Text: "Press the buttons to toggle the images!",
			Wrap: TruePtr(),
		},
		&TextBlock{
			Text:      "Here are some images:",
			IsVisible: FalsePtr(),
			ID:        "textToToggle",
		},
		&ColumnSet{
			Columns: []*Column{
				{
					Items: []Node{
						&Image{
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
		&ActionToggleVisibility{
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
		&ActionToggleVisibility{
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
		&ActionToggleVisibility{
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
		&ActionToggleVisibility{
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
		&Container{
			Items: []Node{
				&TextBlock{
					Text:   "Publish Adaptive Card schema",
					Weight: "bolder",
					Size:   "medium",
				},
				&ColumnSet{
					Columns: []*Column{
						{
							Width: "auto",
							Items: []Node{
								&Image{
									URL:   "https://pbs.twimg.com/profile_images/3647943215/d7f12830b3c17a5a9e4afcc370e3a37e_400x400.jpeg",
									Size:  "small",
									Style: "person",
								},
							},
						},
						{
							Width: "stretch",
							Items: []Node{
								&TextBlock{
									Text:   "Matt Hidinger",
									Weight: "bolder",
									Wrap:   TruePtr(),
								},
								&TextBlock{
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
		&Container{
			Items: []Node{
				&TextBlock{
					Text: "Now that we have defined the main rules...",
					Wrap: TruePtr(),
				},
				&FactSet{
					Facts: []*Fact{
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
		&ActionShowCard{
			Title: "Comment",
			Card: NestedCard{
				Body: []Node{
					&InputText{
						ID:          "comment",
						IsMultiline: TruePtr(),
						Placeholder: "Enter your comment",
					},
				},
				Actions: []Node{
					&ActionSubmit{
						Title: "OK",
					},
				},
			},
		},
		&ActionOpenURL{
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
	err := c.Prepare()
	if err == nil {
		t.Error("expected to have an error, got nil")
	}
}

func TestMedia(t *testing.T) {
	mediaCardJSON := mustReadFile("./test/media.json")
	c := New([]Node{
		&TextBlock{
			Text: "Media supports **audio** and **video** content!",
			Wrap: TruePtr(),
		},
		&TextBlock{
			Text:                "Video",
			HorizontalAlignment: "center",
			Spacing:             "medium",
			Separator:           TruePtr(),
			Size:                "large",
		},
		&Media{
			Poster:  "https://adaptivecards.io/content/poster-video.png",
			AltText: "Adaptive Cards overview video",
			Sources: []*MediaSource{
				{
					MimeType: "video/mp4",
					URL:      "https://adaptivecardsblob.blob.core.windows.net/assets/AdaptiveCardsOverviewVideo.mp4",
				},
			},
		},
		&TextBlock{
			Text:                "Audio",
			HorizontalAlignment: "center",
			Separator:           TruePtr(),
			Size:                "large",
		},
		&Media{
			Poster:  "https://adaptivecards.io/content/poster-audio.jpg",
			AltText: "Adaptive Cards overview audio",
			Sources: []*MediaSource{
				{
					MimeType: "audio/mpeg",
					URL:      "https://adaptivecardsblob.blob.core.windows.net/assets/AdaptiveCardsOverviewVideo.mp3",
				},
			},
		},
	}, []Node{}).WithVersion(Version11).WithSchema(DefaultSchema)
	got, err := c.StringIndent("", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if got != mediaCardJSON {
		t.Errorf("expected:\n%s\nbut got:\n%s", mediaCardJSON, got)
	}
}

func TestRichTextBlock(t *testing.T) {
	richCardJSON := mustReadFile("./test/rich.json")
	c := New([]Node{
		&RichTextBlock{
			Inlines: []*TextRun{
				{
					Text:  "We support colors,",
					Color: "good",
				},
				{
					Text:     " both regular and subtle. ",
					IsSubtle: TruePtr(),
				},
				{
					Text: "Text ",
					Size: "small",
				},
				{
					Text: "sizes! ",
					Size: "extraLarge",
				},
				{
					Text:   "Light weight text. ",
					Weight: "lighter",
				},
				{
					Text:      "Highlights. ",
					Highlight: TruePtr(),
				},
				{
					Text:   "Italics. ",
					Italic: TruePtr(),
				},
				{
					Text:          "Strikethrough. ",
					Strikethrough: TruePtr(),
				},
				{
					Text:     "Monospace too!",
					FontType: "monospace",
				},
			},
		},
		&RichTextBlock{
			Inlines: []*TextRun{
				{
					Text: "Date-Time parsing: {{DATE(2017-02-14T06:08:39Z,LONG)}} {{TIME(2017-02-14T06:08:39Z)}}",
				},
			},
		},
		&RichTextBlock{
			HorizontalAlignment: "center",
			Inlines: []*TextRun{
				{
					Text: "Rich text blocks also support center alignment.",
				},
			},
		},
		&RichTextBlock{
			HorizontalAlignment: "right",
			Inlines: []*TextRun{
				{
					Text: "Rich text blocks also support right alignment.",
				},
			},
		},
	}, []Node{}).WithVersion(Version12).WithSchema(DefaultSchema)
	got, err := c.StringIndent("", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if got != richCardJSON {
		t.Errorf("expected:\n%s\nbut got:\n%s", richCardJSON, got)
	}
}

func TestImageSet(t *testing.T) {
	imagesCardJSON := mustReadFile("./test/images.json")
	c := New([]Node{
		&ImageSet{
			ImageSize: "small",
			Images: []*Image{
				{
					URL: "https://adaptivecards.io/content/cats/1.png",
				},
				{
					URL: "https://adaptivecards.io/content/cats/2.png",
				},
				{
					URL: "https://adaptivecards.io/content/cats/3.png",
				},
			},
		},
	}, []Node{}).WithVersion(Version1).WithSchema(DefaultSchema)
	got, err := c.StringIndent("", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if got != imagesCardJSON {
		t.Errorf("expected:\n%s\nbut got:\n%s", imagesCardJSON, got)
	}
}

func TestActionSet(t *testing.T) {
	actionSetJSON := mustReadFile("./test/actionSet.json")
	c := New([]Node{
		&TextBlock{
			Text: "Cards can have action sets in the middle of their body.",
			Wrap: TruePtr(),
		},
		&ActionSet{
			Type: ActionSetType,
			Actions: []Node{
				&ActionShowCard{
					Type: ActionShowCardType,
					Card: NestedCard{
						Type: AdaptiveCardType,
						Body: []Node{
							&TextBlock{
								Type: TextBlockType,
								Text: "This is a show card",
							},
						},
					},
					Title: "ShowCard",
				},
				&ActionOpenURL{
					Type:  ActionOpenURLType,
					URL:   "https://adaptivecards.io",
					Title: "OpenUrl",
				},
			},
		},
	}, []Node{
		&ActionShowCard{
			Type: ActionShowCardType,
			Card: NestedCard{
				Type: AdaptiveCardType,
				Body: []Node{
					&TextBlock{
						Type: TextBlockType,
						Text: "This is a show card",
					},
				},
			},
			Title: "ShowCard",
		},
		&ActionOpenURL{
			Type:  ActionOpenURLType,
			URL:   "https://adaptivecards.io",
			Title: "OpenUrl",
		},
	}).WithVersion(Version1).WithSchema(DefaultSchema)
	got, err := c.StringIndent("", "  ")
	if err != nil {
		t.Fatal(err)
	}
	if got != actionSetJSON {
		t.Errorf("expected:\n%s\nbut got:\n%s", actionSetJSON, got)
	}
}

func mustReadFile(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}
