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
func TestEndHand(t *testing.T) {
	g := New(opts)
	g.player = []hand{{
		cards: []deck.Card{
			{Rank: deck.Ace},
			{Rank: deck.Ten},
		},
		bet: 150,
	}}
	g.dealer = []deck.Card{
		{Rank: deck.Five},
		{Rank: deck.Four},
	}
	endRound(&g, HumanAI())
	got := g.balance
	want := int(150 * 1.5)
	if got != want {
		t.Errorf("Got %v, Wanted %v", got, want)
	}

}

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
	g.state = statePlayerTurn
	MoveStand(&g)
	if g.state != stateDealerTurn {
		t.Errorf("Should be statePlyaerTurn instead it is %v", g.state)
	}

}
