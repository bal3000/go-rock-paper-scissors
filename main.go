package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	guide, err := loadStrategyGuide("./guide.txt")
	if err != nil {
		panic(err)
	}

	t := chooseMovesAndGetTotals(guide)

	score := calculateTotal(t)

	fmt.Printf("Total Score: %d", score)
}

func calculateTotal(scores []int) int {
	total := 0

	for _, s := range scores {
		total += s
	}

	return total
}

func chooseMovesAndGetTotals(moves []string) []int {
	totals := make([]int, 0)
	for _, m := range moves {
		s := strings.Split(m, "")
		theirs := Move(s[0])
		strat := s[1]

		switch strat {
		case "X":
			// lose
			totals = append(totals, theirs.CalculateScoreOnScenario(Lose))
		case "Y":
			// draw
			totals = append(totals, theirs.CalculateScoreOnScenario(Draw))
		case "Z":
			// win
			totals = append(totals, theirs.CalculateScoreOnScenario(Win))
		}
	}
	return totals
}

func loadStrategyGuide(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	str := strings.Split(
		strings.Replace(
			string(data),
			"\r\n",
			"\n",
			-1,
		),
		"\n",
	)

	guide := make([]string, 0)

	for _, line := range str {
		round := strings.ReplaceAll(line, " ", "")
		guide = append(guide, round)
	}

	return guide, nil
}
