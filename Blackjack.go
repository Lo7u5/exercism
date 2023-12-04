package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch card {
        case "ace":
    	return 11
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
        case "ten":
    	return 10
        case "jack":
    	return 10
        case "queen":
    	return 10
        case "king":
    	return 10
    }
	return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	switch{
        case card1 == "ace" && card2 == "ace":
    	return "P"
        case ParseCard(card1) + ParseCard(card2) == 21 && ParseCard(dealerCard) == 11:
    	return "S"
        case ParseCard(card1) + ParseCard(card2) == 21 && ParseCard(dealerCard) == 10:
    	return "S"
        case ParseCard(card1) + ParseCard(card2) == 21:
    	return "W"
        case 17 <= ParseCard(card1) + ParseCard(card2) && ParseCard(card1) + ParseCard(card2) <= 20:
    	return "S"
        case 12 <= ParseCard(card1) + ParseCard(card2) && ParseCard(card1) + ParseCard(card2) <= 16 && ParseCard(dealerCard) >=7:
    	return "H"
        case 12 <= ParseCard(card1) + ParseCard(card2) && ParseCard(card1) + ParseCard(card2) <= 16:
    	return "S"
        case ParseCard(card1) + ParseCard(card2) <= 11:
    	return "H"
    }
	return "Wrong input"
}
