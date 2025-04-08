package TennisKataGO

import "math"

type Tennis struct {
	firstPlayerScoreTime  int
	secondPlayerScoreTime int
	firstPlayerName       string
	secondPlayerName      string
}

var scoreLookup = map[int]string{
	0: "love",
	3: "forty",
	2: "thirty",
	1: "fifteen",
}

func (t *Tennis) Score() string {
	if t.isScoreDifferent() {
		if t.firstPlayerScoreTime > 3 || t.secondPlayerScoreTime > 3 {
			return t.advState()
		}
		return t.differentScore()
	}
	if t.isDeuce() {
		return t.deuceScore()
	}
	return t.sameScore()
}

func (t *Tennis) isScoreDifferent() bool {
	return t.secondPlayerScoreTime != t.firstPlayerScoreTime
}

func (t *Tennis) deuceScore() string {
	return "deuce"
}

func (t *Tennis) sameScore() string {
	return scoreLookup[t.firstPlayerScoreTime] + " all"
}

func (t *Tennis) isDeuce() bool {
	return t.firstPlayerScoreTime >= 3
}

func (t *Tennis) advState() string {
	if t.isAdv() {
		return t.advScore()
	}
	return t.winScore()
}

func (t *Tennis) isAdv() bool {
	return math.Abs(float64(t.firstPlayerScoreTime)-float64(t.secondPlayerScoreTime)) == 1
}

func (t *Tennis) differentScore() string {
	return scoreLookup[t.firstPlayerScoreTime] + " " + scoreLookup[t.secondPlayerScoreTime]
}

func (t *Tennis) advScore() string {
	if t.firstPlayerScoreTime > t.secondPlayerScoreTime {
		return t.firstPlayerName + " adv."
	}
	return t.secondPlayerName + " adv."
}

func (t *Tennis) winScore() string {
	if t.firstPlayerScoreTime > t.secondPlayerScoreTime {
		return t.firstPlayerName + " win."
	}
	return t.secondPlayerName + " win."
}

func (t *Tennis) firstPlayerScore() {
	t.firstPlayerScoreTime++
}

func (t *Tennis) secondPlayerScore() {
	t.secondPlayerScoreTime++
}

func NewTennis(firstPlayerName string, secondPlayerName string) *Tennis {
	return &Tennis{
		firstPlayerName:  firstPlayerName,
		secondPlayerName: secondPlayerName,
	}
}
