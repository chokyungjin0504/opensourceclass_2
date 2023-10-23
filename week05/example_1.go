package main

import (
	"bufio"
	"fmt"
	"log" // 랜덤 숫자 생성을 위한 패키지 추가
	"os"
	"strconv"
	"strings"
	// 랜덤 숫자 생성 시드를 설정하기 위한 패키지 추가
)

func main() {
	fmt.Print("Input score: ")
	reader := bufio.NewReader(os.Stdin)
	inputScoreString, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	inputScoreString = strings.TrimSpace(inputScoreString)      //공백제거
	inputScore, err := strconv.ParseFloat(inputScoreString, 64) //inputScoreString을 숫자로 변경

	var grade string // grade 변수 선언

	if inputScore >= 90 {
		grade = "A grade" // 변수 할당 수정
	} else {
		grade = "BCDF grade" // 변수 할당 수정
	}
	fmt.Println("You got " + grade)

}
