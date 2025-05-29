package ExtremeSpeed

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTennis(t *testing.T) {

	tests := []struct {
		message           string
		score             string
		firstPlayerScore  int
		secondPlayerScore int
	}{
		{score: "love all", firstPlayerScore: 0, message: "初始分數為love all"},
		{score: "fifteen love", firstPlayerScore: 1, message: "得壹分，15:0"},
		{score: "thirty love", firstPlayerScore: 2, message: "得兩分，30:0"},
		{score: "forty love", firstPlayerScore: 3, message: "得三分，40:0"},
		{score: "love fifteen", firstPlayerScore: 0, secondPlayerScore: 1, message: "對手得分，0:15"},
		{score: "love thirty", firstPlayerScore: 0, secondPlayerScore: 2, message: "對手得分，0:30"},
		{score: "love forty", firstPlayerScore: 0, secondPlayerScore: 3, message: "對手得分，0:40"},
		{score: "fifteen all", firstPlayerScore: 1, secondPlayerScore: 1, message: "雙方平分，15:15"},
		{score: "thirty all", firstPlayerScore: 2, secondPlayerScore: 2, message: "雙方平分，30:30"},
		{score: "deuce", firstPlayerScore: 3, secondPlayerScore: 3, message: "Deuce"},
		{score: "Nick adv.", firstPlayerScore: 4, secondPlayerScore: 3, message: "一位領先"},
		{score: "David adv.", firstPlayerScore: 3, secondPlayerScore: 4, message: "二位領先"},
		{score: "David win.", firstPlayerScore: 3, secondPlayerScore: 5, message: "二位獲勝"},
		{score: "Nick win.", firstPlayerScore: 5, secondPlayerScore: 3, message: "一位獲勝"},
	}
	for _, test := range tests {
		t.Run(test.message, func(t *testing.T) {
			tennis := NewTennis("Nick", "David")
			giveFirstPlayerScore(tennis, test.firstPlayerScore)
			giveSecondPlayerScore(tennis, test.secondPlayerScore)
			scoreShouldBe(t, test.score, tennis.Score(), test.message)
		})
	}

}

func giveSecondPlayerScore(tennis *Tennis, score int) {
	for i := 0; i < score; i++ {
		tennis.SecondPlayerScore()
	}
}

func scoreShouldBe(t *testing.T, expected string, actual string, message string) {
	assert.Equal(t, expected, actual, message)
}

func giveFirstPlayerScore(tennis *Tennis, time int) {
	for i := 0; i < time; i++ {
		tennis.FirstPlayerScore()
	}
}
