package sentence_score

import (
	variable_parameter "github.com/golang-infrastructure/go-variable-parameter"
	"strings"
)

// BuiltinDictionary 内建的词典
var BuiltinDictionary = LoadBuiltinDictionary()

// CalculateScore 为字符串计算得分
func CalculateScore(sentence string, dictionary ...*Dictionary) float64 {

	// 如果没有传递字典的话，则使用默认的字典
	dictionary = variable_parameter.SetDefaultParam(dictionary, BuiltinDictionary)

	// 所有字符统一转为小写
	sentence = strings.ToLower(sentence)

	// 包含的单词的数量
	words := make([]string, 0)
	// 非单词的部分
	fuckParts := make([]string, 0)
	for sentenceIndex := 0; sentenceIndex < len(sentence); {

		exploreIndex := sentenceIndex
		matchIndex := sentenceIndex
		node := dictionary[0].trie.RootTrieNode()
		for exploreIndex < len(sentence) {

			// TODO 允许更多地白名单字符串
			// 对字符检查，如果不是英文字母，则不需要再往后继续
			if sentence[exploreIndex] < 'a' || sentence[exploreIndex] > 'z' {
				break
			}

			// 一层一层的深入下去匹配
			childNode, exists := node.Children[string(sentence[exploreIndex])]
			// 节点不存在，说明本次匹配已经到头了，没必要再继续深入下去了
			if !exists {
				break
			}
			// 这个字符存在，但是字符有绑定到单词吗
			if childNode.Value != "" {
				matchIndex = exploreIndex + 1
			}
			// 匹配到了，继续往下深入
			exploreIndex++
			node = childNode
		}

		// 啥也没匹配到，往后加1
		if matchIndex == sentenceIndex {
			fuckParts = append(fuckParts, sentence[sentenceIndex:sentenceIndex+1])
			sentenceIndex++
			continue
		}

		// 匹配到了一些东西，保留并跳过
		words = append(words, sentence[sentenceIndex:matchIndex])
		sentenceIndex = matchIndex
	}

	// 读取单词，采取最长匹配原则
	return float64(len(words))
}
