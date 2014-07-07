package gseg

import (
	"fmt"
	"sort"
)

//获取最大匹配的备选词
func maxMatch(cache [][]int) [][]int {
	max := 0
	tmp := [][]int{}

	for _, v := range cache {
		c := []int{}
		c = append(c, v...)

		sort.Ints(v)
		if v[2] > max {
			tmp = [][]int{c}
			max = v[2]
		} else if v[2] == max {
			tmp = append(tmp, c)
		}
	}

	return tmp
}

//最大平均词语长度
func largestAvg(cache [][]int) [][]int {
	max := 0
	tmp := [][]int{}

	for _, v := range cache {
		var l, t, a int

		for _, v2 := range v {
			if v2 == 0 {
				continue
			}
			l++
			t = v2 - a + t
			a = v2
		}
		avg := t / l
		if avg > max {
			tmp = [][]int{v}
			max = avg
		} else if avg == max {
			tmp = append(tmp, v)
		}
	}

	return tmp
}

//获取所有可能的3个备选词
func searchWords(s *Seg, w []string) [][]int {
	tmp := [][]int{}
	c1 := s.dict.GetAll(w)
	l := len(w)

	fmt.Println()
	if len(c1) == 0 {
		tmp = append(tmp, []int{0, 0, 0})
		return tmp
	}

	//先循环第一个词的所有可能, 取出第二个词
	for _, v1 := range c1 {
		offset := v1 + 1
		c2 := s.dict.GetAll(w[offset:])
		if len(c2) == 0 {
			c2 = []int{0}
		}
		//循环第二个词的所有可能, 取出第三个词
		for _, v2 := range c2 {
			offset := v1 + v2 + 2
			fmt.Println("offset==>", offset)
			if offset >= l {
				tmp = append(tmp, []int{v1, l, 0})
				continue
			}
			c3 := s.dict.GetAll(w[offset:])
			fmt.Println("c3==>", w[offset:])
			if len(c3) == 0 {
				c3 = []int{0}
			}
			//保存所有的可能
			for _, v3 := range c3 {
				offset := v1 + v2 + 1

				end := offset + v3 + 1
				if end >= l {
					end = 0
				}
				tmp = append(tmp, []int{v1, offset, end})
			}
		}

	}

	return tmp
}
