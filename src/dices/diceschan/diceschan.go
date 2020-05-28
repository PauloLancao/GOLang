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
		rollRes1 := make(chan int)
		rollRes2 := make(chan int)

		defer close(rollRes1)
		defer close(rollRes2)
		defer close(rolls)

		// Required to seed RND
		rand.Seed(time.Now().UnixNano())

		for i := 0; i < rollDicesMax; i++ {
			go rollDices(rollRes1)
			go rollDices(rollRes2)

			r1, r2 := <-rollRes1, <-rollRes2

			rolls <- r1 + r2
		}
	}()

	// Consume the rolls channel
	for rolled := range rolls {
		switch rolled {
		case 7, 11:
			fmt.Printf("DiceValue: %d, Outcome: %s\n", rolled, "NATURAL")
		case 2:
			fmt.Printf("DiceValue: %d, Outcome: %s\n", rolled, "SNAKE-EYES-CRAPS")
		case 3, 12:
			fmt.Printf("DiceValue: %d, Outcome: %s\n", rolled, "LOSS-CRAPS")
		default:
			fmt.Printf("DiceValue: %d, Outcome: %s\n", rolled, "NEUTRAL")
		}
	}
}

func main() {
	rollDicesChannel()
}
