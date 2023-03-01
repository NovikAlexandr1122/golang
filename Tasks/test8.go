// Количество нулей
// По данным числам, определите количество чисел, которые равны нулю.

// Входные данные
// Вводится натуральное число N, а затем N чисел.

package main

import "fmt"

func main() {
	var a, b, result int
	fmt.Scan(&a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b)
		if b == 0 {
			result++
		}
	}
	fmt.Printf("%d", result)
}
