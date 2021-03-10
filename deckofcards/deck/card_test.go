package deck

import (
	"fmt"
	"testing"

	"github.com/marcosgeo/go/deckofcards/deck"
)

func ExampleCard() {
	fmt.Println(deck.Card{Rank: deck.Ace, Suit: deck.Heart})
	fmt.Println(deck.Card{Rank: deck.Two, Suit: deck.Spade})
	fmt.Println(deck.Card{Rank: deck.Nine, Suit: deck.Diamond})
	fmt.Println(deck.Card{Rank: deck.Jack, Suit: deck.Club})
	fmt.Println(deck.Card{Suit: deck.Joker})

	// Output:
	// Ace of Hearts
	// Two of Spades
	// Nine of Diamonds
	// Jack of Clubs
	// Joker
}

func TestNew(t *testing.T) {
	cards := deck.New()
	// 13 ranks * 4 suits
	if len(cards) != 13*4 {
		t.Error("Wrong number of cardds in a new deck.")
	}
}

func TestSort(t *testing.T) {
	cards := deck.New(deck.Sort(deck.Less))
	exp := deck.Card{Rank: deck.Ace, Suit: deck.Spade}
	if cards[0] != exp {
		t.Error("Expected Ace of Spades as first card. Received:", cards[0])
	}
}

func TestShuffle(t *testing.T) {
	// make shuffleRand deterministic
	// First call to shuffleRand.Perm(52) should be:
	// [40, 35, ...]
	//shuffleRand = rand.New(rand.NewSource(0))

	orig := deck.New()
	first := orig[40]
	second := orig[35]
	cards := deck.New(deck.Shuffle)
	if cards[0] != first {
		t.Errorf("Expected the first card to be %s, received %s", first, cards[0])
	}
	if cards[1] != second {
		t.Errorf("Expected the first card to be %s, received %s.", first, second)
	}
}

func TestJokers(t *testing.T) {
	want := 3
	cards := deck.New(deck.Jokers(want))
	count := 0
	for _, c := range cards {
		if c.Suit == deck.Joker {
			count++
		}
	}
	if count != want {
		t.Errorf("Expected %d Jokers, received: %d", want, count)
	}
}

func TestFilter(t *testing.T) {
	filter := func(card deck.Card) bool {
		return card.Rank == deck.Two || card.Rank == deck.Three
	}
	cards := deck.New(deck.Filter(filter))
	for _, c := range cards {
		if c.Rank == deck.Two || c.Rank == deck.Three {
			t.Errorf("Expected all twos and threes to be filtered out.")
		}
	}
}

func TestDeck(t *testing.T) {
	decks := 3
	// 13 ranks * 4 suites * N decks
	expected := 13 * 4 * decks
	cards := deck.New(deck.Deck(decks))
	if len(cards) != expected {
		t.Errorf("Expected %d cards, received %d cards.", expected, len(cards))
	}
}
