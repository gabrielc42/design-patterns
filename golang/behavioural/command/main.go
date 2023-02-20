// command:
// encapsulates what method to call,
// the method's arguments,
// and the object to which the method belongs
// for performing a given action

package main

func main() {
	tv := &tv{}
	onCommand := &onCommand{
		device: tv,
	}
	offCommand := &offCommand{
		device: tv,
	}
	onButton := &button{
		command: onCommand,
	}
	onButton.press()
	offButton := &button{
		command: offCommand,
	}
	offButton.press()
}
