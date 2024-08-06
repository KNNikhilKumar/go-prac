package firstclass

import "fmt"

//func taking input as param

func add(x int, y int) int {
	return x + y
}

func mul(x int, y int) int {
	return x * y
}

func perform(x int, y int, action func(int, int) int) int {
	return action(x, action(x, y))
}

//func returning func Curried Function

func selfConstruct(action func(int, int) int) func(int) int {
	return func(x int) int {
		return action(x, x)
	}
}

func ExecFC() {
	fmt.Println(perform(1, 2, add))
	fmt.Println(perform(1, 2, mul))

	//we use mul to calculate square
	squareFunc := selfConstruct(mul)
	fmt.Println(squareFunc(3))
}
