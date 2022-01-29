package main

import (
	"encoding/json"
	"fmt"
	"log"
)


//creating struct from Json
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
	[
		{
			"first_name": "Clark",
			"last_name": "Kent",
			"hair_color": "black",
			"has_dog": true
		},
		{
			"first_name": "Bruce",
			"last_name": "Wayne",
			"hair_color": "black",
			"has_dog": false
		}
	]`

	var unmarshalled []Person

	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}

	log.Printf("unmarshalled: %v", unmarshalled)   //unmarshalled is: [{Clark Kent black true} {Bruce Wayne black false}]

	//write json from the struct
	var mySlice []Person

	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	var m2 Person
	m2.FirstName = "Diana"
	m2.LastName = "Prince"
	m2.HairColor = "black"
	m2.HasDog = false

	mySlice = append(mySlice, m2)

	newJson, err := json.MarshalIndent(mySlice, "", "     ")
	if err != nil {
		log.Println("error marshaling", err)
	}

	fmt.Println(string(newJson))
}

//Step1
// PS C:\Users\Ramin\go\src\learning-go\Reading and rinting JSON> go run main.go
// 2022/01/15 16:30:22 unmarshalled: [{Clark Kent black true} {Bruce Wayne black false}]


//Step2
// PS C:\Users\Ramin\go\src\learning-go\Reading and rinting JSON> go run main.go
// 2022/01/15 16:48:52 unmarshalled: [{Clark Kent black true} {Bruce Wayne black false}]		(from Step1 - //creating struct from Json)
// [																							(from step2 - //write json from the struct)
//      {
//           "first_name": "Wally",
//           "last_name": "West",  
//           "hair_color": "red",  
//           "has_dog": false      
//      },
//      {
//           "first_name": "Diana",
//           "last_name": "Prince",
//           "hair_color": "black",
//           "has_dog": false
//      }
// ]