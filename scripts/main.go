package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode/utf8"
)

// LoadBuiltinDictionaryAsBuilder 用来生成默认的词典
// 字典来自： https://github.com/mahavivo/english-wordlists
func LoadBuiltinDictionaryAsBuilder() {
	fileSlice, err := ioutil.ReadDir("scripts/temp_data")
	if err != nil {
		fmt.Println("read dir fail:", err)
		panic(err)
	}
	wordSet := make(map[string]struct{}, 0)
	for _, file := range fileSlice {
		fileBytes, err := os.ReadFile("scripts/temp_data/" + file.Name())
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
		split := strings.Split(string(fileBytes), "\n")
		for _, line := range split {
			// 兼容所有词典，处理出真正的单词
			if strings.HasPrefix(line, "*") {
				line = strings.ReplaceAll(line, "*", "")
			}
			if strings.Contains(line, "[") {
				line = strings.SplitN(line, "[", 2)[0]
			}
			if strings.Contains(line, ",") {
				line = strings.SplitN(line, ",", 2)[0]
			}
			if len(line) != utf8.RuneCountInString(line) {
				continue
			}
			line = strings.TrimSpace(line)
			if len(line) < 2 {
				continue
			}
			line = strings.ToLower(line)
			wordSet[line] = struct{}{}
		}
	}

	s := bytes.Buffer{}
	for word := range wordSet {
		s.WriteString(word)
		s.WriteString("\n")
	}

	os.WriteFile("data/builtin-dictionary.txt", s.Bytes(), os.ModePerm)
}

func main() {
	LoadBuiltinDictionaryAsBuilder()
}
