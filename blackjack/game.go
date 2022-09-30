package blackjack

import (
	"fmt"

	"github.com/pwinning1991/deck"
)

type state int8

const (
	statePlayerTurn State = iota
	stateDealerTurn
	stateHandOver
)

func New() Game {
	return Game{
		state:    statePlayerTurn,
		dealerAi: dealerAI{},
		balance:  0,
	}
}

type Game struct {
	// unexported fields
	deck     []deck.Card
	state    state
	player   []deck.Card
	dealer   []deck.Card
	dealerAI AI
	balance  int
}

func (g *Game) currentHand() *[]deck.Card {
	switch g.state {
	case statePlayerTurn:
		return &g.player
	case stateDealerTurn:
		return &g.dealer
	default:
		panic("it isn't current any player's turn")
	}

}

func deal(gs GameState) {
	g := clone(gs)
	g.player = make([]deck.Card, 0, 5)
	g.dealer = make([]deck.Card, 0, 5)
	var card deck.Card
	for i := 0; i < 2; i++ {
		card, g.Deck = draw(g.deck)
		g.player = append(g.player, card)
		card, g.Deck = draw(g.deck)
		g.Dealer = append(g.dealer, card)
	}
	g.state = StatePlayerTurn
}

func (g *Game) Play(ai AI) {
	g.deck = dec.New(deck.Deck(3), deck.Shuffle())
	for i := 0; i < 2; i++ {
		deal(g)
		var input string
		for g.state == statePlayerTurn {
			hand := make([]deck.Card, len(g.player))
			copy(hand, g.player)
			move := ai.Play(hand, g.dealer[0])
			move(g)
		}

		for gs.State == StateDealerTurn {
			hand := make([]deck.Card, len(g.dealer))
			copy(hand, g.dealer)
			move := g.dealerAi.Play(hand, g.dealer[0])
			move(g)
		}
		endHand(g, ai)

	}
}

type Move func(*Game)

func MoveHit(g *Game) {
	hand = g.currentHand()
	var card deck.Card
	card, g.deck = draw(g.deck)
	*hand = append(*hand, card)
	if score(*hand...) > 21 {
		MoveStand(g)
	}
}
func MoveStand(g *Game) {
	g.State++
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

func minScore(hand ...deck.Card) int {
	runningTotal := 0
	for _, card := range hand {
		runningTotal += min(int(card.Rank), 10)
	}
	return runningTotal
}

func Soft(hand ...deck.Card) bool {
	minScore := minScore(hand...)
	score := Score(hand)
	return minScore != score
}

func Score(hand ...deck.Card) int {
	minScore := minScore(hand...)
	if minScore > 11 {
		return minScore
	}

	for _, card := range hand {
		if card.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func endHand(g *Game, ai AI) {
	pScore, dScore := Score(g.player...), Score(g.dealer...)
	//TODO : figure ouyt winnings and add/sub
	switch {
	case pScore > 21:
		fmt.Println("You Busted")
		g.balance--
	case dScore > 21:
		fmt.Println("Dealer Busted")
		g.balance++
	case pScore > dScore:
		fmt.Println("You Win")
		g.balance++
	case pScore < dScore:
		fmt.Println("You Lose")
		g.balance--
	case pScore == dScore:
		fmt.Println("Draw")
	}
	fmt.Println()
	ai.Results([][]deck.Card{g.player}, g.dealer)
	g.player = nil
	g.dealer = nil
}
