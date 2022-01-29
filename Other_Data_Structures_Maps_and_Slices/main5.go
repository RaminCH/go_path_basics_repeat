package main 

import "log"

func main() {
	numbers := []int{1,2,3,4,5,6,7,8,9}

	log.Println(numbers)

	log.Println(numbers[6:9])
}

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Other Data Structures_Maps and Slices> go run main5.go
// 2021/12/17 21:36:06 [1 2 3 4 5 6 7 8 9]
// 2021/12/17 21:36:06 [7 8 9]