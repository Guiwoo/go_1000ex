package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"unicode/utf8"
)

func ex02() {
	var file string
	_, file = path.Split("css/main.css")
	fmt.Println(file)
}

// When use a short declaration ?
// if you don't know initial value ? do not use it
// when you need a package scoped variable

var globalValue int

func ex03() {
	// score := 0 // DONT
	var score int // 👍Already score = 0

	// better readability grouping the value like  this
	var (
		video    string
		duration int
		current  int
	)
	fmt.Println(score, video, duration, current)
}

// But Go Developer love short init way
func ex04() {
	// If you know the specific value ? use this
	width, height := 100, 160 //👍 Good

	width, color := 50, "red" // 👍Look better

	fmt.Println(width, height, color)
}

// type conversion
func ex05() {
	// You can't use value belong to different types together
	speed := 100 // int
	force := 2.5 // float64

	// You need to explicitly convert the values
	// conversion can change value it self
	speed = speed * int(force)

	speed = int(float64(speed) * force)
}

func ex06() {
	// A slice can store multiple values
	var Args []string // Args's type is slice of strings
	// Args[0] => path to the program Path to the program
	Args = os.Args
	fmt.Printf("%#v\n", Args)
	fmt.Println("Paht", Args[0])
	fmt.Println("1st", Args[1])
	fmt.Println("Number of items", len(Args))
}

func s01() {
	// Both are string type
	// string literal => single line interpreted
	// raw string literal => multi line not interpreted

	var s string
	s = "How are you ?"
	s = `How are you?`

	s = "<html>\n\t<body>\"Hello\"</body>\n</html>"
	s = `
<html>
	<body>"Hello"</body>
</html>`
	fmt.Println(s)
	fmt.Println("c:\\my\\dir\\file")
	fmt.Println(`c:\my\dir\file`)
}

func s02() {
	name := "ㄱㄴㄷ"
	fmt.Println(utf8.RuneCountInString(name))
	arr := []rune(name)
	fmt.Println(len(arr))
}

func s03() {
	msg := os.Args[1]
	s := strings.ToUpper(msg) + strings.Repeat("!", len(msg))
	fmt.Println(s)
}

func iota01() {
	// iota is a built on constant generate which generates ever increasing number
	// Expressions with iota, So the other constants will repeat the expressions
	const (
		monday = iota + 1
		tuesday
		wednesday
		thursday
		friday
		saturday
		sunday
	)

	fmt.Println(monday, tuesday, wednesday, thursday, friday, saturday, sunday)
}

func blankIdentifier02() {
	// which formula do i need to inialize constants with correct timezone value ?
	const (
		EST = -(5 + iota)
		_
		MST
		PST
	)
}

func main() {
	iota01()
}
