package sentence_score

import (
	"fmt"
	"testing"
)

func TestLoadBuiltinDictionary(t *testing.T) {
	dictionary := LoadBuiltinDictionary()
	fmt.Println(dictionary)
}
