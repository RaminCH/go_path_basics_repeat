package main

import "log"

func main() {
	myVar := "cat"

	switch myVar {
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("cat is set to dog")
	case "fish":
		log.Println("cat is set to fish")

	default: 
		log.Println("cat is set to default")
	}
}

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Decision structures> go run .\main6.go
// 2021/12/26 13:05:09 cat is set to cat