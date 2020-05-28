package main

import (
	"fmt"
	"math/rand"
	"time"
)

const rollDicesMax int = 50

func rollDices(res chan int) {

	min := 1
	max := 6

	res <- rand.Intn(max-min+1) + min
}

func rollDicesChannel() {

	rolls := make(chan int, rollDicesMax)

	go func() {
		var dicesCount int = 0
		rollRes1 := make(chan int)
		rollRes2 := make(chan int)

		defer close(rollRes1)
		defer close(rollRes2)
		defer close(rolls)

		// Required to seed RND
		rand.Seed(time.Now().UnixNano())

		for {

			go rollDices(rollRes1)
			go rollDices(rollRes2)

			r1, r2 := <-rollRes1, <-rollRes2

			rolls <- r1 + r2
			dicesCount++
			if dicesCount == rollDicesMax {
				break
			}
		}
	}()

	// Consume the rolls channel
	var c int = 1
	for rolled := range rolls {
		switch rolled {
		case 7, 11:
			fmt.Printf("IDX:%d DiceValue: %d, Outcome: %s\n", c, rolled, "NATURAL")
		case 2:
			fmt.Printf("IDX:%d DiceValue: %d, Outcome: %s\n", c, rolled, "SNAKE-EYES-CRAPS")
		case 3, 12:
			fmt.Printf("IDX:%d DiceValue: %d, Outcome: %s\n", c, rolled, "LOSS-CRAPS")
		default:
			fmt.Printf("IDX:%d DiceValue: %d, Outcome: %s\n", c, rolled, "NEUTRAL")
		}
		c++
	}
}

func main() {
	rollDicesChannel()
}
