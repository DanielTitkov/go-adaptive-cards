package main

import (
	"fmt"
	"log"

	cards "github.com/DanielTitkov/go-adaptive-cards"
)

func main() {
	log.Println("Go Adaptive Cards example")
	log.Println("Serializing card...")

	c := cards.New([]cards.Node{
		&cards.Container{
			Items: []cards.Node{
				&cards.TextBlock{
					Text:     "foo",
					IsSubtle: cards.FalsePtr(),
				},
			},
		},
	}, []cards.Node{}).
		WithSchema(cards.DefaultSchema).
		WithVersion(cards.Version12)

	s, err := c.StringIndent("", "    ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(s)
	log.Println("Done")
}
