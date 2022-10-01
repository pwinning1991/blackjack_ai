package blackjack

import (
	"testing"

	"github.com/pwinning1991/deck"
)

var (
	opts = Options{
		Decks:           3,
		Hands:           2,
		BlackjackPayout: 1.5,
	}
)

//TODO figure how to test this with new hand type
//func TestEndHand(t *testing.T) {
//g := New(opts)
//g.playerBet = 150
//g.player = []hand{
//cards: []deck.Card{{Rank: deck.Ace},
//{Rank: deck.Ten},
//},
//bet: 150,
//}
//endRound(&g, HumanAI())
//got := g.balance
//want := int(150 * 1.5)
//if got != want {
//t.Errorf("Got %v, Wanted %v", got, want)
//}

//}

func TestBlackJack(t *testing.T) {
	hand1 := []deck.Card{
		{Rank: deck.Ace},
		{Rank: deck.Ten},
	}
	got := Blackjack(hand1...)
	want := true
	if got != want {
		t.Errorf("Got %v, Wanted %v", got, want)
	}

}

func TestMoveStand(t *testing.T) {
	g := New(opts)
	g.state = stateBet
	MoveStand(&g)
	if g.state != statePlayerTurn {
		t.Errorf("Should be statePlyaerTurn instead it is %v", g.state)
	}

}
