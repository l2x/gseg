package gseg

import (
	"bufio"
	//"fmt"
	"github.com/l2x/gseg/trie"
	"os"
	"strings"
)

//加载字典, 支持加载多个
func (s *Seg) LoadDict(file string) error {
	if s.dict == nil {
		s.dict = trie.New()
	}

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	//读取词典
	var w []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		w = strings.Split(scanner.Text(), "")
		if len(w) > MaxWordLen {
			MaxWordLen = len(w)
		}
		s.dict.Insert(w)
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
