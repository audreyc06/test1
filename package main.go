package main

import "fmt"

func Multiply(n int) int {
	result := n * 2
	return result
}

func main() {
	fmt.Println(Multiply(3))
}
