package TennisKata

type Tennis struct {
	firstPlayerName  string
	secondPlayerName string
	firstPlayer      int
	secondPlayer     int
}

var scoreLookup = map[int]string{
	0: "love",
	1: "fifteen",
	2: "thirty",
	3: "forty",
}

func NewTennis(firstPlayerName, secondPlayerName string) Tennis {
	return Tennis{
		firstPlayerName:  firstPlayerName,
		secondPlayerName: secondPlayerName,
		firstPlayer:      0,
		secondPlayer:     0,
	}
}

func (t *Tennis) Score() string {
	if t.isScoreDifferent() {
		if t.isAdv() {
			return t.advState()
		}
		return t.lookupScore()
	}
	if t.isDeuce() {
		return t.deuceScore()
	}
	return t.sameScore()
}

func (t *Tennis) advState() string {
	if t.isWin() {
		return t.winScore()
	}
	return t.advScore()
}

func (t *Tennis) isDeuce() bool {
	return t.firstPlayer >= 3
}

func (t *Tennis) isWin() bool {
	return abs(t.firstPlayer-t.secondPlayer) > 1
}

func (t *Tennis) isAdv() bool {
	return t.firstPlayer > 3 || t.secondPlayer > 3
}

func (t *Tennis) isScoreDifferent() bool {
	return t.firstPlayer != t.secondPlayer
}

func (t *Tennis) sameScore() string {
	return scoreLookup[t.firstPlayer] + " all"
}

func (t *Tennis) deuceScore() string {
	return "deuce"
}

func (t *Tennis) lookupScore() string {
	return scoreLookup[t.firstPlayer] + " " + scoreLookup[t.secondPlayer]
}

func (t *Tennis) winScore() string {
	if t.firstPlayer > t.secondPlayer {
		return t.firstPlayerName + " wins"
	}
	return t.secondPlayerName + " wins"
}

func (t *Tennis) advScore() string {
	if t.firstPlayer > t.secondPlayer {
		return t.firstPlayerName + " adv."
	}
	return t.secondPlayerName + " adv."
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func (t *Tennis) FirstPlayerScore() {
	t.firstPlayer++
}

func (t *Tennis) SecondPlayerScore() {
	t.secondPlayer++

}
