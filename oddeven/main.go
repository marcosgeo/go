package main

import "fmt"

func main(){
	fmt.Println()
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i+1)
	}

	for _, n := range numbers {
		if n % 2  == 0 {
			fmt.Printf("%v is even\n", n)
		} else {
			fmt.Printf("%v is odd\n", n)
		}
	}
}