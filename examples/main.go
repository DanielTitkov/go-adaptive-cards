package main

import (
	"fmt"
	"log"

	cards "github.com/DanielTitkov/go-adaptive-cards"
)

func main() {
	log.Println("Go Adaptive Cards example")

	c := cards.New([]cards.Node{
		cards.Container{
			Type: cards.ContainerType,
			Items: []cards.Node{
				cards.TextBlock{
					Type:     cards.TextBlockType,
					Text:     "foo",
					IsSubtle: true,
				},
			},
		},
	})

	s, err := c.StringIndent("", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
	log.Println("Done")
}
