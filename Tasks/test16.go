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
