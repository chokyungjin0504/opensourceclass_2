package main

import "fmt"

func main() {
	// var primes [3]int
	// primes[0] = 2
	// primes[1] = 3
	// primes[2] = 5
	// fmt.Println(primes, primes[1])

	// var primes [3]int = [3]int{2, 3, 5}
	// fmt.Println(primes, primes[1])

	primes := [3]int{2, 3, 5}
	fmt.Println(primes, primes[1])

	test := [5]bool{true, true, true}
	fmt.Println(test[3]) //boolean 타입의 제로값, false

	fmt.Println(primes[4]) // invalid argurment: index 4 out of bounds [1:3]

	i := 0
	//for i < 4 { //panic: runtime error: index out of range [3] with length
	for i < len(primes) {
		fmt.Println(primes[i])
		i++
	}

	for _, prime : = range primes {
		fmt.Println(prime)
	}

	fmt.Printf("%#v\n", test)
	fmt.Println(test)
	fmt.Printf("%#v\n", primes)
	fmt.Println(primes)

	// for idx, primes := range primes { //값만 출력하려 했으나 인덱스가 출력됨
	// 	fmt.Println(idx, primes)
	// }

	// //for j := 0; j<len(primes); i++ {
	// 	fmt.Prinltn(primes[j])
	// }
}
