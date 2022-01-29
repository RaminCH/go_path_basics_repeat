package main 

import "log"

func main() {
	names := []string{"sucdip", "sevendik", "kombayn", "gulchokhra"}

	log.Println(names)

	log.Println(names[1:3])
}

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Other Data Structures_Maps and Slices> go run main6.go
// 2021/12/17 21:39:42 [sucdip sevendik kombayn gulchokhra]
// 2021/12/17 21:39:42 [sevendik kombayn]