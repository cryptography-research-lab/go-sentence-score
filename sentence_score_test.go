package sentence_score

import (
	"fmt"
	"testing"
)

func TestCalculateScore(t *testing.T) {
	s := "1123thisisverybkjkjbsecurityfghfkeyassdfdfadasdasdasdasdasdasdasdgoodasdasdking"
	words, score := CalculateScore(s)
	fmt.Println(words)
	fmt.Println(score)
}

func TestCalculateScoreAndDescSort(t *testing.T) {
	slice := []string{
		"1123thisisverybkjkjbsecurityfghfkeygoodasdasdking",
		"1123thisisverybkjkintenetjbsecurityfghfkeygoodasdasdking",
		"1123thisisverybknamejkjbsecurityfghfkeygoodasdasdking",
	}
	sort := CalculateScoreAndDescSort(slice)
	fmt.Println(sort)
}
