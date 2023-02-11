package main

func main() {
	item := newItem("Green Tea")
	observerFirst := &customer{id: "me@mail.pigeon"}
	observerSecond := &customer{id: "you@mail.pigeon"}
	item.register(observerFirst)
	item.register(observerSecond)
	item.updateAvailability()
}
