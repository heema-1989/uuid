package main

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
)

func main() {
	var err error
	uuid.Must(uuid.NewV4(), err)

	fmt.Println("Generated uuid is: ")
	//parsing uuid from string
	u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s", u2)
}
