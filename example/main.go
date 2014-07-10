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

	s := "直到现在，整个中国科技生态都还无法完全摆脱“山寨”的刻板印象。而作为中国新一代创业公司的代表，既然心怀国际化的志向，就应该学会尊重国际上的玩法。小米本已经拿了手好牌，但如果摆脱不了抄袭的惯性，那只会重走上一代山寨公司的老路。从野蛮到文明的进化固然痛苦，但是如果衣服一旦脱下来，就不是那么好穿上了。"
	//fmt.Println(gs.Simple(s))
	fmt.Println(gs.Complex(s))

}
