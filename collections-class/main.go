package main

import "fmt"

type gradings map[string]float64

func (g gradings) output() {
	for key, item := range g {
		fmt.Println(key, item)
	}
}

func main() {
	userNames := make([]string, 2, 5)
	userNames = append(userNames, "aaron", "allen")
	userNames[0] = "natalia"

	courseRatings := make(gradings, 1)

	courseRatings["c#"] = 4.5
	courseRatings["go"] = 4.0
	courseRatings["c"] = 5

	fmt.Println(userNames)
	for index, item := range userNames {
		fmt.Println(index, item)
	}
	courseRatings.output()
	//exercise()
}

func createMaps() {
	webs := map[string]string{
		"Google":             "https://google.com",
		"Amazon Web Service": "https://aws.com",
	}
	fmt.Println(webs)
	webs["linkedin"] = "https://linkedin.com"
	delete(webs, "Google")
	fmt.Println(webs)
}

func exercise() {
	var prices = []float64{0.99, 1.99}
	prices = append(prices, 2.99)
	fmt.Println(prices)

	fmt.Println("comienza el ejercicio")
	hobbies := [3]string{"coding", "reading", "running"}
	fmt.Println("1)", hobbies)

	fmt.Println("2)", hobbies[:1])
	fmt.Println("2)", hobbies[1:])

	auxHobbies := hobbies[:2]
	interesringHobbies1 := auxHobbies[0:]
	interesringHobbies2 := auxHobbies[:2]
	fmt.Println("3)", interesringHobbies1)
	fmt.Println("3)", interesringHobbies2)
	interesringHobbies1 = interesringHobbies2[1:3]
	fmt.Println("4)", interesringHobbies1)

	var goals []string
	goals = append(goals, "learn go", "build my webapi")
	fmt.Println(goals)
	fmt.Println("5)", goals)

	goals[1] = "Lear japanese"
	goals = append(goals, "get my phd")
	fmt.Println("6)", goals)

}

// func main() {
// 	prices := [5]float64{6.3, 88.3, 11.8, 1.03, 3.99} //    aqui se declarar que un array tiene 4 elementos del tipo float64
// 	var names [4]string
// 	fmt.Println(prices)
// 	fmt.Println(names)

// 	fmt.Println(prices[1:4])
// 	fmt.Println(prices[:3])
// 	auxPrices := prices[3:]
// 	auxPrices[1] = 0.0
// 	fmt.Println("auxPrices", auxPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(auxPrices))
// 	fmt.Println(cap(auxPrices))
// 	fmt.Println(prices)
// }
