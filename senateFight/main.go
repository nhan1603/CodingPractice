package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(predictPartyVictory("DRR"))
}

func predictPartyVictory(senate string) string {
	partCount := make(map[byte]int)
	partCount['R'] = 0
	partCount['D'] = 0
	for strings.Contains(senate, "R") && strings.Contains(senate, "D") {
		val := senate[0]
		// increase a party power by present
		partCount[val] = partCount[val] + 1
		// decrease a party power by
		// - banning an opponent
		// - take slot for a ban vote
		partCount[opposite(val)] = partCount[opposite(val)] - 1
		if partCount[val] > 0 {
			senate = senate[1:]
			senate = senate + string(val)
		} else {
			senate = senate[1:]
		}
	}
	if !strings.Contains(senate, "R") {
		return "Dire"
	}
	return "Radiant"
}

func opposite(party byte) byte {
	if party == 'R' {
		return 'D'
	}
	return 'R'
}
