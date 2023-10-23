package main

import (
	"fmt"
	"strings"
)

func main() {
	HotSpur := "hm ? j madi?"
	replacePlayer := strings.NewReplacer("?", "son")
	player := replacePlayer.Replace(HotSpur)
	fmt.Println(player)
}
