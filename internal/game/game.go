package game

import (
	"math/rand"
)

func GenerateBotChoice() string {
	choices := []string{"Rock", "Paper", "Scissors"}
	return choices[rand.Intn(len(choices))]
}

func CompareChoices(userChoice, botChoice string) string {
	if userChoice == botChoice {
		return "tie"
	}
	switch userChoice {
	case "Rock":
		if botChoice == "Scissors" {
			return "win"
		}
	case "Paper":
		if botChoice == "Rock" {
			return "win"
		}
	case "Scissors":
		if botChoice == "Paper" {
			return "win"
		}
	}
	return "lose"
}
