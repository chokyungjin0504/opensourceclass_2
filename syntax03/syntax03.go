package main

import (
	"bufio"
	"fmt"
	"log" //Fatal 함수 이용: 에러 확인 후 프로그램종료
	"os"
	"strconv" //TrimSpace
	"strings" //ParseInt
)

func main() {
	fmt.Print("단 입력: ")
	rd := bufio.NewReader(os.Stdin)
	in, err := rd.ReadString('\n')

	if err != nil { //에러가 발생
		log.Fatal(err)
	}
	in = strings.TrimSpace(in)
	dan, err := strconv.Atoi(in)
	//dan, err := strconv.ParseInt(in, 10, 32) //10진수, int32 8비트
	if err != nil {
		log.Fatal(err)
	}
	// for i := 1; i < 10; i++ {
	// 	fmt.Printf("%d * %d = %d\n", dan, i, (dan * i))
	// 	//fmt.Println(dan, " * ", i, " = ", (dan * i)) //int32와 int64로 맞지 않아 dan전체를 int로 형변환
	// }
	//다른 언어의 while문 구현
	i := 1
	for i < 10 {
		fmt.Printf("%d * %d = %d\n", dan, i, (dan * i))
		i++
	}

	//fmt.Println(dan * 2)

	// fmt.Print("숫자입력 : ")            //Print와 Println은 줄바꿈 차이
	// rd := bufio.NewReader(os.Stdin) //standard input: 키보드로 입력받은 값
	// in, _ := rd.ReadString('\n')    // enterkey까지 입력받는다
	// fmt.Println(in)

}
