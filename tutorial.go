package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fatih/color"
)

func main() {
	PlayGame()
}

//export PlayGame
func PlayGame() {
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().Unix())
	number := rand.Int63n(100)
	game_won := false

	guessed_numbers := []int64{}

	for i := 5; i > 0; i-- {
		color.New(color.FgGreen).Printf("You have %d chances left to guess\n>", i)
		input, err := reader.ReadString('\n')
		HandleError(err)
		input = strings.TrimSuffix(input, "\n")
		user_guess, err := strconv.ParseInt(input, 10, 32)
		if err != nil || user_guess > 100 || user_guess < 0 {
			color.New(color.FgYellow).Println("Enter a valid number from 0-100!\n")
			i++
			continue
		}

		if i64InArray(guessed_numbers, user_guess) {
			color.New(color.FgYellow).Println("You have already guessed it!\n")
			i++
			continue
		}

		guessed_numbers = append(guessed_numbers, user_guess)

		if user_guess == number {
			fmt.Println("You've found it! It was", number)
			game_won = true
			break
		} else if user_guess > number {
			fmt.Println("Your guess is too high")
		} else {
			fmt.Println("Your guess is too low")
		}
		fmt.Print("\n")
	}

	if !game_won {
		color.New(color.FgRed).Println("You lost! The number was", number)
	}
}

func i64InArray(arr []int64, num int64) bool {
	for _, v := range arr {
		if v == num {
			return true
		}
	}
	return false
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
