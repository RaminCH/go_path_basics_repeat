package main

import "log"

func main() {
	var myVar bool
	myVar = true

	if myVar == true {
		log.Println("myVar is:", myVar)
	} else {
		log.Println("myVar is:", myVar)
	}
}

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Decision structures> go run .\main1.go
// 2021/12/26 12:36:21 myVar is: true