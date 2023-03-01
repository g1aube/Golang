package main

import (
	"fmt"
)

func main() {
	my_list := []string{"hello", "world", "!"} // создание списка
	fmt.Print(my_list[2], my_list[1])

	for _, i := range my_list {
		fmt.Print(i)
	}

	for x := 2; x <= 2; x++ {
		fmt.Print(x)
	}

	var h int = 5

	for h <= 10 {
		fmt.Print(h)
		h = h + 1
	}
}
