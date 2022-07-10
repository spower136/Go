package main

import (
	"fmt"
	"calculator"
)

func main() {
	// problem #1
	fmt.Println(calculator.Calc("rectangle", "area", 20, 30))
	fmt.Println(calculator.Calc("rectangle", "perimeter", 20, 30))
	fmt.Println(calculator.Calc("square", "area", 20, 20))
	fmt.Println(calculator.Calc("square", "perimeter", 20, 20))

	// problem #2
	maxNum := 10
	threeMultSum, fiveMultSum := calculator.Sums(maxNum)
	fmt.Println("Enter the upper bound:", maxNum)
	fmt.Println("The sum of numbers less than", maxNum, "divisible by 3 is", threeMultSum)
	fmt.Println("The sum of numbers less than", maxNum, "divisible by 5 is", fiveMultSum)

	// problem #3
	calculator.Transform("aBcD")
	calculator.Transform("αΒγΔ")
	calculator.Transform("h￿ello%3 !world")
	calculator.Transform("a!% b£c $€")
}
