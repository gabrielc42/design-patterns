// memento:
// implement undoable actions, saves state of object at
// instance and restoring if actions performed need to be undone

package main

import "fmt"

func main() {
	caretaker := &caretaker{
		mementoArray: make([]*memento, 0),
	}
	originator := &originator{
		state: "🌄",
	}
	fmt.Printf("Originator Current State: %s \n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.setState("🌇")
	fmt.Printf("Originator Current State: %s \n", originator.getState())

	caretaker.addMemento(originator.createMemento())
	originator.setState("🌆")

	fmt.Printf("Originator Current State: %s \n", originator.getState())
	caretaker.addMemento(originator.createMemento())

	originator.restoreMemento(caretaker.getMemento(1))
	fmt.Printf("Restored to State: %s 👝\n", originator.getState())
}
