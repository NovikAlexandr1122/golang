// Напишите программу, которая выводит квадраты натуральных чисел от 1 до 10. Квадрат каждого числа должен выводится в новой строке.
package main

import "fmt"

func main() {
	i := 1
	for i < 11 {
		fmt.Println(i * i)
		i++
	}
}