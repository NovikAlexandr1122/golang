// Из натурального числа удалить заданную цифру.

// Входные данные

// Вводятся натуральное число и цифра, которую нужно удалить.

// Выходные данные

// Вывести число без заданных цифр.

// Sample Input:

// 38012732
// 3
// Sample Output:

// 801272

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	c := strconv.Itoa(a)
	d := strconv.Itoa(b)
	i := strings.Replace(c, d, "", -1)
	fmt.Println(i)

}
