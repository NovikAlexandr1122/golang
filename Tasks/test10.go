// Цифровой корень
// Цифровой корень натурального числа — это цифра, полученная в результате итеративного процесса суммирования цифр, на каждой итерации которого для подсчета суммы цифр берут результат, полученный на предыдущей итерации. Этот процесс повторяется до тех пор, пока не будет получена одна цифра.

// Например цифровой корень 65536 это 7 , потому что 6+5+5+3+6=25 и 2+5=7 .

// По данному числу определите его цифровой корень.

// Входные данные

// Вводится одно натуральное число n, не превышающее
// 1
// 0
// 7
// 10
// 7
//  .

// Выходные данные
// Вывести цифровой корень числа n.

// Sample Input:

// 3456
// Sample Output:

// 9
package main

import (
	"fmt"
)

func main() {
	var a, sum, result int
	fmt.Scan(&a)

	for a > 0 {
		sum = a % 10
		//fmt.Println("sum", sum)
		a /= 10
		//fmt.Println("a", a)
		result += sum
		//fmt.Println("result", result)
	}
	b := result
	sum = b % 10
	//fmt.Println("sum", sum)
	b /= 10
	//fmt.Println("result", result)
	b += sum

	fmt.Println(b)
}
