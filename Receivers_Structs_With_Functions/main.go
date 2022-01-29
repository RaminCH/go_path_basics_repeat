package main

import "log"

type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	var myVar myStruct
	myVar.FirstName = "John"


	myVar2 := myStruct{
		FirstName: "Mary",
	}

	log.Println("myVar is set to:", myVar.FirstName)
	log.Println("myVar2 is set to:", myVar2.FirstName)

	//-----------------------------------

	log.Println("myVar is set to:", myVar.printFirstName())
	log.Println("myVar2 is set to:", myVar2.printFirstName())
}

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Receivers_StructsWithFunctions> go run main.go
// 2021/12/17 20:46:05 myVar is set to: John
// 2021/12/17 20:46:05 myVar2 is set to: Mary
// 2021/12/17 20:46:05 myVar is set to: John 
// 2021/12/17 20:46:05 myVar2 is set to: Mary