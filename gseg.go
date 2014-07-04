package gseg

import (
	"fmt"
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

func (s *Seg) Simple(words string) []int {
	w := strings.Split(words, "")

	max := len(w)
	start := 0
	ind := 0
	res := []int{}

	for start < max {
		reserve := []int{}

		for end := start + 1; (end-start) <= MaxWordLen && end < max; end++ {
			if start+end >= max {
				end = max
			}

			ind = s.search(start, end, w[start:end])

			fmt.Println(start, end, ind)

			if ind != 0 {
				reserve = append(reserve, ind)
			}
		}

		//没有匹配
		fmt.Println(reserve, start)
		lr := len(reserve)
		if lr == 0 {
			reserve = append(reserve, start+1)
		}

		start = reserve[len(reserve)-1]

		res = append(res, start)
	}

	fmt.Println(res)

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
