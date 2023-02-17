/*На ввод подается целое число. Если число положительное - вывести сообщение "Число положительное", если число отрицательное - "Число отрицательное". Если подается ноль - вывести сообщение "Ноль". Выводить сообщение без кавычек.

Sample Input:5
Sample Output:Число положительное*/

package main

import "fmt"

func main() {

	var a int32 //для экономии памяти
	fmt.Print("Enter your number:")
	fmt.Scan(&a)
	switch {
	case a < 0:
		fmt.Println("Число отрицательное.")
	case a > 0:
		fmt.Println("Число положительное.")
	case a == 0:
		fmt.Println("Ноль.")

	default:
		fmt.Println("Вам нужно ввести целое число.")
	}
}
