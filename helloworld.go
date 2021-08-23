package main

import (
	"fmt"
)

const GO_NAME string = "Go"

//This is main function, will start first when your program start
func main() {
	fmt.Println("Hello Duong!")

	var name string = "Duong"
	var isSingle bool = false
	var isHaveChild = true
	var gradeAvg float32 = 8.00001
	var age int64 = 31
	var numOfChild uint
	numOfWife := 1

	numOfChild = 2

	fmt.Printf("name Type: %T, value: %v\n", name, name)
	fmt.Printf("isSingle Type: %T, value: %v\n", isSingle, isSingle)
	fmt.Printf("isHaveChild Type: %T, value: %v\n", isHaveChild, isHaveChild)
	fmt.Printf("gradeAvg Type: %T, value: %v\n", gradeAvg, gradeAvg)
	fmt.Printf("age Type: %T, value: %v\n", age, age)
	fmt.Printf("numOfChild Type: %T, value: %v\n", numOfChild, numOfChild)
	fmt.Printf("numOfWife Type: %T, value: %v\n", numOfWife, numOfWife)
}

/*
This is multi-line comments
It'll help write a lot of comments here
*/
