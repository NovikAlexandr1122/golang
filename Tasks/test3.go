package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	if a < 86400 && a >= 0 {
		h := a / 3600
		m := (a - (h * 3600)) / 60

		fmt.Printf("It is %d hours %d minutes.", h, m)
	}
}
