package trie

type Trie struct {
	Value    string
	Count    int
	End      int
	Children map[string]*Trie
}

func New() *Trie {
	return &Trie{
		Children: make(map[string]*Trie),
	}
}

func (t *Trie) Insert(w []string) error {
	l := len(w)
	if l == 0 {
		return nil
	}

	for _, v := range w {
		if v == " " {
			continue
		}
		t = t.insert(v)
	}
	t.last()

	return nil
}

func (t *Trie) insert(w string) *Trie {
	if _, ok := t.Children[w]; !ok {
		t.Children[w] = New()
		t.Children[w].Value = w
	}

	t.Count++
	return t.Children[w]
}

func (t *Trie) Get(w []string) (*Trie, bool) {
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

func (t *Trie) get(w string) (*Trie, bool) {
	if _, ok := t.Children[w]; ok {
		return t.Children[w], true
	}

	return t, false
}

func (t *Trie) GetMax(w []string) (*Trie, int) {
	var ok bool

	arr := []int{0}
	for k, v := range w {
		t, ok = t.get(v)
		if !ok {
			break
		}
		arr = append(arr, k)
	}

	return t, arr[len(arr)-1] + 1
}

func (t *Trie) GetAll(w []string) []int {
	var ok bool

	arr := []int{}
	for k, v := range w {
		t, ok = t.get(v)
		if !ok {
			break
		}
		arr = append(arr, k)
	}
	return arr
}

func (t *Trie) last() {
	t.End++
}
