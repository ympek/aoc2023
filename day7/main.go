package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	HighCard     = 0
	OnePair      = 1
	TwoPair      = 2
	ThreeOfAKind = 3
	FullHouse    = 4
	FourOfAKind  = 5
	FiveOfAKind  = 6
)

type biddedHand struct {
	hand string
	bid  int
}

func main() {
	contents, err := os.ReadFile("./input")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(contents), "\n")

	var hands []biddedHand

	for _, line := range lines {
		if len(line) < 1 {
			continue
		}
		fields := strings.Fields(line)
		hand := fields[0]
		bid, _ := strconv.Atoi(fields[1])
		hands = append(hands, biddedHand{
			hand: hand,
			bid:  bid,
		})
	}

	slices.SortFunc(hands, compareHands)
	ansPart1 := 0
	for i, h := range hands {
		ansPart1 += (i + 1) * h.bid
	}
	fmt.Printf("Answer to part 1: %d\n", ansPart1)

	slices.SortFunc(hands, compareHandsWithJokers)
	ansPart2 := 0
	for i, h := range hands {
		ansPart2 += (i + 1) * h.bid
	}
	fmt.Printf("Answer to part 2: %d\n", ansPart2)
}

func calcNumericValueOfACard(val rune) int {
	switch val {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	}
	return int(val - '0')
}

func calcNumericValueOfACardWithJokers(val rune) int {
	switch val {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 1
	case 'T':
		return 10
	}
	return int(val - '0')
}

func calcHandType(hand string) int {
	set := make(map[rune]int)

	for _, s := range hand {
		val, exists := set[s]
		if exists {
			set[s] = val + 1
		} else {
			set[s] = 1
		}
	}

	numDistinct := len(set)

	if numDistinct == 1 {
		return FiveOfAKind
	}
	if numDistinct == 2 {
		for _, x := range set {
			if x == 4 {
				return FourOfAKind
			}
		}
		return FullHouse
	}
	if numDistinct == 3 {
		for _, x := range set {
			if x == 3 {
				return ThreeOfAKind
			}
		}
		return TwoPair
	}
	if numDistinct == 4 {
		return OnePair
	}
	if numDistinct == 5 {
		return HighCard
	}
	panic("Illegal hand type")
}

func calcHandTypeWithJokers(hand string) int {
	set := make(map[rune]int)

	for _, s := range hand {
		val, exists := set[s]
		if exists {
			set[s] = val + 1
		} else {
			set[s] = 1
		}
	}
	numJokers := set['J']

	if numJokers == 5 {
		return FiveOfAKind
	}
	if numJokers == 0 {
		return calcHandType(hand)
	}

	// yeah, let's just brute-force the solution cuz why not xD
	potentialHighestType := 0
	substitutions := []string{"A", "K", "Q", "J", "T", "9", "8", "7", "6", "5", "4", "3", "2"}
	for i := 1; i <= numJokers; i++ {
		for j := 0; j < len(substitutions); j++ {
			newHand := strings.Replace(hand, "J", substitutions[j], i)
			ht := calcHandType(newHand)
			if ht >= potentialHighestType {
				potentialHighestType = ht
			}
		}
	}

	return potentialHighestType
}

func compareHands(a, b biddedHand) int {
	// as per docs:
	// cmp(a, b) should return a negative number when a < b, a positive number when a > b and zero when a == b.
	htA := calcHandType(a.hand)
	htB := calcHandType(b.hand)
	if htA > htB {
		return 1
	}
	if htA < htB {
		return -1
	}
	// now compare 1by1
	runesA := []rune(a.hand)
	runesB := []rune(b.hand)

	for i := 0; i < len(runesA); i++ {
		valA := calcNumericValueOfACard(runesA[i])
		valB := calcNumericValueOfACard(runesB[i])
		if valA > valB {
			return 1
		}
		if valA < valB {
			return -1
		}
	}
	panic("Illegal")
}

func compareHandsWithJokers(a, b biddedHand) int {
	htA := calcHandTypeWithJokers(a.hand)
	htB := calcHandTypeWithJokers(b.hand)
	if htA > htB {
		return 1
	}
	if htA < htB {
		return -1
	}

	runesA := []rune(a.hand)
	runesB := []rune(b.hand)

	for i := 0; i < len(runesA); i++ {
		valA := calcNumericValueOfACardWithJokers(runesA[i])
		valB := calcNumericValueOfACardWithJokers(runesB[i])
		if valA > valB {
			return 1
		}
		if valA < valB {
			return -1
		}
	}
	panic("Illegal")
}
