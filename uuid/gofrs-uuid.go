package main

import (
	"fmt"
	"github.com/gofrs/uuid"
	"log"
)

func main() {
	v1, err := uuid.NewV1()
	if err != nil {
		log.Fatal("cannot generate v1 uuid")
	}
	fmt.Printf("v1 uuid: %v\n", v1)
	//gofrs does not support version 2 uuid

	// version 3 uuid
	v3 := uuid.NewV3(uuid.NamespaceURL, "https://example.com")
	fmt.Printf("v3 uuid: %v\n", v3)

	// version 4 uuid
	v4, err := uuid.NewV4()
	if err != nil {
		log.Fatal("cannot generate v4 uuid")
	}
	fmt.Printf("v4 uuid: %v\n", v4)

	// version 5 uuid
	v5 := uuid.NewV5(uuid.NamespaceURL, "https://example.com")
	fmt.Printf("v5 uuid: %v\n", v5)
}
