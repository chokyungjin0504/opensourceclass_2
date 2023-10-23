package main

import (
	"fmt"
)

func main() {
	a := 10
	var b int = 20
	var pa *int = &a

	fmt.Printf("%T %T\n", &a, pa)
	fmt.Printf("%x %x %x\n", &a, pa, &pa)
	fmt.Println(&a, pa, &pa)
	fmt.Println(*pa)
	pa = &b
	fmt.Println(*pa)
}

// package main

// import "fmt"

// func swap(n1 *int, n2 *int) {
// 	temp := *n1
// 	*n1 = *n2
// 	*n2 = temp
// }

// func main() {
// 	a := 10
// 	b := 20

// 	c := &a
// 	fmt.Printf("%T\n", c)
// 	fmt.Println(a, b)
// 	swap(&a, &b)
// 	fmt.Println(a, b)
// }

// func double(n *int) {
// 	*n = *n * 2
// }

// func main() {
// 	var a int = 6
// 	double(&a)
// 	fmt.Println(a)
// }
