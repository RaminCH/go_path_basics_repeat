package main

import "log"

func main() {
	myNum := 100
	isTrue := false

	if myNum > 99 && !isTrue {
		log.Println("myNum is greater that 99 and is set to true")
	} else if myNum > 1000 || isTrue == true {
		log.Println("1")
	}
}