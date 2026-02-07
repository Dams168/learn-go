package main

import "errors"

var (
	ValidationError = errors.New("Validation Error")
	NotFoundError   = errors.New("Not Found Error")
)

func GetById(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "john" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := GetById("john")

	if err != nil {
		if err == ValidationError {
			println("Validation error occurred")
		} else if err == NotFoundError {
			println("Not found error occurred")
		} else {
			println("An unknown error occurred")
		}
	} else {
		println("No error, user found")
	}
}
