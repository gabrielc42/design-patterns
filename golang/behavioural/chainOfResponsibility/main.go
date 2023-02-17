// // chain of responsibility:
// processing objects responsible for a command,
// forwards command to next processor, etc.

// decoupling, the client doesn't choose the reciever
// because more than one object can handle request

package main

func main() {
	cashier := &cashier{}

	medical := &medical{}
	medical.setNext(cashier)

	doctor := &doctor{}
	doctor.setNext(medical)

	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{name: "Mr. Patient"}
	reception.execute(patient)
}
