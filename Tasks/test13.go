// По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.

// Входные данные

// Вводится натуральное число.

package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Printf("1 ")
	for i := 2; i <= a; {
		fmt.Printf("%d ", i)
		i = i + i
	}
}
