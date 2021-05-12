package main

import "fmt"

type item struct {
	observerList []Observer
	name         string
	inSock       bool
}

func newItem(name string) *item {
	return &item{name: name}
}

func (receiver *item) updateAvailability() {
	fmt.Println("Item %s is now in stock\n", receiver.name)
	receiver.inSock = true
	receiver.notifyAll()
}

func (i *item) register(o Observer) {
	i.observerList = append(i.observerList, o)
}

func (i *item) deregister(o Observer) {
	i.observerList = removeFromlice(i.observerList, o)
}

func removeFromlice(observerList []Observer, observerToRemove Observer) []Observer {
	observerListLength := len(observerList)
	for i, observer := range observerList {
		if observerToRemove.getID() == observer.getID() {
			observerList[observerListLength-1], observerList[i] = observerList[i], observerList[observerListLength-1]
			return observerList[:observerListLength-1]
		}
	}
	return observerList
}

func (receiver *item) notifyAll() {
	for _, observer := range receiver.observerList {
		observer.update(receiver.name)
	}
}
