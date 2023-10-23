package main

//import 여러개를 묶어서 쓸 수 있다
import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var c rune = '가' //rune: 작은 따옴표(var 변수이름 변수타입)
	//var a int16 = 7  //2byte
	//var a = 7
	a := 7
	var b float64 = 5.3
	a = int(b) //Type Conversion, Casting

	d := false
	fmt.Println(d)        //결과: false
	fmt.Printf("%T\n", d) //결과: bool

	fmt.Println(a)
	fmt.Printf("%T\n", a)

	fmt.Println(c)        //유니코드 출력(UTF-8)
	fmt.Printf("%c\n", c) //한글자 출력
	fmt.Printf("%T\n", c) //rune은 int32의 별명(4byte)

	fmt.Println(math.Ceil(2.71)) //올림 함수
	//Floor: 내림 //Round: 반올림
	fmt.Println("Hello Go~")
	fmt.Println(strings.Title("go git github java"))
}
