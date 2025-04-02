package TennisKata

import "testing"

func TestTennis(t *testing.T) {
	tests := []struct {
		name               string
		firstPlayerScores  int
		secondPlayerScores int
		expectedScore      string
	}{
		{"love all", 0, 0, "love all"},
		{"fifteen love", 1, 0, "fifteen love"},
		{"thirty love", 2, 0, "thirty love"},
		{"forty love", 3, 0, "forty love"},
		{"love fifteen", 0, 1, "love fifteen"},
		{"love thirty", 0, 2, "love thirty"},
		{"fifteen all", 1, 1, "fifteen all"},
		{"thirty all", 2, 2, "thirty all"},
		{"deuce", 3, 3, "deuce"},
		{"first player advances", 4, 3, "Nick adv."},
		{"second player advances", 3, 4, "David adv."},
		{"first player wins", 5, 3, "Nick wins"},
		{"second player wins", 3, 5, "David wins"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tennis := NewTennis("Nick", "David")
			firstPlayerScoreTimes(&tennis, test.firstPlayerScores)
			secondPlayerScoreTimes(&tennis, test.secondPlayerScores)
			scoreShouldBe(t, &tennis, test.expectedScore)
		})
	}
}

func secondPlayerScoreTimes(tennis *Tennis, times int) {
	for i := 0; i < times; i++ {
		tennis.SecondPlayerScore()
	}
}

func firstPlayerScoreTimes(tennis *Tennis, times int) {
	for i := 0; i < times; i++ {
		tennis.FirstPlayerScore()
	}
}

func scoreShouldBe(t *testing.T, tennis *Tennis, want string) {
	t.Helper()
	got := tennis.Score()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
