package main

import (
	"awesomeProject/app/patterns/dof/conceptual/memento/model"
	_ "awesomeProject/app/patterns/dof/conceptual/memento/model"
	"fmt"
)

func main() {
	caretaker := &model.Caretaker{MementoArray: make([]*model.Memento, 0)}
	originator := model.Originator{State: "A"}
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState("B")
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.SetState("C")
	fmt.Printf("Originator Current State: %s\n", originator.GetState())
	caretaker.AddMemento(originator.CreateMemento())

	originator.RestoreMemento(caretaker.GetMemento(1))
	fmt.Printf("Restored to State: %s\n", originator.GetState())

	originator.RestoreMemento(caretaker.GetMemento(0))
	fmt.Printf("Restored to State: %s\n", originator.GetState())
}
