package sentence_score

import (
	"github.com/cryptography-research-lab/go-sentence-score/data"
	"github.com/golang-infrastructure/go-trie"
	"strings"
)

type Dictionary struct {
	trie *trie.Trie[string]
}

// LoadBuiltinDictionary 加载内建的词典
func LoadBuiltinDictionary() *Dictionary {
	trie := trie.New[string]()

	for _, line := range strings.Split(data.BuiltinDictionary, "\n") {
		_ = trie.Add(line, line)
	}

	return &Dictionary{
		trie: trie,
	}
}

// AddWord 往词典中增加一个单词
func (x *Dictionary) AddWord(word string) error {
	return x.trie.Add(word, word)
}

// RemoveWord 从词典中删除单词
func (x *Dictionary) RemoveWord(word string) error {
	return x.trie.Remove(word)
}

func (x *Dictionary) Exists(word string) bool {
	value, err := x.trie.Query(word)
	if err != nil {
		return false
	}
	return value == word
}
