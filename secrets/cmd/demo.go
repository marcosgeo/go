package main

import (
	"fmt"

	"github.com/marcosgeo/go/secrets"
)

func main() {
	v := secrets.File("my-fake-key", "./secrets/secrets")
	// err := v.Set("demo_key1", "123 some crazy value")
	// if err != nil {
	// 	panic(err)
	// }
	// err = v.Set("demo_key2", "456 some crazy value")
	// if err != nil {
	// 	panic(err)
	// }
	// err = v.Set("demo_key3", "789 some crazy value")
	// if err != nil {
	// 	panic(err)
	// }
	plain, err := v.Get("demo_key1")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain:", plain)
	plain, err = v.Get("demo_key2")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain:", plain)
	plain, err = v.Get("demo_key3")
	if err != nil {
		panic(err)
	}
	fmt.Println("Plain:", plain)
}
