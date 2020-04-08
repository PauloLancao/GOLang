package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

const rollDicesMax int = 50

func rollDices() (int, error) {
	min := 1
	max := 6

	var r1 = rand.Intn(max-min+1) + min

	if r1 < 1 || r1 > 6 {
		return 0, errors.New("Dices have to be between 1 and 6")
	}

	return r1, nil
}

func rollDicesNTimes() []int {
	var dicesCount int = 0
	var resultRollDices = make([]int, rollDicesMax)
	rand.Seed(time.Now().UnixNano())

	for {
		r1, err := rollDices()
		r2, err1 := rollDices()

		if err == nil && err1 == nil {
			resultRollDices[dicesCount] = r1 + r2
			dicesCount++
			if dicesCount == rollDicesMax {
				break
			}
		}
	}

	return resultRollDices
}

func evaluateDicesOutcome(results []int) {

	for diceVal := range results {
		switch results[diceVal] {
		case 7, 11:
			fmt.Printf("DiceValue: %d, Outcome: %s\n", results[diceVal], "NATURAL")
		case 2:
			fmt.Printf("DiceValue: %d, Outcome: %s\n", results[diceVal], "SNAKE-EYES-CRAPS")
		case 3, 12:
			fmt.Printf("DiceValue: %d, Outcome: %s\n", results[diceVal], "LOSS-CRAPS")
		default:
			fmt.Printf("DiceValue: %d, Outcome: %s\n", results[diceVal], "NEUTRAL")
		}
	}
}

// Write a program so that two dice (1 to 6) are rolled 50 times.
// The resulting rolls are to be processed in the following manner.
// 7 and 11 are to be called NATURAL,
// 2 SNAKE-EYES-CRAPS ,
// 3, 12 are LOSS-CRAPS,
// any other combinations are to be called NEUTRAL.
// Display the number rolls and the outcomes in sequential order.
func main() {
	evaluateDicesOutcome(rollDicesNTimes())
}
