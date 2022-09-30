package ai

import (
	"fmt"

	"github.com/pwinning1991/deck"
)

type AI interface {
	Bet() int
	Play(hand []deck.Card, dealer deck.Card)
	Results(hand [][]deck.Card, dealer []deck.Card)
}

type HumanAI struct{}

func (ai *HumanAI) Bet() int {
	return 1
}

type dealerAI struct{}

func (ai dealerAI) Bet() int {
	return 1
}

func (ai dealerAI) Play(hand []deck.Card, dealer deck.Card) Move {
	dScore := Score(hand...)
	if dScore <= 16 || dScore == 17 && Soft(g.dealer...) {
		return MoveHit
	} else {
		return MoveStand
	}
}

func (ai dealerAI) Results(hand [][]deck.Card, dealer []deck.Card) {
	// noop
}

func (ai *HumanAI) Play(hand []deck.Card, dealer deck.Card) {
	var input string
	for {
		fmt.Println("Player:", player)
		fmt.Println("Dealer:", dealer)
		fmt.Println("What will you do? (h)it, (s)tand")
		fmt.Scanf("%s\n", &input)
		switch input {
		case "h":
			return MoveHit()
		case "s":
			return MoveStand()
		default:
			fmt.Println("Invalid option: chose either \"h\" or \"s\"")
		}

	}

}

func (ai *HumanAI) Results(hand [][]deck.Card, dealer []deck.Card) {
	fmt.Println("==FINAL HANDS==")
	fmt.Println("Player:", hand)
	fmt.Println("Dealer:", dealer)

}
