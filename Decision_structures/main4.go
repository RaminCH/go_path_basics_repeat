package main

import "log"

func main() {
	myNum := 100
	isTrue := false

	if myNum > 99 && isTrue == false {
		log.Println("myNum is greater that 99 and is set to true")
	}
}

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Decision structures> go run .\main4.go
// 2021/12/26 12:51:29 myNum is greater that 99 and is set to true