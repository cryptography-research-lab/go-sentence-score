package sentence_score

import (
	"fmt"
	"testing"
)

func TestCalculateScore(t *testing.T) {
	s := "1123thisisverybkjkjbsecurityfghfkeyassdfdfadasdasdasdasdasdasdasdgoodasdasdking"
	score := CalculateScore(s)
	fmt.Println(score)
}
