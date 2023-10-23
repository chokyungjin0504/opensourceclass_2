//두 수 사이의 소수 모두 출력

package main

import (
	"fmt"
	"log"
	"os"
)

func isPrime(n int) (bool, error) {

	for n < 2 {
		return false, fmt.Errorf("%d 는 소수가 아님", n)
	}

	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

func prime(number int) {
	p, err := isPrime(number)
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
	if p {
		fmt.Println(number, "은 소수")
	} else {
		fmt.Println(number, "은 소수 아님")
	}
}

func primeRange(a int, b int) {
	if a > b {
		temp := a
		a = b
		b = temp
	}

	for i := a; i <= b; i++ {
		p, err := isPrime(i)
		if err != nil {
			fmt.Println(err)
			os.Exit(0)
		}
		if p {
			fmt.Print(i, " ")
		}
	}
}

func main() {
	var menu int

	for true { //무한 루프
		fmt.Print("1) 소수판정 2) 구간 소수판정")
		_, err := fmt.Scanln(&menu)

		if err != nil {
			log.Fatal(err)
		}

		switch menu {
		case 1:
			var in int
			fmt.Print("정수 1개 입력: ")
			_, err := fmt.Scanln(&in)

			if err != nil {
				log.Fatal(err)
			}
			prime(in)
		case 2:
			var n1, n2 int
			fmt.Print("정수 2개 입력: ")
			_, err := fmt.Scanln(&n1, &n2)

			if err != nil {
				log.Fatal(err)
			}
			primeRange(n1, n2)
		default:
			fmt.Print("프로그램을 종료합니다")
			os.Exit(1)
		}
	}
}

// //두 수 사이의 소수 모두 출력

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func isPrime(n int) (bool, error) {

// 	for n < 2 {
// 		return false, fmt.Errorf("%d 는 소수가 아님", n)
// 	}

// 	for i := 2; i < n; i++ {
// 		if n%i == 0 {
// 			return false, nil
// 		}
// 	}
// 	return true, nil
// }

// func main() {
// 	var a, b int

// 	fmt.Print("정수입력:")
// 	_, err := fmt.Scanln(&a, &b)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if a > b {
// 		temp := a
// 		a = b
// 		b = temp
// 	}

// 	for i := a; i <= b; i++ {
// 		p, err := isPrime(i)
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(0)
// 		}
// 		if p {
// 			fmt.Print(i, " ")
// 		}
// 	}
// }

// //두 수 사이의 소수 모두 출력

// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func isPrime(n int) (bool, error) {
// 	prime := true

// 	for n < 2 {
// 		return false, fmt.Errorf("%d 는 소수가 아님", n)
// 	}

// 	for i := 2; i < n; i++ {
// 		if n%i == 0 {
// 			prime = false
// 			break
// 		}
// 	}
// 	return prime, nil
// }

// func main() {
// 	var a, b int

// 	fmt.Print("정수입력:")
// 	_, err := fmt.Scanln(&a, &b)

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if a > b {
// 		temp := a
// 		a = b
// 		b = temp
// 	}

// 	for i := a; i <= b; i++ {
// 		p, err := isPrime(i)
// 		if err != nil {
// 			fmt.Println(err)
// 			os.Exit(0)
// 		}
// 		if p {
// 			fmt.Print(i, " ")
// 		}
// 	}
// }

// // 소수 판정 프로그램: 함수 적용, err 리턴
// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func isPrime(n int) (bool, error) {
// 	prime := true

// 	for n < 2 {
// 		return false, fmt.Errorf("%d 는 소수가 아님", n)
// 	}

// 	for i := 2; i < n; i++ {
// 		if n%1 == 0 {
// 			prime = false
// 			break
// 		}
// 	}

// 	return prime, nil //소수 판정 값, 정상 데이터
// }

// func main() {
// 	var number int

// 	fmt.Print("정수입력:")
// 	_, err := fmt.Scanln(&number)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	p, err := isPrime(number)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(o)
// 	}

// 	if p {
// 		fmt.Println(number, "은 소수")
// 	} else {
// 		fmt.Println(number, "은 소수 아님")
// 	}
// }

// // 소수 판정 프로그램: 함수 적용
// package main

// import (
// 	"fmt"
// 	"log"
// 	"os"
// )

// func isPrime(n int) (bool) {
// 	prime := true
// 	for i := 2; i < n; i++ {
// 		if n%1 == 0 {
// 			prime = false
// 			break
// 		}
// 	}

// 	return prime
// }

// func main() {
// 	var number int

// 	fmt.Print("정수입력:")
// 	_, err := fmt.Scanln(&number)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	for number < 2 {
// 		fmt.Println(number, "은 소수가 아님")
// 		os.Exit(0)
// 	}

// 	if isPrime(number) {
// 		fmt.Println(number, "은 소수")
// 	} else {
// 		fmt.Println(number, "은 소수 아님")
// 	}
// }

// // after(multi return)
// package main

// import "fmt"

// func processScore(kor int, eng int, math int) (int, int) {
// 	totalScore := kor + eng + math
// 	average := (kor + eng + math) / 3.0

// 	return totalScore, average

// 	//fmt.Printf("%s의 총점은 %d 평균은 %d\n", name, totalScore, average)
// }

// func main() {
// 	var t int
// 	var a int
// 	t, a = processScore(100, 90, 93)
// 	fmt.Printf("%s의 총점은 %d 평균은 %d\n", "조경진", t, a)
// 	t, a = processScore(90, 30, 43)
// 	fmt.Printf("%s의 총점은 %d 평균은 %d\n", "조경민", t, a)
// }

// // after
// package main

// import "fmt"

// func processScore(name string, kor int, eng int, math int) {
// 	totalScore := kor + eng + math
// 	average := (kor + eng + math) / 3.0

// 	fmt.Printf("%s의 총점은 %d 평균은 %d\n", name, totalScore, average)
// }

// func main() {
// 	processScore("조경진", 100, 90, 93)
// 	processScore("조경민", 90, 30, 43)
// }

// before
// package main

// import "fmt"

// func main() {
// 	kor := 100
// 	eng := 90
// 	math := 93
// 	name := "조경진"

// 	fmt.Println(name, "의 총점은", (kor + eng + math), "평균은", (kor+eng+math)/3.0)

// 	kor = 99
// 	eng = 91
// 	math = 92
// 	name = "조경민"

// 	fmt.Println(name, "의 총점은", (kor + eng + math), "평균은", (kor+eng+math)/3.0)
// }

// package main

// import "fmt"

// func sayHello() {
// 	fmt.Println("Hello")
// }

// func main() {
// 	sayHello()
// }
