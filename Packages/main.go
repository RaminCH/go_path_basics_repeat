package main

import (
	"github.com/ramin/myniceprogram/helpers"	
	"log"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType	//here starts step2

	myVar.TypeName = "Some Name"
	log.Println(myVar.TypeName)
}



// Step1 (before creating helpers folder)

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Packages> go mod init github.com/ramin/myniceprogram    (<- enter this line)
// go: creating new go.mod: module github.com/ramin/myniceprogram
// go: to add module requirements and sums:
//         go mod tidy								(<- will appear go mod file)

// Step2 (after creating helper folder and helepr.go)
// PS C:\Users\Ramin\go\src\learning-go\learning-go\Packages> go run main.go
// 2022/01/11 17:30:41 Hello
// 2022/01/11 17:30:42 Some Name