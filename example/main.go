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
	//gs.Complex2(s)
	gs.Complex(s)
}
