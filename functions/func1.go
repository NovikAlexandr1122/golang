// Напишите функцию, находящую наименьшее из четырех введённых в этой же функции чисел.

// Входные данные

// Вводится четыре числа.

// Выходные данные

// Необходимо вернуть из функции наименьшее из 4-х данных чисел

// Sample Input:

// 4 5 6 7
// Sample Output:

// 4

package main

import "fmt"

func main() {
	a := minimumFromFour()
	fmt.Println(a)
}
func minimumFromFour() int {
	//print your code here
	var a, b, c, d, minNumber int
	fmt.Scan(&a, &b, &c, &d)

	for i := 0; i <= a && i <= b && i <= c && i <= d; i++ {
		minNumber = i
	}
	return minNumber

}
