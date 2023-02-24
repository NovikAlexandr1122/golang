package main

import "fmt"

func main() {
	array := [5]int{}
	var a int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a)
		array[i] = a
	}
	// здесь ваш код
	// ...
	var max int
	for c := 0; c < len(array); c++ {
		b := array[c]

		if max < b {
			max = b
		}
		//fmt.Println(max)
	}
	fmt.Println(max)
}
