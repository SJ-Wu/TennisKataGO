package TennisKataGO

import "testing"

func TestTennis(t *testing.T) {
	tests := []struct {
		testName           string
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
		{"first player advances.", 4, 3, "Nick adv."},
		{"second player advances.", 3, 4, "David adv."},
		{"second player wins.", 3, 5, "David win."},
		{"first player wins.", 5, 3, "Nick win."},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			tennis := NewTennis("Nick", "David")
			firstPlayerScoreTime(tennis, test.firstPlayerScores)
			secondPlayerScoreTime(tennis, test.secondPlayerScores)
			scoreShouldBe(t, tennis.Score(), test.expectedScore)
		})
	}
}

func firstPlayerScoreTime(tennis *Tennis, time int) {
	for i := 0; i < time; i++ {
		tennis.firstPlayerScore()
	}
}

func secondPlayerScoreTime(tennis *Tennis, time int) {
	for i := 0; i < time; i++ {
		tennis.secondPlayerScore()
	}
}

func scoreShouldBe(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s, want %s", got, want)
	}
}
