package main

import (
	"fmt"
)

func status(name string) {
	balls := map[string]int{"성기훈": 20, "오일남": 0}
	ball := balls[name]
	fmt.Println(ball)
	ball, exists := balls[name]
	if !exists {
		fmt.Println(name, "남은 게임 참여자가 아닙니다.")
	} else if ball < 1 {
		fmt.Println(name, "님은", balls[name], "개로 탈락! 탕탕")
	} else {
		fmt.Println(name, "남은 게임에서 승리하였습니다.")
	}
}

func main() {
	status("오일남")
	status("강철")
	status("남주")
	status("여주")

}

// func main() {
// 	var balls map[string]int
// 	//balls := make(map[string]int)
// 	fmt.Printf("%#v\n", balls)
// 	balls["성기훈"] = 20
// 	fmt.Println(balls["성기훈"])
// 	fmt.Println(balls["오일남"])
// }

// func main() {
// 	games := map[int]string{
// 		456: "성기훈",
// 		218: "박해수",
// 		067: "강새벽",
// 		001: "오일남",
// 		199: "알리",
// 		101: "아이오아이",
// 	}

// 	name, ok := games[100]
// 	fmt.Println(ok, name)

// 	for _, v := range games {
// 		fmt.Println(v)
// 	}
// 	games[101] = "장덕수" //update

// 	for k, v := range games {
// 		fmt.Println(k, v)
// 	}
// }
