package main 

import "log"

var s = "seven"

func main()  {
	var s2 = "six"
	
	// s := "eight"

	log.Println("s is: ", s)
	log.Println("s2 is: ", s2)

	saySomething("xxx")
}

func saySomething(s3 string) (string, string) {
	log.Println("s from the saySomething func is: ", s)
	return s3, "World"
}

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Types and Structs> go run main1.go
// 2021/12/08 22:04:37 s is:  seven
// 2021/12/08 22:04:37 s2 is:  six
// 2021/12/08 22:04:37 s from the saySomething func is:  seven