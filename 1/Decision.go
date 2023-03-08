//По данному целому числу найти его квадрат(на входе целое число, на выходе его квадрат)

package main

import "fmt"

func main() {
	var a int
	fmt.Print("Hi. This program squares the number that you specify. Specify any integer:")
	fmt.Scan(&a)

	b := a * a

	fmt.Println(b)

}
