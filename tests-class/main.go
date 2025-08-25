package main

import "errors"

func main() {
	result, err := divide(x,y)
}

func divide (x,y float64 ) (float64, error) {
	var result float64

	if y == 0 {
		return result, errors.New("Cannot divide by 0")
	}

	return x/y, nil
}