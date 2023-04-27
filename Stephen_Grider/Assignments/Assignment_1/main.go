package main

import "fmt"

func main() {
	number := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, num := range number {
		if number[i]%2 == 0 {
			fmt.Println("Number", num, "is even")
		} else {
			fmt.Println("Number", num, "is odd")
		}
	}
}
