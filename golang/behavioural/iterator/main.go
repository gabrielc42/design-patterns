// iterator:
// iterates through objects

package main

import "fmt"

func main() {
	user1 := &user{
		name: "Toma",
		age:  25,
	}
	user2 := &user{
		name: "Wangari",
		age:  24,
	}
	userCollection := &userCollection{
		users: []*user{user1, user2},
	}
	iterator := userCollection.createIterator()
	for iterator.hasNext() {
		user := iterator.getNext()
		fmt.Printf("User is %+v\n ", user)
	}
}
