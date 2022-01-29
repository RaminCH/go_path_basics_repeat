package main

import "log"

func main() {
	cat := "cat2"

	if cat == "cat" {
		log.Println("cat is:", cat)
	} else {
		log.Println("cat is not cat")
	}
}

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Decision structures> go run .\main3.go       (cat := "cat")
// 2021/12/26 12:41:45 cat is: cat
// PS C:\Users\Ramin\go\src\learning-go\learning-go\Decision structures> go run .\main3.go		 (cat := "cat2")
// 2021/12/26 12:41:55 cat is not cat