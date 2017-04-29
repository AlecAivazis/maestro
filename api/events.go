package main

import (
	"fmt"

	"github.com/nautilus/events"
)

func (s *MaestroAPI) HandleAction(a *events.Action) {
	fmt.Println(a.Type)
}
