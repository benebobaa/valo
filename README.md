# Valo: Validator in Go

Valo is a simple and lightweight validation library for Go. It allows you to define validation rules using struct tags and validate your data structures easily. This library supports various validation rules such as `notblank`, `sizeMin`, `sizeMax`, `min`, `max`, `notnil`, and recursive struct validation with `valid`.

## Features

- Validate struct fields with simple tags
- Supports basic validation rules: `notblank`, `sizeMin`, `sizeMax`, `min`, `max`, `notnil`
- Validation for nested structs
- Easy to use and integrate

## Installation

To install Valo, use the following command:

```sh
go get -u github.com/benebobaa/valo
```


## Quick Start

Here's a quick example to get you started:

```go
package main

import (
	"fmt"
	"valo"
)

type User struct {
	Name    string   `valo:"notblank,sizeMin=2,sizeMax=50"`
	Age     int      `valo:"min=18,max=120"`
	Friends []string `valo:"notnil"`
	Address Address  `valo:"valid"`
}

type Address struct {
	Street  string `valo:"notblank"`
	City    string `valo:"notblank"`
	Country string `valo:"notblank"`
}

func main() {
	user := User{
		Name: "Okay",
		Age:  30,
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
```

# Validation Rules

Valo supports the following validation rules, which can be specified using struct tags:

- **`notblank`**: Ensures the string is not empty or only whitespace.
- **`sizeMin=N`**: Ensures the string has at least N characters.
- **`sizeMax=N`**: Ensures the string has no more than N characters.
- **`min=N`**: Ensures the integer is at least N.
- **`max=N`**: Ensures the integer is at most N.
- **`notnil`**: Ensures the pointer or slice is not nil.
- **`valid`**: Recursively validates nested structs.

## Note

This library is created for experimental learning purposes. It is not intended for production use.

## Custom Validation

You can also define custom validation functions and integrate them with Valo. (Coming soon :D)

## License

This project is licensed under the MIT License. See the LICENSE file for details.
