// Вклад в банке составляет x рублей. Ежегодно он увеличивается на p процентов, после чего дробная часть копеек отбрасывается. Каждый год сумма вклада становится больше. Определите, через сколько лет вклад составит не менее y рублей.

// Входные данные

// Программа получает на вход три натуральных числа: x, p, y.

// Выходные данные

// Программа должна вывести одно целое число.

// Sample Input:

// 100
// 10
// 200
// Sample Output:

// 8

package main

import "fmt"

func main() {
	var x, p, y, i float32
	fmt.Scan(&x, &p, &y)
	for x <= y {
		if p == 1 {
			x = x + (x / 100)
			i++
		} else {
			x = x + (x*p)/100
			i++
			//fmt.Println(i)
		}
	}
	fmt.Println(i)
}
