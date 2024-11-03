package main

//in golang always need a library if you need print anything on screen
import "fmt" //format library to print on screen

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

	fmt.Println(name) // you can print the variable with the word fmt and the word Println and the name of the variable
}
