package main

import (
	"fmt"
	"matchLib/pkg/calculations"
)

func main() {
	n := int64(5) // пример входного числа
	result := calculations.Calculate(n, true)
	fmt.Printf("%d! = %d\n", n, result)
}
