package main

// A = Rock, B = Paper, C = Scissors
type Move string

const (
	A Move = "A"
	B Move = "B"
	C Move = "C"
)

type scenario string

const (
	Win  scenario = "win"
	Draw scenario = "draw"
	Lose scenario = "lose"
)

var AllPossible map[Move]map[scenario]int = map[Move]map[scenario]int{
	A: {
		Win:  2 + 6,
		Draw: 1 + 3,
		Lose: 3 + 0,
	},
	B: {
		Win:  3 + 6,
		Draw: 2 + 3,
		Lose: 1 + 0,
	},
	C: {
		Win:  1 + 6,
		Draw: 3 + 3,
		Lose: 2 + 0,
	},
}

func (m Move) CalculateScoreOnScenario(s scenario) int {
	return AllPossible[m][s]
}

type roundCombinations string

const (
	AY roundCombinations = "AY"
	AX roundCombinations = "AX"
	AZ roundCombinations = "AZ"
	BY roundCombinations = "BY"
	BX roundCombinations = "BX"
	BZ roundCombinations = "BZ"
	CY roundCombinations = "CY"
	CX roundCombinations = "CX"
	CZ roundCombinations = "CZ"
)

type yourMoves string

const (
	X yourMoves = "Rock"
	Y yourMoves = "Paper"
	Z yourMoves = "Scissors"
)

type roundScores map[roundCombinations]int

func NewRoundScoresMap() roundScores {
	r := map[roundCombinations]int{
		AX: 3,
		AY: 6,
		AZ: 0,
		BY: 3,
		BZ: 6,
		BX: 0,
		CZ: 3,
		CX: 6,
		CY: 0,
	}

	y := map[yourMoves]int{
		"Rock":     1,
		"Paper":    2,
		"Scissors": 3,
	}

	// A = Rock, B = Paper, C = Scissors
	// X = Rock, Y = Paper, Z = Scissors

	return map[roundCombinations]int{
		AY: y[Y] + r[AY],
		AX: y[X] + r[AX],
		AZ: y[Z] + r[AZ],
		BY: y[Y] + r[BY],
		BX: y[X] + r[BX],
		BZ: y[Z] + r[BZ],
		CY: y[Y] + r[CY],
		CX: y[X] + r[CX],
		CZ: y[Z] + r[CZ],
	}
}

func (rs roundScores) CalculateRoundScore(round string) int {
	s, ok := rs[roundCombinations(round)]
	if !ok {
		return 0
	}

	return s
}

func (rs roundScores) ConvertToTotals(rounds []string) []int {
	totals := make([]int, 0)
	for _, r := range rounds {
		s := rs.CalculateRoundScore(r)
		totals = append(totals, s)
	}

	return totals
}
