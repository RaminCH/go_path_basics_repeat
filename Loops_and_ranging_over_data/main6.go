package main

import "log"

func main() {
	type User struct {
		FirstName 	string
		LastName 	string
		Email 		string
		Age 		int
	}

	var users []User
	users = append(users, User{"John", "Smith", "j.smith@gmail.com", 30})
	users = append(users, User{"Mary", "Jones", "m.jones@gmail.com", 20})
	users = append(users, User{"Sally", "Brown", "s.brown@gmail.com", 45})
	users = append(users, User{"Alex", "Anderson", "a.anderson@gmail.com", 17})


	for _, l := range users {
		log.Println(l.FirstName, l.LastName, l.Email, l.Age)
	}
}

// PS C:\Users\Ramin\go\src\learning-go\learning-go\Loops and ranging over data> go run main6.go
// 2021/12/31 13:16:31 John Smith j.smith@gmail.com 30
// 2021/12/31 13:16:31 Mary Jones m.jones@gmail.com 20      
// 2021/12/31 13:16:31 Sally Brown s.brown@gmail.com 45     
// 2021/12/31 13:16:31 Alex Anderson a.anderson@gmail.com 17