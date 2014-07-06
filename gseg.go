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

	cache := s.search(w)
	fmt.Println(cache)

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

	//TODO
	//maximum matching

	//largest average word length

	//smallest variance of word lengths

	//largest sum of degree of morphemic freedom of one-character words

}

func (s *Seg) search(w []string) [][]int {
	tmp := [][]int{}
	c1 := s.dict.GetAll(w)
	l := len(w)

	for _, v1 := range c1 {
		offset := v1 + 1
		c2 := s.dict.GetAll(w[offset:])
		if len(c2) == 0 {
			end := offset
			if end == l {
				end = 0
			}
			tmp = append(tmp, []int{v1, end, 0})
			continue
		}
		for _, v2 := range c2 {
			offset := v1 + v2 + 2
			c3 := s.dict.GetAll(w[offset:])
			if len(c3) == 0 {
				end := offset
				if end == l {
					end = 0
				}
				tmp = append(tmp, []int{v1, offset - 1, end})
				continue
			}
			for _, v3 := range c3 {
				offset := v1 + v2 + 1
				end := offset + v3 + 1
				if end > l {
					end = l
				}
				tmp = append(tmp, []int{v1, offset, end})
			}
		}

	}

	return tmp
}

func maxMatch(cache []Cache) ([]Cache, bool) {

	return cache, false
}
