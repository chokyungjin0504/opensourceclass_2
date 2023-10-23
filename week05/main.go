package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

/*
func main() {

	var now time.Time
	now = time.Now()
	year := now.Year()
	var month string = now.Month().String()
	fmt.Println(year, month, now.Hour(), now.Minute(), now.Second())


	HotSpur := "hm ? j madi?"
	replacePlayer := strings.NewReplayer("?","son")
	player := replacePlayer.Replace(HotSpur)
	fmt.Println(player)
*/
/*
	fmt.Print("Input score:")
	reader := bufio.newReader(os.Stdin)
	InputScore := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	inputScoreString = strings.TrimSpace(inputScoreString)
	inputScore, err := strconv.ParseFloat(inputScoreString, 32)

	if inputSccore >= 90 {
		grade := "A grade"
	} else {
		grade := "BCDF grade"
	}
	fmt.Println("You got" +grade)


	dice := reand.Intn(6) + 1
	fmt.Println(dice)
*/

func main() {
	rand.Seed(time.Now().Unix()) //get the current data and time as an integer
	answer := rand.Intn(100) + 1 // random integer number (1~100)
	fmt.Println("Guess Number Game~!")
	fmt.Println(answer)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 10; i++ {
		fmt.Println("You have ", 10-i, "chances~")
		fmt.Print("Input guess number")
		inputNumberString, err := reader.ReadString('\n')
		if err != nil {

			log.Fatal(err)
		}
		inputNumberString = strings.TrimSpace(inputNumberString)
		inputNumber, err := strconv.Atoi(inputNumberString) //문자열을 정수로 변환
		if err != nil {
			log.Fatal(err)
		}
		if inputNumber == answer {
			fmt.Println("Great You got the Number. Congrates~") // You got the answer!
		} else if inputNumber < answer {
			fmt.Println("Guess number is lower than answer") // The answer is higher!
		} else if inputNumber > answer {
			fmt.Println("Guess number is higher than answer") // The answer is lower!
		}
	}
}
