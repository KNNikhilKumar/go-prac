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

func ExecFC() {
	fmt.Println(perform(1, 2, add))
	fmt.Println(perform(1, 2, mul))
}
