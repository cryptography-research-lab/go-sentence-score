# 句子打分 

# 一、概述

用来区分一个句子中含有的有效的单词的数量，其实就是一个简单的分词，用于在一些暴力枚举破解字典的场景下把尽可能真实的解找出来。

# 二、安装

```bash
go get -u github.com/cryptography-research-lab/go-sentence-score
```

# 三、示例代码

```go
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
```





