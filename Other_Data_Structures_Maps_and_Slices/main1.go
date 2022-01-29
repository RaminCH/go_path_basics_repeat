package main

import "log"

func main() {
	myMap := make(map[string]int)

	myMap["First"] = 1
	myMap["Second"] = 2

	log.Println(myMap["First"],"\n", myMap["Second"])
}

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Other Data Structures_Maps and Slices> go run main1.go
// 2021/12/17 21:13:18 1 
//  2