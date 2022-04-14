package blackjack

func ParseCard(card string) int {
	switch card {
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
		case "ace":
			return 11
		default:
			return 0
	}
}

func IsBlackjack(card1, card2 string) bool {
	return (ParseCard(card1) + ParseCard(card2)) == 21
}

func LargeHand(isBlackjack bool, dealerScore int) string {
	switch {
	case isBlackjack && dealerScore < 10:
		return "W"
	case isBlackjack && dealerScore >= 10:
		return "S"
	default:
		return "P"
	}
}

func SmallHand(handScore, dealerScore int) string {
	switch {
	case handScore >= 17:
		return "S"
	case handScore <= 11:
		return "H"
	case dealerScore >= 7:
		return "H"
	default:
		return "S"
	}
}
