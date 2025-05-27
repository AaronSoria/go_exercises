package main // todos necesitan esta seccion

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 6.5
	//var investmentAmount, years float64 = 1000, 10
	//investmentAmount, years := 1000.0, 10.0
	var investmentAmount float64
	years, expectedReturnRate := 10.0, 5.5
	// expectedReturnRate := 5.5
	// var years float64 = 10

	fmt.Print("please specify an investmentAmount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("please specify years of investment: ")
	fmt.Scan(&years)

	fmt.Print("please specify an expected rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow((1+(expectedReturnRate/100)), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(`future Value:
	
	`, futureValue)
	fmt.Printf("Future VAlue (adjusted for Inflation): %#v", futureRealValue)
}
