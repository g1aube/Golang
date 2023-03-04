package main

import "fmt"

func main() {
	var array [3]string // массив, длина 3, состоит из строк
	fmt.Println(len(array))
	arr := make([]int, 10)
	fmt.Println(arr)

	for _, x := range arr {
		fmt.Println(x)
	}
	array[1] = "be"
	fmt.Println(array)
	arr_int := []int{1, 2, 3, 4}
	fmt.Println(arr_int)
}
