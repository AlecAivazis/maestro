package main

import (
	"fmt"

	"github.com/nautilus/events"
)

func (s *MaestroBuild) HandleAction(a *events.Action) {
	fmt.Println("received action", a.Type)
}
