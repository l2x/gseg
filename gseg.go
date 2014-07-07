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

	segment := []string{}
	start := 0
	end := 0
	max := len(w)

	for start < max {
		cache := searchWords(s, w[start:])

		if len(cache) == 1 {
			end = start + cache[0][0] + 1
			segment = append(segment, strings.Join(w[start:end], ""))
			start = end
		}

		fmt.Println("segment=>", segment)

		/*
			fmt.Println(cache)
			for _, v := range cache {
				fmt.Println(w[:v[0]+1])
				if v[1] != 0 {
					end := v[1] + 1
					if v[1]+1 > len(w) {
						end = len(w)
					}
					fmt.Println(w[v[0]+1 : end])
				}
				if v[2] != 0 {
					end := v[2] + 1
					if v[2]+1 > len(w) {
						end = len(w)
					}
					fmt.Println(w[v[1]+1 : end])
				}
				fmt.Println("+++++")
			}
		*/

		fmt.Println("cache=>", cache)

		//TODO
		//maximum matching
		cache = maxMatch(cache)
		if len(cache) == 1 {
			end = start + cache[0][0] + 1
			segment = append(segment, strings.Join(w[start:end], ""))
			start = end
		}
		fmt.Println("cache=>", cache)

		//largest average word length

		//smallest variance of word lengths

		//largest sum of degree of morphemic freedom of one-character words

		start = max
	}

}
