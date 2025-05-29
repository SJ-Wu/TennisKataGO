package ExtremeSpeed

import "math"

type Tennis struct {
	firstPlayerScore  int
	secondPlayerScore int
	firstPlayerName   string
	secondPlayerName  string
}

func NewTennis(firstPlayerName string, secondPlayerName string) *Tennis {
	return &Tennis{firstPlayerName: firstPlayerName, secondPlayerName: secondPlayerName}
}

var scoreLookup = map[int]string{
	0: "love",
	1: "fifteen",
	2: "thirty",
	3: "forty",
}

func (t *Tennis) Score() string {
	if t.isScoreDifferent() {
		if t.isGameSetPoint() {
			return t.advState()
		}
		return t.diffScore()
	}
	if t.isDeuce() {
		return t.deuceScore()
	}
	return t.sameScore()
}

func (t *Tennis) advState() string {
	if t.isAdv() {
		return t.advScore()
	}
	return t.winScore()
}

func (t *Tennis) sameScore() string {
	return scoreLookup[t.firstPlayerScore] + " all"
}

func (t *Tennis) deuceScore() string {
	return "deuce"
}

func (t *Tennis) isDeuce() bool {
	return t.firstPlayerScore >= 3 || t.secondPlayerScore >= 3
}

func (t *Tennis) isScoreDifferent() bool {
	return t.firstPlayerScore != t.secondPlayerScore
}

func (t *Tennis) diffScore() string {
	return scoreLookup[t.firstPlayerScore] + " " + scoreLookup[t.secondPlayerScore]
}

func (t *Tennis) isGameSetPoint() bool {
	return t.firstPlayerScore > 3 || t.secondPlayerScore > 3
}

func (t *Tennis) isAdv() bool {
	return math.Abs(float64(t.firstPlayerScore-t.secondPlayerScore)) == 1
}

func (t *Tennis) advScore() string {
	return t.getAdvPlayer() + " adv."
}

func (t *Tennis) winScore() string {
	return t.getAdvPlayer() + " win."
}

func (t *Tennis) getAdvPlayer() string {
	var playerName string
	if t.firstPlayerScore > t.secondPlayerScore {
		playerName = t.firstPlayerName
	} else {
		playerName = t.secondPlayerName
	}
	return playerName
}

func (t *Tennis) FirstPlayerScore() {
	t.firstPlayerScore++
}

func (t *Tennis) SecondPlayerScore() {
	t.secondPlayerScore++
}
