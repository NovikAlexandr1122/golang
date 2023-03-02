// Найдите самое большее число на отрезке от a до b, кратное 7 .

// Входные данные
// Вводится два целых числа a и b (a≤b).

package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)
	switch {
	case a == b:
		fmt.Println("NO")
	case a < b:
		for i := b; i >= a; i-- {
			if i%7 == 0 {
				fmt.Printf("%d", i)
				break
			}
			if i < a+1 {
				fmt.Printf("NO")
				break
			}
		}
	}
}
