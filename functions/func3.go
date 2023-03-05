package main

import "fmt"

func main() {
	a, b := sumInt(1, 0)
	fmt.Println(a, b)
}

func sumInt(num ...int) (int, int) {
	sum := 0
	i := 0
	for _, j := range num {
		sum += j
		i++
	}
	return i, sum
}
