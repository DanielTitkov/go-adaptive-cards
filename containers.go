package cards

import (
	"errors"
	"fmt"
)

// Container groups items together.
type Container struct {
	Type  string `json:"type"`  // required
	Items []Node `json:"items"` // required
}

func (n Container) validate() error {
	if n.Type != ContainerType {
		return fmt.Errorf("type must be %s", ContainerType)
	}
	if len(n.Items) < 1 {
		return errors.New("container must have elements")
	}
	for _, node := range n.Items {
		if err := node.validate(); err != nil {
			return err
		}
	}
	return nil
}
