package main

import (
	"log"
)

func main() {
	var myString string
	myString = "Green"

	log.Println("myString is set to: ", myString)

	changeUsingPointer(&myString)
	log.Println("after func call myString is set to: ", myString)
}

func changeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue
}

// PS C:\Users\Ramin\go\src\learning-go\Pointers> go run main.go
// 2021/12/07 21:06:45 myString is set to:  Green
// 2021/12/07 21:06:45 s is set to 0xc00005a230
// 2021/12/07 21:06:45 after func call myString is set to:  Red
