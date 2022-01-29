package main

import "log"

func main() {
	myMap := make(map[string]string)

	myMap["dog"] = "Samson"
	myMap["other Dog"] = "Cassie"

	// myMap["dog"] = "Fido"


	log.Println(myMap["dog"])
	log.Println(myMap["other Dog"])
}

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Other Data Structures_Maps and Slices> go run main.go
// 2021/12/17 21:08:09 Samson
// 2021/12/17 21:08:09 Cassie
// PS C:\Users\Ramin\go\src\learning-go\learning-go\Other Data Structures_Maps and Slices> go run main.go
// 2021/12/17 21:08:38 Fido
// 2021/12/17 21:08:38 Cassie