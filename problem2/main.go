package main

import "fmt"

func main() {
	left, right := 1, 1
	sum := 0

	for right <= 4_000_000 {
		left, right = right, left+right
		if right%2 == 0 {
			sum += right
		}
	}
	fmt.Println(sum)
}
