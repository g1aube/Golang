package main

import "fmt" // импорт библиотеки, позволяющей писать в терминал

// Syntax
// var имя_переменной тип_данных
// string, int, bool, float64/32

var x int
var u string
var y int32
var f float64

func main() {
	var hello_message = "Hello!!!"
	var random_int = 2
	fmt.Println("Hello World!") // пишет в терминал Hello World!
	fmt.Println(hello_message)
	fmt.Println((random_int + 2))
}
