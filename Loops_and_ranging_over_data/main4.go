package main

import "log"

func main() {
	animals := make(map[string]string)
	animals["dog"] = "Fido"
	animals["cat"] = "Fluffy"

	for animalType, animal := range animals {
		log.Println(animalType, animal)
	}
}

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Loops and ranging over data> go run main4.go
// 2022/01/08 00:10:36 dog Fido
// 2022/01/08 00:10:36 cat Fluffy