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

type Cache struct {
	Ind int
	C   []Cache
}

func New() Seg {
	return Seg{}
}

//Simple 匹配方法
func (s *Seg) Simple(words string) []string {
	var w []string = strings.Split(words, "")
	var res []string = []string{}
	var start, end, i, max int
	max = len(w)

	for start < max {
		_, i = s.dict.GetMax(w[start:])
		end = start + i
		res = append(res, strings.Join(w[start:end], ""))
		start = end
	}

	return res
}

//complex 匹配方法
func (s *Seg) Complex(words string) {
	var w []string = strings.Split(words, "")

	//分别取出3组备选词
	cache := s.search(w, 0)

	for k, v := range cache {
		offset := v.Ind + 1
		c2 := s.search(w[offset:], offset)
		cache[k].C = c2

		for k2, v2 := range c2 {
			offset := v2.Ind + 1
			c3 := s.search(w[offset:], offset)
			cache[k].C[k2].C = c3
		}
	}
	//TODO
	//maximum matching

	//largest average word length

	//smallest variance of word lengths

	//largest sum of degree of morphemic freedom of one-character words

}

func (s *Seg) search(w []string, offset int) []Cache {
	cache := []Cache{}
	res := s.dict.GetAll(w)
	for _, v := range res {
		c := Cache{
			Ind: v + offset,
			C:   []Cache{},
		}
		cache = append(cache, c)
	}

	return cache
}
