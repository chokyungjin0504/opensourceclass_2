package main

import "fmt"

func sum(n1 int, n2 int) {
	temp := n1
	n1 = n2
	n2 = temp
}

func main() {
	a := 10
	//b := 20 //var b int == 20

	c := &a
	fmt.Printf("%T\n", c)
	fmt.Println(a, b)
	//swap(a,b)
	fmt.Println(a, b)
}

// func double(n *int) {
// 	*n = *n * 2
// }

// func main() {
// 	var a int = 6
// 	double(&a)
// 	fmt.Println(a)
// }
