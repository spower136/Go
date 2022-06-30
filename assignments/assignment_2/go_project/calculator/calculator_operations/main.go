package main

import (
	"calculator"
	"fmt"
)

func main() {
	step1 := calculator.Multiply(3, 2)
	step2 := calculator.Add(step1, 6)
	step3 := calculator.Multiply(3, 4)
	step4 := calculator.Add(step3, 2)
	result := calculator.Divide(step2, step4)
	fmt.Println(result)
}
