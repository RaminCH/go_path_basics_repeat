package main

import(
	"time"
	"log"	
) 

type User struct {
	FirstName 	string			//Capital letters are used to let to use the out of this code, small letters make variables private!
	LastName 	string
	PhoneNumber string
	Age			int
	BirthDate 	time.Time
}

func main() {
	user := User{
		FirstName: "Ramin",
		LastName:  "Chopurov",
	}

	log.Println("\n","name:", user.FirstName,"\n", "lastname:", user.LastName,"\n", "birth date is:", user.BirthDate)
}

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Types and Structs> go run main.go
// 2021/12/08 22:29:18
//  name: Ramin
//  lastname: Chopurov
//  birth date is: 0001-01-01 00:00:00 +0000 UTC