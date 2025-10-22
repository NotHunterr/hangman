// TODO
// WHEN GUESSER GUESSES, CHECK IF ITS CORRECTLY A LETTER AND NOT A STRING, IF WRONG MAKE THEM GUESS AGAIN
// BE ABLE TO SEE YOUR PREVIOUS GUESSES
// DETERMINE CHOSEN WORD AND PRELOAD ALL CHARACTER SPACES BEFORE FOR LOOP
//DISPLAY CURRENT SCORE E.X: H_APP_ (HAPPY)
// MAKE SURE YOU CANT GUESS THE SAME THING AGAIN
// CHECK FOR TWO OF SAME LETTER IN WORD.

package main

import (
	"fmt"
	"math/rand"
	"strconv"
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

	secretWord := randWords[randIndex]
	fmt.Println(startGameASCII)
	fmt.Println("Welcome to Hangman! Your only goal is to guess the random word that the computer has chosen!")
	fmt.Println(secretWord)
	var won bool = false
	var letter string
	var misses int = 0
	var totalRight int = 0

	// Game Loop
	for !won {
		fmt.Println("Please take your guess now!")
		fmt.Scanln(&letter)

		if !strings.Contains(secretWord, letter) {
			misses += 1
			fmt.Printf("I'm sorry, the letter %s is not in the word!", letter)

		} else {
			// Inits length of secret word, the creates a loop and iterates over each letter,
			// checking if it is the same as the one guessed
			wordLen := len(secretWord)
			i := 0

			for i < wordLen {
				slice := secretWord[i:(i + 1)]
				if slice != letter {
					i += 1
				} else {
					fmt.Printf("You are correct! The letter %s is in the %s spot\n", letter, strconv.Itoa(i+1))
					totalRight += 1

					if totalRight == wordLen {
						fmt.Printf("Congratulations, you have solved the word! The word was %s", secretWord)
						won = true
						break
					}
					break
				}

			}
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
			fmt.Printf("You have lost I'm so very sorry! The correct word was %s", secretWord)

		}
		// If you have missed 6 letters, the game is over
		if misses == 6 {
			break
		}

	}

}
