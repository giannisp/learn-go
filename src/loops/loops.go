package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		if i == 0 {
			break
		}
	}
}
