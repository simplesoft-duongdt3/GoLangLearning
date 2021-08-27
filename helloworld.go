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

func demo003Slice() {
	fmt.Println("\ndemo003Slice")
	var sliceEmpty1 = []string{}
	fmt.Printf("sliceEmpty1 Type: %T, value: %v\n", sliceEmpty1, sliceEmpty1)

	sliceEmpty2 := []string{}
	fmt.Printf("sliceEmpty2 Type: %T, value: %v\n", sliceEmpty2, sliceEmpty2)

	sliceFull := []int{1, 2, 3}
	fmt.Printf("sliceFull Type: %T, value: %v len=%v cap=%v\n", sliceFull, sliceFull, len(sliceFull), cap(sliceFull))

	var childArray = [4]string{"Chip", "Khia", "Cua", "Cong"}
	var sliceChild1 = childArray[0:2]
	fmt.Printf("sliceChild1 Type: %T, value: %v len=%v cap=%v\n", sliceChild1, sliceChild1, len(sliceChild1), cap(sliceChild1))

	//cap depen on lengh of array (from start to end)
	var sliceChild2 = childArray[1:3]
	fmt.Printf("sliceChild2 Type: %T, value: %v len=%v cap=%v\n", sliceChild2, sliceChild2, len(sliceChild2), cap(sliceChild2))

	sliceAnimal1 := make([]int, 5, 10)
	fmt.Printf("sliceAnimal1 Type: %T, value: %v len=%v cap=%v\n", sliceAnimal1, sliceAnimal1, len(sliceAnimal1), cap(sliceAnimal1))

	sliceAnimal2 := make([]string, 5)
	fmt.Printf("sliceAnimal2 Type: %T, value: %v len=%v cap=%v\n", sliceAnimal2, sliceAnimal2, len(sliceAnimal2), cap(sliceAnimal2))

	var child1 = sliceChild1[0]
	fmt.Printf("child1 Type: %T, value: %v\n", child1, child1)

	var animal1 = sliceAnimal2[0]
	fmt.Printf("animal1 Type: %T, value: %v len=%v\n", animal1, animal1, len(animal1))

	sliceAnimal2 = append(sliceAnimal2, "Heo", "ga")
	fmt.Printf("sliceAnimal2 Type: %T, value: %v len=%v cap=%v\n", sliceAnimal2, sliceAnimal2, len(sliceAnimal2), cap(sliceAnimal2))
}

func demo004If() {
	fmt.Println("\ndemo004ControlFlow")
	var englishGrade float32 = 7.8
	var mathGrade float32 = 6.5
	if englishGrade > mathGrade {
		fmt.Printf("You learn English better Math englishGrade: %v, mathGrade: %v\n", englishGrade, mathGrade)
	} else if mathGrade > englishGrade {
		fmt.Printf("You learn Math better English mathGrade: %v, englishGrade: %v\n", mathGrade, englishGrade)
	} else {
		fmt.Printf("You learn Math equal English mathGrade: %v, englishGrade: %v\n", mathGrade, englishGrade)
	}
}

func demo005Switch() {
	fmt.Println("\ndemo005Switch")

	day := 4

	fmt.Printf("day: %v\n", day)
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Not a weekday")
	}

	switch day {
	case 1, 3, 5, 7, 9:
		fmt.Println("Odd")
	case 2, 4, 6, 8:
		fmt.Println("Even")
	}
}

func demo006For() {
	fmt.Println("\ndemo006For")
	for i := 0; i <= 100; i++ {
		if i == 88 {
			break
		}
		if i%3 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

//This is main function, will start first when your program start
func main() {
	demo001Type()
	demo002Array()
	demo003Slice()
	demo004If()
	demo005Switch()
	demo006For()
}

/*
This is multi-line comments
It'll help write a lot of comments here
*/
