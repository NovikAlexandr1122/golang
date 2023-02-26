package main

import "fmt"

func main() {
	var a int
	//var b int
	fmt.Scan(&a)
	h := a / 3600
	m := a % 3600
	fmt.Println(h, m)

}
