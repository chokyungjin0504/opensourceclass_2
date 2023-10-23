package main

import (
	"fmt"
	//"reflect"
	"strings"
)

func main() {
	texts := "G@ M@ney~"
	fmt.Println(texts)
	r := strings.NewReplacer("@", "o")
	newTexts := r.Replace(texts)
	fmt.Println(newTexts)

	// //zero value
	// var e string
	// var d bool
	// var c rune
	// var b float64
	// var a int

	// //naming convention
	// var studentId string
	// var i int32
	// fmt.Println(studentId)
	// fmt.Println(i)

	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)
	// fmt.Println(d)
	// fmt.Println(e)

	// fmt.Printf("%T\n", d)
	// fmt.Printf("%T\n", e)

	// fmt.Println(reflect.TypeOf(d))
	// fmt.Println(reflect.TypeOf(e))
}
