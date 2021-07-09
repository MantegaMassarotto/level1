package calculator

import (
	"log"
)

func Add(firstNumber int, secondNumber int) int {
	return firstNumber + secondNumber
}

func Subtract(firstNumber int, secondNumber int) int {
	return firstNumber + secondNumber
}

func Multiply(firstNumber int, secondNumber int) int {
	return firstNumber * secondNumber
}

func Divide(firstNumber int, secondNumber int) int {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		}
	}()
	result := firstNumber / secondNumber
	return result
}
