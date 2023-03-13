//Часовая стрелка повернулась с начала суток на d градусов. Определите, сколько сейчас целых часов h и целых минут m.

//Входные данные

//На вход программе подается целое число d (0 < d < 360).

//Выходные данные

//Выведите на экран фразу:

//It is ... hours ... minutes.

//Вместо многоточий программа должна выводить значения h и m, отделяя их от слов ровно одним пробелом.

package main

import "fmt"

func main() {
	var h int

	fmt.Print("Введите колличесвто градусов поворота стрелки: ")
	fmt.Scan(&h)

	hours := h / 30
	min := 2 * (h % 30)
	fmt.Println("It is ", hours, "hours ", min, "minutes.")
}