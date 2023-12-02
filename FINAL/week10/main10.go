package main

import (
	"FINAL/week10/src/greeting"
	"fmt"
	"week10/src/mymath"
)

func main() {
	fmt.Println(mymath.MyPower(2, 9))
	fmt.Println(mymath.MyAbs(-99))
	greeting.Hello()
	fmt.Println(mymath.MyAbs(17))
	greeting.Hi()
}

// package main

// import "week10/src/greeting.go"
// func main() {
// 	greeting.Hello()
// 	greeting.Hi()
// }
