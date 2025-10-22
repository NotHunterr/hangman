// TODO
// WHEN GUESSER GUESSES, CHECK IF ITS CORRECTLY A LETTER AND NOT A STRING, IF WRONG MAKE THEM GUESS AGAIN
// BE ABLE TO SEE YOUR PREVIOUS GUESSES
// DETERMINE CHOSEN WORD AND PRELOAD ALL CHARACTER SPACES BEFORE FOR LOOP
//DISPLAY CURRENT SCORE E.X: H_APP_ (HAPPY)

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	var startGameASCII = `
    +---+
    |   |
        |
        |
        |
        |
    =========
	`

	randWords := []string{"Hunter", "Brook", "Love", "Party", "Dingleberries"}

	randIndex := rand.Intn(len(randWords))

	chosenWord := randWords[randIndex]
	fmt.Println(startGameASCII)
	fmt.Println("Welcome to Hangman! Your only goal is to guess the random word that the computer has chosen!")
	var won bool = false
	var letter string
	var misses int = 0

	for !won {
		fmt.Println("Please take your guess now!")
		fmt.Scanln(&letter)

		if !strings.Contains(chosenWord, letter) {
			misses += 1
		} else {

			//iterate over chosen word with selected letter to find the spot that it is in
			fmt.Printf("The letter %s is in the secret phrase! It is in the %s spot", letter)
		}

		switch misses {
		case 1:
			fmt.Println(oneL)
		case 2:
			fmt.Println(twoL)
		case 3:
			fmt.Println(threeL)
		case 4:
			fmt.Println(fourL)
		case 5:
			fmt.Println(fiveL)
		case 6:
			fmt.Println(sixL)
			fmt.Printf("You have lost I'm so very sorry! The correct word was %s", chosenWord)
		}

	}

}
