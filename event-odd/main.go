package main

import "fmt"

func main() {
	odd, even := oddEven(100)
	fmt.Println(odd)
	fmt.Println(even)
}

func oddEven(number int) ([]int, []int) {
	var odd, even []int
	for i := 1; i <= number; i++ {
		if i%2 == 1 {
			odd = append(odd, i)
		} else {
			even = append(even, i)
		}
	}
	return odd, even
}
