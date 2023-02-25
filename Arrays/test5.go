// Дана последовательность, состоящая из целых чисел. Напишите программу, которая подсчитывает количество положительных чисел среди элементов последовательности.

// Входные данные

// Сначала задано число
// �
// N — количество элементов в последовательности (
// 1
// ≤
// �
// ≤
// 100
// 1≤N≤100). Далее через пробел записаны
// �
// N чисел — элементы последовательности. Последовательность состоит из целых чисел.

// Выходные данные

// Необходимо вывести единственное число - количество положительных элементов в последовательности.

// Sample Input:

// 5
// 1 2 3 -1 -4
// Sample Output:

// 3
package main

import "fmt"

func main() {
	var a, n, c int
	fmt.Scan(&n)
	if n > 0 && n <= 100 {
		for i := 0; i < n; i++ {
			fmt.Scan(&a)
			if a > 0 {
				c++
			}
		}
		fmt.Printf("%d", c)
	} else {
		fmt.Printf("\nЧисло эл-ов последовательности должно быть больше 0 и меньше 101 а у Вас:%d", n)
	}
}
