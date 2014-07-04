package gseg

import (
	//	"fmt"
	"github.com/l2x/gseg/trie"
	"strings"
)

var (
	MaxWordLen = 9
)

type Seg struct {
	dict *trie.Trie
}

func New() Seg {
	return Seg{}
}

func (s *Seg) Simple(words string) []string {
	w := strings.Split(words, "")
	res := []string{}
	start := 0
	end := 0
	max := len(w)

	for start < max {
		_, i := s.dict.GetMax(w[start:])
		end = start + i
		res = append(res, strings.Join(w[start:end], ""))
		start = end
	}

	return res
}

func (s *Seg) Complex(words string) {

}

func (s *Seg) search(start int, end int, w []string) int {
	var ok bool
	_, ok = s.dict.Get(w)
	if ok {
		return end
	}
	return 0
}
