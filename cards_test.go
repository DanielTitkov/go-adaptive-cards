package cards

import "testing"

var exampleCardJSON string = `{
	"$schema": "http://adaptivecards.io/schemas/adaptive-card.json",
	"type": "AdaptiveCard",
	"version": "1.0",
	"body": [
	  {
		"type": "Container",
		"items": [
		  {
			"type": "TextBlock",
			"text": "Publish Adaptive Card schema",
			"weight": "bolder",
			"size": "medium"
		  },
		  {
			"type": "ColumnSet",
			"columns": [
			  {
				"type": "Column",
				"width": "auto",
				"items": [
				  {
					"type": "Image",
					"url": "https://pbs.twimg.com/profile_images/3647943215/d7f12830b3c17a5a9e4afcc370e3a37e_400x400.jpeg",
					"size": "small",
					"style": "person"
				  }
				]
			  },
			  {
				"type": "Column",
				"width": "stretch",
				"items": [
				  {
					"type": "TextBlock",
					"text": "Matt Hidinger",
					"weight": "bolder",
					"wrap": true
				  },
				  {
					"type": "TextBlock",
					"spacing": "none",
					"text": "Created {{DATE(2017-02-14T06:08:39Z, SHORT)}}",
					"isSubtle": true,
					"wrap": true
				  }
				]
			  }
			]
		  }
		]
	  },
	  {
		"type": "Container",
		"items": [
		  {
			"type": "TextBlock",
			"text": "Now that we have defined the main rules...",
			"wrap": true
		  },
		  {
			"type": "FactSet",
			"facts": [
			  {
				"title": "Board:",
				"value": "Adaptive Card"
			  },
			  {
				"title": "List:",
				"value": "Backlog"
			  },
			  {
				"title": "Assigned to:",
				"value": "Matt Hidinger"
			  },
			  {
				"title": "Due date:",
				"value": "Not set"
			  }
			]
		  }
		]
	  }
	],
	"actions": [
	  {
		"type": "Action.ShowCard",
		"title": "Comment",
		"card": {
		  "type": "AdaptiveCard",
		  "body": [
			{
			  "type": "Input.Text",
			  "id": "comment",
			  "isMultiline": true,
			  "placeholder": "Enter your comment"
			}
		  ],
		  "actions": [
			{
			  "type": "Action.Submit",
			  "title": "OK"
			}
		  ]
		}
	  },
	  {
		"type": "Action.OpenUrl",
		"title": "View",
		"url": "https://adaptivecards.io"
	  }
	]
  }`

func TestExampleCard(t *testing.T) {
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
			},
		},
	}).WithVersion(Version1)
	got, err := c.StringIndent("", "  ")
	if err != nil {
		t.Error(err)
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
