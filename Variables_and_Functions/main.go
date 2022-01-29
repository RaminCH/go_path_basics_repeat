package main

import "fmt"

func main() {

	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println("The function returned",whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "Something", "else"
}


// PS C:\Users\Ramin\go\src\learning-go\learning-go\Variables and Functions> go run main.go
// The function returned Something else