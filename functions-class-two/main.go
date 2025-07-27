package main

import (
	"fmt"
)

type intTransformation func(int) int

// type globalTransformation func[T int | float64](T) T  esto no se puede hacer, me da error

func main() {
	collection := []int{1, 2, 3, 4}
	collection2 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(collection)
	result := transformNumbers(&collection, func(x int) int { return x * 4 })
	fmt.Println(result)
	result = transformNumbers(&collection, getOperation())
	fmt.Println(result)

	fmt.Println(collection2)
	result2 := transformAllNumbers(&collection2, doubleFloat)
	fmt.Println(result2)

	byTen := createTransformation(10)
	result3 := transformAllNumbers(&collection, byTen)
	fmt.Println(result3)

	fmt.Println(sumup(10, 20, 30, 40))
	fmt.Println(sumup(result3...))

}

// func transformNumbers(numbers *[]int, transformation func(int) int) []int {
func transformNumbers(numbers *[]int, transformation intTransformation) []int {
	newNumbers := make([]int, len(*numbers))
	for index, item := range *numbers {
		newNumbers[index] = transformation(item)
	}

	return newNumbers
}

// func transformNumbers(numbers *[]int, transformation func(int) int) []int {
func transformAllNumbers[T int | float64](numbers *[]T, transformation func(T) T) []T {
	newNumbers := make([]T, len(*numbers))
	for index, item := range *numbers {
		newNumbers[index] = transformation(item)
	}

	return newNumbers
}

func double(value int) int {
	return value * 2
}

func doubleFloat(value float64) float64 {
	return value * 2
}

func getOperation() intTransformation {
	return double
}

func createTransformation(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func sumup(numbers ...int) int {
	sum := 0
	for _, val := range numbers {
		sum += val
	}

	return sum
}
