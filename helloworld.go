package main

import (
	"fmt"
)

const GO_NAME string = "Go"

func demo001Type() {
	fmt.Println("\ndemo001Type")

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

func demo002Array() {
	fmt.Println("\ndemo002Array")
	var fixSizeGrades = [5]float32{8, 7, 9, 6.5, 3}

	fmt.Printf("fixSizeGrades Type: %T, value: %v\n", fixSizeGrades, fixSizeGrades)
	fmt.Printf("fixSizeGrades[1] Type: %T, value: %v\n", fixSizeGrades[1], fixSizeGrades[1])

	autoSizeGrades := [...]float32{1.5, 9.3, 8, 7, 9, 6.5, 3}
	fmt.Printf("autoSizeGrades Type: %T, value: %v\n", autoSizeGrades, autoSizeGrades)
	var defaultZeroGrades = [10]float32{8, 9.1}
	var defaultInitWithIndexGrades = [10]float32{0: 8, 5: 9.1}

	fmt.Printf("defaultZeroGrades Type: %T, value: %v\n", defaultZeroGrades, defaultZeroGrades)
	fmt.Printf("defaultInitWithIndexGrades Type: %T, value: %v\n", defaultInitWithIndexGrades, defaultInitWithIndexGrades)
}

//This is main function, will start first when your program start
func main() {
	demo001Type()
	demo002Array()
}

/*
This is multi-line comments
It'll help write a lot of comments here
*/
