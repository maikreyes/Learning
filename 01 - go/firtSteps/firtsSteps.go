package main

//in golang always need a library if you need print anything on screen
import (
	"fmt"     //format library to print on screen
	"reflect" //reflect library to print the type of the variable
)

//Firts Steps  in golang Languaje

func main() {

	fmt.Println("Hello World") // you don´t need semicolon when you finish a comand line  in thos case "fmt" it´s a reservb word for format text

	//Languaje Fundaments in golang

	/*
		In golang always need when you create a varible need to specify the type of data

		data types in golang
		- int
		- float
		- string
		- bool
		- complex
		- byte
		- rune
		- uint
		- uintptr
		- array
		- slice
		- map
		- struct
		- interface
		- pointer
		- function
		- channel
		- error

		you can create a variable in golang with the next sintaxys
	*/

	var name string = "Jhon" // you can create a variable with the word var and the name of the variable and the type of data and the value of the variable

	// you can declarate first the variable and then assign the value of the variable
	var lastName string
	lastName = "Doe"

	fmt.Println(name, lastName) // you can print the variable with the word fmt and the word Println and the name of the variable

	//also you can change the value of the variable

	name = "Jhonatan"

	fmt.Println(reflect.TypeOf(name)) // you can print the type of the variable with the word reflect and the word TypeOf and the name of the variable

	fmt.Println(name) // you can print the variable with the word fmt and the word Println and the name of the variable

	//in go you don´t have class you have structs

	type MyStruct struct {
		name string
		age  int
	}

	mystruct := MyStruct{name: "Jhon", age: 25}

	fmt.Println(mystruct)

}
