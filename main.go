package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"
)

type Choice struct {
	name    string
	chances int
}

type BestScore struct {
	Attempts int       `json:"attempts"`
	Level    string    `json:"level"`
	Time     string    `json:"time"`
	Day      time.Time `json:"day"`
}

func main() {
	randomNumber := rand.Intn(100)

	choices := []Choice{}
	choices = append(choices, Choice{name: "Easy", chances: 10})
	choices = append(choices, Choice{name: "Medium", chances: 5})
	choices = append(choices, Choice{name: "Hard", chances: 3})

	run(choices, &randomNumber)

}
func intro() {
	fmt.Println("Welcome to the Number Guessing Game!\nI'm thinking of a number between 1 and 100.\nYou have 5 chances to guess the correct number.")
	fmt.Printf("\nPlease select the difficulty level:\n1. Easy (10 chances)\n2. Medium (5 chances)\n3. Hard (3 chances)")
	fmt.Println()
	fmt.Print("\nEnter your choice: ")
}
func game(choices []Choice, randomRumber int) {
	var bestscore *BestScore
	data, err := os.ReadFile("bestscore.json")
	if err != nil {
		fmt.Println("File read error :", err)
		return
	}
	if len(data) == 0 {
		*bestscore = BestScore{Attempts: 100000000}
	} else {
		err = json.Unmarshal(data, &bestscore)
		if err != nil {
			fmt.Println("Json unmarshall error :", err)
			return
		}
	}

	var choice int
	fmt.Scanln(&choice)
	if choice < 1 || choice > 3 {
		fmt.Println("Choice is invalid")
		return
	}
	chances := choices[choice-1].chances

	fmt.Println()
	fmt.Printf("Great! You have selected the %s difficulty level.\nLet's start the game!\n", choices[choice-1].name)
	start := time.Now()

	for i := 0; i < chances; i++ {
		fmt.Print("Enter your guess: ")
		var guess int
		fmt.Scanln(&guess)
		if guess == randomRumber {
			elapsed := time.Since(start)
			if i < bestscore.Attempts {
				fmt.Println("-----Best Score-----")
				bestscore.Attempts = i + 1
				bestscore.Day = time.Now()
				bestscore.Level = choices[choice-1].name
				bestscore.Time = fmt.Sprintf("%.2f", elapsed.Seconds())
			}
			data, err := json.Marshal(bestscore)
			if err != nil {
				fmt.Println("json marshall error :", err)
				return
			}
			err = os.WriteFile("bestscore.json", data, 0644)
			if err != nil {
				fmt.Println("File write error:", err)
				return
			}

			fmt.Printf("Congratulations! You guessed the correct number in %d attempts in a %.2f second.\n", i+1, elapsed.Seconds())
			return
		} else {
			if guess > randomRumber {
				fmt.Println("Incorrect! The number is less than", guess)
			} else {
				fmt.Println("Incorrect! The number is greater than", guess)
			}
		}
	}
	fmt.Println("Game Over!!!\nRandom number:", randomRumber)
}

func run(choices []Choice, randomNumber *int) {
	for true {

		var play string
		intro()
		game(choices, *randomNumber)

		for play != "y" && play != "n" {
			fmt.Print("Do you want to play again? Yes -> y || No -> n : ")
			fmt.Scanln(&play)
		}
		if play == "n" {
			return
		} else {
			fmt.Printf("\n\n")
			fmt.Println("New Game Started!!!")
			*randomNumber = rand.Intn(100)
			fmt.Printf("\n\n")
		}
	}
}
