package main

import (
	"fmt"
	sentence_score "github.com/cryptography-research-lab/go-sentence-score"
)

func main() {

	// 计算单个句子
	s := "1123thisisverybkjkjbsecurityfghfkeyassdfdfadasdasdasdasdasdasdasdgoodasdasdking"
	words, score := sentence_score.CalculateScore(s)
	fmt.Println(words)
	fmt.Println(score)

	// 计算多个句子并排序
	slice := []string{
		"1123thisisverybkjkjbsecurityfghfkeygoodasdasdking",
		"1123thisisverybkjkintenetjbsecurityfghfkeygoodasdasdking",
		"1123thisisverybknamejkjbsecurityfghfkeygoodasdasdking",
	}
	sort := sentence_score.CalculateScoreAndDescSort(slice)
	fmt.Println(sort)

}
