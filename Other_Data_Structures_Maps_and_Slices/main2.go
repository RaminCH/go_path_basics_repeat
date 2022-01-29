package main

import "log"

type User struct {
	FirstName string
	LastName string
}

func main() {
	myMap := make(map[string]User)

	me := User{
		FirstName: "John",
		LastName: "Smith",
	}

	myMap["me"] = me

	log.Println(myMap["me"].FirstName)
}

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Other Data Structures_Maps and Slices> go run main2.go
// 2021/12/17 21:19:15 John