// Количество минимумов
// Найдите количество минимальных элементов в последовательности.

// Входные данные

// Вводится натуральное число N, а затем N целых чисел последовательности.

package main

import (
	"fmt"
	"sort"
)

func main() {
	var a, result int

	fmt.Scan(&a)
	b := make([]int, a, a)
	for i := 0; i < a; i++ {
		fmt.Scan(&b[i])
	}
	sort.Ints(b)
	min := b[0]

	for _, elem := range b {
		if min == elem {
			result++
		}
	}
	fmt.Printf("%d", result)

	//fmt.Printf("%d", result)
}
