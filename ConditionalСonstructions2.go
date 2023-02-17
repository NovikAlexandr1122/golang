/*Выведите "YES", если все цифры числа различны, в противном случае - "NO".

Sample Input 1:

237
Sample Output 1:

YES
Sample Input 2:

117
Sample Output 2:

NO*/

package main

import "fmt"

func main() {

	var a int
	fmt.Scan(&a)

	f := a / 100
	s := (a % 100) / 10
	t := a % 10

	switch {
	case f != s && s != t && t != f:
		fmt.Println("YES")
	default:
		fmt.Println("NO")
	}
	//fmt.Println(f)
	//fmt.Println(s)
	//fmt.Println(t)

}
