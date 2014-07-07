package gseg

import (
	//"fmt"
	"math"
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
	var max float64
	tmp := [][]int{}

	for _, v := range cache {
		avg := getAvg(v)
		if avg > max {
			tmp = [][]int{v}
			max = avg
		} else if avg == max {
			tmp = append(tmp, v)
		}
	}

	return tmp
}

//最小变化频率
func smallestVariance(cache [][]int) [][]int {
	var max float64
	tmp := [][]int{}

	for k, v := range cache {
		var a, sd float64
		avg := getAvg(v)

		for _, v2 := range v {
			sd += math.Pow(float64(v2)-a-avg, 2)
			a = float64(v2)
		}

		s := math.Sqrt(sd / float64(len(v)))
		if k == 0 {
			max = s
		}

		if s < max {
			tmp = [][]int{v}
			max = s
		} else if s == max {
			tmp = append(tmp, v)
		}
	}

	return tmp
}

func getAvg(cache []int) float64 {
	var l, t, a int
	var avg float64

	for _, v := range cache {
		if v == 0 {
			continue
		}
		l++
		t = v - a + t
		a = v
	}
	avg = float64(t) / float64(l)

	return avg
}

//获取所有可能的3个备选词
func searchWords(s *Seg, w []string) [][]int {
	tmp := [][]int{}
	c1 := s.dict.GetAll(w)
	l := len(w)

	if len(c1) == 0 {
		tmp = append(tmp, []int{0, 0, 0})
		return tmp
	}

	//先循环第一个词的所有可能, 取出第二个词
	for _, v1 := range c1 {
		offset := v1 + 1
		if v1 == l-1 {
			tmp = append(tmp, []int{v1, 0, 0})
			continue
		}
		c2 := s.dict.GetAll(w[offset:])
		if len(c2) == 0 {
			c2 = []int{0}
		}
		//循环第二个词的所有可能, 取出第三个词
		for _, v2 := range c2 {
			offset := v1 + v2 + 2
			if offset >= l {
				tmp = append(tmp, []int{v1, l - 1, 0})
				continue
			}
			c3 := s.dict.GetAll(w[offset:])
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
