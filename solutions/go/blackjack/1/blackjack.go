package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten", "jack", "queen", "king":
		return 10
	case "ace":
		return 11
	default:
		return 0
	}
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	card1Val, card2Val, dealerCardVal := ParseCard(card1), ParseCard(card2), ParseCard(dealerCard)
	// Naïve translation of the strategy described in the README:
	switch {
	case card1Val == card2Val && card1Val == 11:
		return "P"
	case card1Val+card2Val == 21 && (dealerCardVal != 10 && dealerCardVal != 11):
		return "W"
	case card1Val+card2Val == 21 && (dealerCardVal == 10 || dealerCardVal == 11):
		return "S"
	case inRange(card1Val+card2Val, 17, 20):
		return "S"
	case inRange(card1Val+card2Val, 12, 16) && dealerCardVal < 7:
		return "S"
	case inRange(card1Val+card2Val, 12, 16) && dealerCardVal >= 7:
		return "H"
	case inRange(card1Val+card2Val, 0, 11):
		return "H"
	default:
		// Suboptimal error reporting.
		panic("Unexpected combination.")
	}
}

func inRange(x, lo, hi int) bool {
	return lo <= x && x <= hi
}
