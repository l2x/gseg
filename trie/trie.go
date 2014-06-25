package trie

type Trie struct {
	Value    byte
	Count    int
	End      int
	Children map[byte]*Trie
}

func New() *Trie {
	return &Trie{
		Children: make(map[byte]*Trie),
	}
}

func (t *Trie) Insert(w []byte) error {
	l := len(w)
	if l == 0 {
		return nil
	}

	for _, v := range w {
		t = t.insert(v)
	}
	t.last()

	return nil
}

func (t *Trie) insert(w byte) *Trie {
	if _, ok := t.Children[w]; !ok {
		t.Children[w] = New()
		t.Children[w].Value = w
	}

	t.Count++
	return t.Children[w]
}

func (t *Trie) Get(w []byte) (*Trie, bool) {
	var ok bool

	for _, v := range w {
		t, ok = t.get(v)
		if !ok {
			return t, false
		}
	}

	if t.End == 0 {
		return t, false
	}

	return t, true
}

func (t *Trie) get(w byte) (*Trie, bool) {
	if _, ok := t.Children[w]; ok {
		return t.Children[w], true
	}

	return t, false
}

func (t *Trie) last() {
	t.End++
}
