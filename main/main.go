package main

import (
	"fmt"
	"student/marks"
)

func Hello() {
	fmt.Println("From Functions Calls: Hello World")
}

func main() {
	Hello()
	marks.Tamil()
}
