package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	var difficulty int
	var chances int
	var diff string
	var guess int

	attempts := 0

	randomNumber := rand.IntN(100) + 1

	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100.")

	fmt.Println("Please select the difficulty level:")
	fmt.Println("1. Easy (10 chances)")
	fmt.Println("2. Medium (5 chances)")
	fmt.Println("3. Hard (3 chances)")

	fmt.Print("Enter your choice: ")
	_, err := fmt.Scan(&difficulty)
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return
	}

	switch difficulty {
	case 1:
		chances = 10
		diff = "Easy"
	case 2:
		chances = 5
		diff = "Medium"
	case 3:
		chances = 3
		diff = "Hard"
	default:
		fmt.Println("Invalid choice, exiting game.")
		return
	}

	fmt.Printf("Great! You have selected the %v difficulty level. You have %v chances.\n", diff, chances)
	fmt.Println("Let's start the game!")

	for chances > 0 {
		fmt.Println("Enter your guess: ")
		_, err = fmt.Scan(&guess)
		if err != nil {
			fmt.Println("Error reading input: ", err)
			return
		}

		chances--
		attempts++

		if guess > randomNumber {
			fmt.Printf("Incorrect! The number is less than %v.\n", guess)
		} else if guess < randomNumber {
			fmt.Printf("Incorrect! The number is greater than %v.\n", guess)
		} else {
			fmt.Println("Congratulations! You guessed the correct number!")
			return
		}
	}

	fmt.Printf("\nGame over! You ran out of chances. The number was %v.\n", randomNumber)
}
