package main

import (
	"fmt"
	"math/rand"
)

var upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var lower = "abcdefghijklmnopqrstuvwxyz"
var number = "0123456789"

func main() {
	var uppercase, lowercase, digit, symbol int

	fmt.Println("Enter: ")
	fmt.Scanf("%d %d %d %d", &uppercase, &lowercase, &digit, &symbol)

	password := generate(uppercase, lowercase, digit, symbol)
	fmt.Println(password)
}

func generate(uppercase, lowercase, digit, symbol int) string {
	if uppercase+lowercase+digit > symbol && uppercase < 0 && lowercase < 0 && digit < 0 && symbol < 0 {
		fmt.Println("Please try again...")
		return ""
	}

	randomUpper := make([]byte, uppercase)
	randomLower := make([]byte, lowercase)
	randomDigit := make([]byte, digit)
	randomSymbol := make([]byte, symbol-(uppercase+lowercase+digit))

	//upperPass := upper[rand.Intn(uppercase)]

	for i := range randomUpper {
		randomUpper[i] = upper[rand.Intn(uppercase)]
	}

	for i := range randomLower {
		randomLower[i] = lower[rand.Intn(lowercase)]
	}

	for i := range randomDigit {
		randomDigit[i] = number[rand.Intn(digit)]
	}

	for i := range randomSymbol {
		randomSymbol[i] = upper[rand.Intn(symbol)]
	}

	password := fmt.Sprintf("%s%s%s%s", string(randomUpper), string(randomLower), string(randomDigit), string(randomSymbol))

	return password
}
