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
	cache := []Cache{}
	res := s.dict.GetAll(w)
	for _, v := range res {
		c := Cache{
			Ind: v,
			C:   []Cache{},
		}
		cache = append(cache, c)
	}

	fmt.Println(cache)

	for k, cv := range cache {
		fmt.Println(w[:cv.Ind])

		i := cv.Ind + 1

		res = s.dict.GetAll(w[i:])
		for _, v := range res {
			c := Cache{
				Ind: v + i,
				C:   []Cache{},
			}
			fmt.Println(w[cv.Ind:c.Ind])
			cache[k].C = append(cache[k].C, c)
		}
		fmt.Println("---")
	}

	fmt.Println(cache)

}

func (s *Seg) search(start int, end int, w []string) int {
	var ok bool
	_, ok = s.dict.Get(w)
	if ok {
		return end
	}
	return 0
}
