//Дано натуральное число, выведите его последнюю цифру.

//Формат входных данных
//На вход дается натуральное число N, не превосходящее 10000.

//Формат выходных данных
//Выведите одно целое число - ответ на задачу.

//Sample Input:123
//Sample Output:3

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a int
	fmt.Scan(&a)
	switch {
	case a < 999999 || a != 0 || a >= 100000:
		b := strconv.Itoa(a)
		fmt.Println(string(b[0]))

	default:
		fmt.Println("Число должно быть целым и не отрицательным.")
	}

}
