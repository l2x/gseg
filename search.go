package gseg

import (
	//"fmt"
	"sort"
)

func maxMatch(cache [][]int) [][]int {
	max := 0
	tmp := [][]int{}

	for _, v := range cache {
		sort.Ints(v)
		if v[2] > max {
			tmp = [][]int{v}
		} else if v[2] == max {
			tmp = append(tmp, v)
		}
	}

	return tmp
}

func searchWords(s *Seg, w []string) [][]int {
	tmp := [][]int{}
	c1 := s.dict.GetAll(w)
	l := len(w)

	if len(c1) == 0 {
		tmp = append(tmp, []int{0, 0, 0})
		return tmp
	}

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
