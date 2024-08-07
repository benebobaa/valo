package main

import (
	"fmt"

	"github.com/benebobaa/valo"
)

type User struct {
	Name     string    `valo:"notblank,sizeMin=2,sizeMax=50"`
	Email    string    `valo:"email"`
	Age      int       `valo:"min=1,max=120"`
	Friends  []string  `valo:"notnil,sizeMin=1"`
	Address  Address   `valo:"valid"`
	Products []Product `valo:"sizeMin=1,valid"`
}

type Address struct {
	Street  string `valo:"notblank"`
	City    string `valo:"notblank"`
	Country string `valo:"notblank"`
}

type Product struct {
	Name     string `valo:"notblank"`
	Quantity int
}

func main() {
	user := User{
		Name:    "Okay",
		Age:     1,
		Email:   "bene@beneboba.me",
		Friends: []string{""},
		Address: Address{
			Street:  "Jalan Depan Gang",
			City:    "Kokas Macet",
			Country: "Jawa",
		},
		Products: []Product{
			{
				Name:     "",
				Quantity: 0,
			},
		},
	}

	if err := valo.Validate(user); err != nil {
		fmt.Println("Validation error:", err)
	} else {
		fmt.Println("Validation passed")
	}
}
