package main

import "log"

func main() {
	var myVar bool
	myVar = false

	if myVar == true {
		log.Println("myVar is:", myVar)
	} else {
		log.Println("myVar is:", myVar)
	}
}

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Decision structures> go run .\main2.go
// 2021/12/26 12:37:46 myVar is: false