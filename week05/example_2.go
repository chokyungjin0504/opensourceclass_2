package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	HotSpur := "hm ? j madi?"
	replacePlayer := strings.NewReplacer("?", "son")
	player := replacePlayer.Replace(HotSpur)
	fmt.Println(player)

	rand.Seed(time.Now().UnixNano()) // 랜덤 숫자 생성 시드 설정
	dice := rand.Intn(6) + 1
	fmt.Println(dice)
}
