package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

type participant struct {
	name, steamID      string
	isGifter, isGiftee bool
}

// Creates and returns a participant struct of user inputted data
// TODO: Add Reading from file later
func AddPerson() participant {

	var inputName string
	var inputSteamID string

	fmt.Println("Enter their name.")
	fmt.Scanln(&inputName)

	fmt.Println("Enter their steamID.")
	fmt.Scanln(&inputSteamID)

	temp := participant{name: inputName, steamID: inputSteamID, isGifter: false, isGiftee: false}

	return temp
}

// Returns a random number between 0 and number
func RandomNum(number int) int {
	r1 := rand.NewSource(time.Now().UnixNano())
	s1 := rand.New(r1)
	return s1.Intn(number)
}

func Randomize(users []participant) {
	if len(users) < 2 {
		fmt.Println("The number of participants must be > 2 in order to randomize. Please add more.")
	} else {
		// TODO: Account for odd number of participants
		fmt.Println("Randomizing.")

		var gifter int
		var giftee int
		matches := [][]int{}

		for i := 0; i < len(users); i++ {
			gifter = RandomNum(len(users))
			giftee = RandomNum(len(users))

			// If they have already been gifted/gifter re-randomize
			if users[gifter].isGifter == true {
				for users[gifter].isGifter == true {
					gifter = RandomNum(len(users))
				}
			}

			if users[giftee].isGiftee == true {
				for users[giftee].isGiftee == true {
					giftee = RandomNum(len(users))
				}
			}

			// Add the match to the array
			match := []int{gifter, giftee}

			matches = append(matches, match)

			users[gifter].isGifter = true
			users[giftee].isGiftee = true

		}

		fmt.Println("Matches:")
		for k := 0; k < len(users)/2; k++ {
			fmt.Println(users[matches[k][0]], "->", users[matches[k+1][0]])
		}
	}
}

func Menu() {
	var users []participant
	var userInput string
	notExit := false

	for notExit != true {
		fmt.Println("Enter 1 to add a person to the list.\nEnter 2 to randomize.\nEnter anything else to exit.")
		fmt.Scanln(&userInput)

		switch userInput {
		case "1":
			users = append(users, AddPerson())
		case "2":
			Randomize(users)
		default:
			fmt.Println("Goodbye!")
			os.Exit(3)
		}
	}
}

func main() {
	Menu()
}
