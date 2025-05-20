package main // todos necesitan esta seccion

import (
	"fmt"
	"math"
)

func main() {
	//var investmentAmount, years float64 = 1000, 10
	//investmentAmount, years := 1000.0, 10.0
	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5
	// expectedReturnRate := 5.5
	// var years float64 = 10

	futureValue := investmentAmount * math.Pow((1+(expectedReturnRate/100)), years)
	fmt.Println(futureValue)
}
