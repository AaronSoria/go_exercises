package main // todos necesitan esta seccion

import (
	"fmt"
	"math"
)

func main() {
	var revenue, expenses, taxRate float64
	fmt.Print("revenue ")
	fmt.Scan(&revenue)
	fmt.Print("expenses ")
	fmt.Scan(&expenses)
	fmt.Print("taxRate ")
	fmt.Scan(&taxRate)

	ebt, profit, ratio := computeProfits(revenue, expenses, taxRate)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func computeProfits(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

func main_1() {
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

	// futureValue := investmentAmount * math.Pow((1+(expectedReturnRate/100)), years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
	futureValue, futureRealValue := calculateFutureValue(investmentAmount, expectedReturnRate, inflationRate, years)
	fmt.Println(`future Value:
	
	`, futureValue)
	fmt.Printf("Future VAlue (adjusted for Inflation): %#v", futureRealValue)
}

func outputText(text1, text2 string) {
	fmt.Printf(text1)
	fmt.Printf(text2)
}

// func calculateFutureValue(investmentAmount float64, expectedReturnRate float64, inflationRate float64, years float64) (float64, float64) {
// 	futureValue := investmentAmount * math.Pow((1+(expectedReturnRate/100)), years)
// 	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)
// 	return futureValue, futureRealValue
// }

func calculateFutureValue(investmentAmount float64, expectedReturnRate float64, inflationRate float64, years float64) (fv float64, frv float64) {
	fv = investmentAmount * math.Pow((1+(expectedReturnRate/100)), years)
	frv = fv / math.Pow(1+inflationRate/100, years)
	return
}
