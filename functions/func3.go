// Напишите "функцию голосования", возвращающую то значение (0 или 1), которое среди значений ее аргументов x, y, z встречается чаще.

// Входные данные

// Вводится 3 числа - x, y и z (x, y и z равны 0 или 1).

// Выходные данные

// Необходимо вернуть значение функции от x, y и z.

// Sample Input:

// 0 0 1
// Sample Output:

// 0

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
