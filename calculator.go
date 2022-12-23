package main

import (
	"fmt"

	"github.com/Venkatesh-8203/Calculator/Add"

	AddOperation "github.com/Venkatesh-8203/Calculator/Add"
	DivOperation "github.com/Venkatesh-8203/Calculator/Division"
	MulOperation "github.com/Venkatesh-8203/Calculator/Multiply"
	SubOperation "github.com/Venkatesh-8203/Calculator/Subtract"
)

func main() {
	a := 100
	b := 50

	addResult := AddOperation.Add(a, b)
	subResult := SubOperation.Subtract(a, b)
	mulResult := MulOperation.Multiply(a, b)
	divResult := DivOperation.Divide(a, b)

	fmt.Println("Add: ", addResult)
	fmt.Println("Sub: ", subResult)
	fmt.Println("Mul: ", mulResult)
	fmt.Println("Div: ", divResult)

	addResult1 := Add.Add(a, b)
	fmt.Println("addResult1: ", addResult1)
}
