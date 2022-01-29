package main

import (
	"log"

	"github.com/ramin/chanelsprogram/helper" //Don't pay attention to red mark, its working
)

const numPool = 1000

func CalculateValue(intChan chan int) {
	randomNumber := helper.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalculateValue(intChan)

	num := <-intChan
	log.Println(num)
}

// Step1
// create main.go , helper folder and in it helper.go then
// PS C:\Users\Ramin\go\src\learning-go\Chanels> go mod init github.com/ramin/chanelsprogram
// go: creating new go.mod: module github.com/ramin/chanels
// go: to add module requirements and sums:
//         go mod tidy

//Step2
// PS C:\Users\Ramin\go\src\learning-go> cd .\Chanels\
// PS C:\Users\Ramin\go\src\learning-go\Chanels> go run main.go
// 2022/01/13 23:50:45 1
// PS C:\Users\Ramin\go\src\learning-go\Chanels> go run main.go
// 2022/01/13 23:50:47 1

//After adding rand.Seed... in helper.go

// PS C:\Users\Ramin\go\src\learning-go\Chanels> go run main.go
// 2022/01/13 23:51:49 0
// PS C:\Users\Ramin\go\src\learning-go\Chanels> go run main.go
// 2022/01/13 23:51:52 9
// PS C:\Users\Ramin\go\src\learning-go\Chanels> go run main.go
// 2022/01/13 23:52:02 724
// PS C:\Users\Ramin\go\src\learning-go\Chanels> go run main.go
// 2022/01/13 23:52:05 775
// PS C:\Users\Ramin\go\src\learning-go\Chanels> go run main.go
// 2022/01/13 23:52:09 825
