package main

import "fmt"

func main()  {
	revenue := userInput("Revenue: ")
	expenses := userInput("Expenses: ")
	taxRate := userInput("Tax rate: ")
	
	calcFinac(revenue, expenses, taxRate)

	
}

func userInput(inputTopic string) float64 {
	fmt.Print(inputTopic)
	var userInput float64
	fmt.Scan(&userInput)
	return userInput
}

func calcFinac(revenue, expenses, taxRate float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate / 100)
	ratio := ebt / profit
	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
} 