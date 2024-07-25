package main

import (
	"fmt"

	"github.com/benebobaa/valo"
)

type User struct {
	Name string `valo:"notblank,sizeMin=2,sizeMax=50"`
	Age  int    `valo:"min=1,max=120"`
	// Friends []string `valo:"notnil"`
	Address Address `valo:"valid"`
}

type Address struct {
	Street  string `valo:"notblank"`
	City    string `valo:"notblank"`
	Country string `valo:"notblank"`
}

func main() {
	user := User{
		Name: "Okay",
		Age:  1,
		Address: Address{
			Street:  "Jalan Depan Gang",
			City:    "Kokas Macet",
			Country: "Jawa",
		},
	}

	if err := valo.Validate(user); err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Validation passed")
	}
}
