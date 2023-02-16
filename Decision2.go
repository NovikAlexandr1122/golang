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

	var p int
	fmt.Print("Ввеите число от 1-10000: ")
	fmt.Scan(&p)
	if p > 10000 {
		fmt.Println("Число должно быть меньше 10000!")
		return
	} else {

		a := strconv.Itoa(p)
		b := len(a)
		fmt.Println(string(a[b-1]))
	}

}
