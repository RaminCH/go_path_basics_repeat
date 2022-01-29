package main

import "log"

func main() {
	animals := []string{"dog","fish","horse","cow","cat"}

	for i, animal := range animals {
		log.Println(i, animal)
	}
}

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Loops and ranging over data> go run main1.go
// 2021/12/31 12:43:39 0 dog
// 2021/12/31 12:43:39 1 fish 
// 2021/12/31 12:43:39 2 horse
// 2021/12/31 12:43:39 3 cow  
// 2021/12/31 12:43:39 4 cat 