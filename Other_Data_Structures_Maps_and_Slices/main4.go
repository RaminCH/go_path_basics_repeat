package main

import (
	"log"
	"sort"
)

func main() {
	var mySlice []int
	
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 3)
	
	sort.Ints(mySlice)

	log.Println(mySlice)
}

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Other Data Structures_Maps and Slices> go run main4.go
// 2021/12/17 21:32:30 [1 2 3]