package main

//shadowing
func main() {

}

// package main

// import (
// 	"fmt"
// 	"log"
// )

// // 입력된 수의 소수 판정 프로그램
// func main() {
// 	var number int
// 	fmt.Print("정수입력: ")
// 	_, err := fmt.Scanln(&number)

// 	//fmt.Println(n) //갯수 몇개인지 안나옴

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	isPrime := true
// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			break
// 		}
// 	}
// 	if isPrime {
// 		fmt.Println(number, "은 소수")
// 	} else {
// 		fmt.Println(number, "은 소수아님")
// 	}
// }

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// // 입력된 수의 소수 판정 프로그램
// func main() {

// 	fmt.Print("정수 입력: ")
// 	rd := bufio.NewReader(os.Stdin)
// 	in, err := rd.ReadString('\n')

// 	if err != nil { //에러가 발생
// 		log.Fatal(err)
// 	}
// 	in = strings.TrimSpace(in)
// 	number, err := strconv.Atoi(in)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	isPrime := true //초기화

// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			break
// 		}
// 	}
// 	if isPrime {
// 		fmt.Println(number, "은 소수")
// 	} else {
// 		fmt.Println(number, "은 소수아님")
// 	}
// }
