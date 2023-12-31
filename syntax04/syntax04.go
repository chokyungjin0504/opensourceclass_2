package main

import (
	"fmt"
	"math/rand"
	"time" //Seed 생성용 패키지
)

// 난수 추출된 수의 소수 판정 프로그램
func main() {
	// seed 설정
	seed := time.Now().Unix()
	rand.Seed(seed)

	//가독성을 위해 count를 isPrime으로 변경
	isPrime := true
	number := rand.Intn(150) + 2 //0과 1제외, 2~151 사이의 수
	fmt.Println("임의로 추출된 수 :", number)

	for i := 2; i < number; i++ {
		if number%i == 0 {
			isPrime = false
			break //첫번째 약수가 발견되면 반복문 강제 종료
			//count++
			//count = count + 1
		}
		//fmt.Print(i, " ")
	}
	if isPrime { //default가 true여서 줄이기 가능
		fmt.Println(number, "은 소수")
	} else {
		fmt.Println(number, "은 소수아님")
	}
}

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time" //Seed 생성용 패키지
// )

// // 난수 추출된 수의 소수 판정 프로그램
// func main() {
// 	// seed 설정
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	//가독성을 위해 count를 isPrime으로 변경
// 	isPrime := true
// 	number := rand.Intn(150) + 2 //0과 1제외, 2~151 사이의 수
// 	fmt.Println("임의로 추출된 수 :", number)

// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			//count++
// 			//count = count + 1
// 		}
// 		fmt.Print(i, " ")
// 	}
// 	if isPrime { //default가 true여서 줄이기 가능
// 		fmt.Println(number, "은 소수")
// 	} else {
// 		fmt.Println(number, "은 소수아님")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time" //Seed 생성용 패키지
// )

// // 난수 추출된 수의 소수 판정 프로그램
// func main() {
// 	// seed 설정
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	//가독성을 위해 count를 isPrime으로 변경
// 	isPrime := true
// 	number := rand.Intn(150) + 2 //0과 1제외, 2~151 사이의 수
// 	fmt.Println("임의로 추출된 수 :", number)

// 	for i := 2; i < number; i++ {
// 		if number%i == 0 {
// 			isPrime = false
// 			//count++
// 			//count = count + 1
// 		}
// 	}
// 	if isPrime == true {
// 		fmt.Println(number, "은 소수")
// 	} else {
// 		fmt.Println(number, "은 소수아님")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time" //Seed 생성용 패키지
// )

// // 난수 추출된 수의 소수 판정 프로그램
// func main() {
// 	// seed 설정
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	count := 0
// 	number := rand.Intn(150) + 2 //0과 1제외, 2~151 사이의 수
// 	fmt.Println("임의로 추출된 수 :", number)

// 	for i := 1; i <= number; i++ {
// 		if number%i == 0 {
// 			count++
// 		}
// 	}
// 	if count == 2 {
// 		fmt.Println(number, "는 소수입니다!")
// 	} else {
// 		fmt.Println(number, "는 소수가 아닙니다.")
// 	}
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// 	"time" //Seed 생성용 패키지
// )

// func main() {
// 	// seed 설정
// 	seed := time.Now().Unix()
// 	rand.Seed(seed)

// 	for i := 1; i < 6; i++ {
// 		dice := rand.Intn(6) + 1
// 		fmt.Println(dice)
// 	}
// }
