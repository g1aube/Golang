package main

import "fmt"

var first_name string
var last_name string
var age int

func main() {
	const name string = "Ivan" // объявляем константу (ее изменить нельзя)
	const surname string = "Kadetov"

	fmt.Printf("%s %s", name, surname)
	fmt.Printf("%T %T", name, surname)

	fmt.Println("Enter your first name, please: ")
	fmt.Scanln(&first_name)

	fmt.Println("Enter your last name, please: ")
	fmt.Scanln(&last_name)

	fmt.Println("Enter your age, please: ")
	fmt.Scanln(&age)

	fmt.Print("Your info is:")
	fmt.Print(first_name+" "+last_name, "\nAge: ", age)

}
