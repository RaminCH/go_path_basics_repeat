package main 

import "log"

type customer struct {
	Name string
	LastName string
}

func (c *customer) printName() string {
	return c.Name
}

func main() {

	var number1 customer
	number1.Name = "John"
	number1.LastName = "Jinss"

	number2 := customer{
		Name: "Harry",
		LastName: "Micks",
	}

	log.Println("customer1 is:", number1.printName())
	log.Println("customer1 is:", number2.printName())
}


// PS C:\Users\Ramin\go\src\learning-go\learning-go\Receivers_StructsWithFunctions> go run main1.go
// 2021/12/17 20:58:45 customer1 is: John
// 2021/12/17 20:58:45 customer1 is: Harry