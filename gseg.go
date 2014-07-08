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

type Cache struct {
	Ind int
	C   []Cache
}

func New() Seg {
	return Seg{}
}

//Simple 匹配方法
func (s *Seg) Simple(words string) []string {
	var w []string = wordsInit(words)
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
func (s *Seg) Complex(words string) []string {
	var w []string = wordsInit(words)

	segment := []string{}
	start := 0
	end := 0
	max := len(w)

	for start < max {
		cache := searchWords(s, w[start:])

		//如果第一个词只有一个, 那就把这个词作为第一个词
		if len(cache) == 1 {
			end = start + cache[0][0] + 1
			segment = append(segment, strings.Join(w[start:end], ""))
			start = end
			continue
		}

		//maximum matching
		cache = maxMatch(cache)

		if len(cache) == 1 {
			end = start + cache[0][0] + 1
			segment = append(segment, strings.Join(w[start:end], ""))
			start = end

			continue
		}

		//largest average word length
		cache = largestAvg(cache)
		if len(cache) == 1 {
			end = start + cache[0][0] + 1
			segment = append(segment, strings.Join(w[start:end], ""))
			start = end

			continue
		}

		//smallest variance of word lengths
		cache = smallestVariance(cache)
		if len(cache) == 1 {
			end = start + cache[0][0] + 1
			segment = append(segment, strings.Join(w[start:end], ""))
			start = end

			continue
		}

		//TODO
		//largest sum of degree of morphemic freedom of one-character words
		end = start + cache[0][0] + 1
		segment = append(segment, strings.Join(w[start:end], ""))
		start = end

	}

	return segment
}
