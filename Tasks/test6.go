package main

import "fmt"

func main() {
	var a, b float32
	fmt.Scan(&a, &b)
	if a > 0 && b > 0 {
		c := (a + b) / 2
		fmt.Println(c)
	}
}
