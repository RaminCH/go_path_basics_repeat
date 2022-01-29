package main

import "log"

func main() {
	animals := make(map[string]string)
	animals["dog"] = "Fido"
	animals["cat"] = "Fluffy"

	for _, animal := range animals {
		log.Println(animal)
	}
}

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Loops and ranging over data> go run main3.go
// 2021/12/31 13:00:05 Fido
// 2021/12/31 13:00:05 Fluffy