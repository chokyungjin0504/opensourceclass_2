// average calculates the average of several numbers.
package main

import (
	"fmt"
	"log"

	"github.com/headfirstgo/datafile"
)

func main() {
	numbers, err := datafile.GetFloats("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	sampleCount := float64(len(numbers))
	fmt.Printf("Average: %0.2f\n", sum/sampleCount)
}

// package main

// import "fmt"

// func main() {
// 	var a []string
// 	var b []bool
// 	// a = make([]string, 4, 5)

// 	b = append(b, true)
// 	fmt.Printf("%#v %#v\n", a, b)
// 	fmt.Println(a, len(a), cap(a))
// 	fmt.Println(b, len(b), cap(b))
// }
