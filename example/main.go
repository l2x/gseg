package main

import (
	"fmt"
	"github.com/l2x/gseg"
)

func main() {
	gs := gseg.New()
	err := gs.LoadDict("../data/words.dic")
	if err != nil {
		fmt.Println(err)
		return
	}

	s := "那就是说我不回家了"
	//fmt.Println(gs.Simple(s))
	//gs.Complex(s)

	//s = "研究生"
	//s = "研究生命起源"
	s = "这个四个过滤规则中，如果使用simple的匹配方法，只能使用第一个规则过滤，如果使用complex的匹配方法，则四个规则都可以使用。实际使用中，一般都是使用complex的匹配方法＋四个规则过滤。（simple的匹配方法实质上就是正向最大匹配，实际中很少只用这一个方法）"
	fmt.Println(gs.Simple(s))
	fmt.Println(gs.Complex(s))

}
