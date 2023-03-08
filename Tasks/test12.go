// По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.

// Входные данные

// Вводится натуральное число.

// Выходные данные

// Выведите ответ на задачу.

// Sample Input:

// 50
// Sample Output:

// 1 2 4 8 16 32

package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	switch {
	case a == 1 || a == 21 || a == 31 || a == 41 || a == 51 || a == 61 || a == 71 || a == 81 || a == 91:
		fmt.Printf("%d korova", a)
	case a/10 != 1 && (a%10 == 2 || a%10 == 3 || a%10 == 4):
		fmt.Printf("%d korovy", a)
	default:
		fmt.Printf("%d korov", a)
	}
}
