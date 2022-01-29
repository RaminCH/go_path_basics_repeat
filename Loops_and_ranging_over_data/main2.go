package main

import "log"

func main() {
	animals := []string{"dog","fish","horse","cow","cat"}

	for _, animal := range animals {
		log.Println(animal)
	}
}

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Loops and ranging over data> go run main2.go
// 2021/12/31 12:50:34 dog
// 2021/12/31 12:50:34 fish 
// 2021/12/31 12:50:34 horse
// 2021/12/31 12:50:34 cow  
// 2021/12/31 12:50:34 cat 