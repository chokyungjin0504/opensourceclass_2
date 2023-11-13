package main

import "fmt"

func main() {
	a := make([]string, 4, 5) // 타입, 개수, 수용량
	a[0] = "a"
	a[2] = "C"
	a[3] = "d"
	as := a[0:2]
	as[1] = "z"

	fmt.Println(a, len(a), cap(a)) //cap: 내장함수
}

// func main() {
// 	a := []string{"a", "b", "c", "d"}
// 	as := a[:2]
// 	as[1] = "z"
// 	fmt.Println(a)
// 	fmt.Println(as)

// 	b := [4]int{4, 3, 2, 1}
// 	bs := b[1:3]
// 	fmt.Println(bs)
// }

// var s []int
// s = make([]int, 5)

//s := make([]int, 5) //단축 연산자

// s := []int{0, 0, 91, 0, 99} //슬라이스 리터럴 이용하여 슬라이스 생성 및 메모리 할당, 초기화 진행

// for _, value := range s {
// 	fmt.Println(value)
// }

// copyS := s[1:4]
// for _, value := range copyS {
// 	fmt.Println(value)
// }

// test := [3]string{"inha", "go", "student"} //배열 리터럴을 이용해서 test 배열 생성
// //testS := test[0:4]                         //invalid argument: index 4 out of bounds [0:4]
// testS := test[:2]
// testS2 := test[1:]
// testS2[0] = "python"
// fmt.Println(testS2[1])
// fmt.Println(testS, len(testS))
// fmt.Println(test)
