package main

import "log"

func main() {
	var mySlice []string
	
	mySlice = append(mySlice, "John")
	mySlice = append(mySlice, "Trevor")
	

	log.Println(mySlice)
}

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Other Data Structures_Maps and Slices> go run main3.go
// 2021/12/17 21:25:51 [John Trevor]