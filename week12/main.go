package main

import "fmt"

func main() {
	// var s []int
	// s = make([]int, 5)

	//s := make([]int, 5) //단축 연산자

	s := []int{0, 0, 0, 0, 0} //슬라이스 리터럴 이용하여 슬라이스 생성 및 메모리 할당, 초기화 진행

	for _, value := range s {
		fmt.Println(value)
	}

	copyS := s[1:4]
	for _, value := range copyS {
		fmt.Println(value)
	}

}
