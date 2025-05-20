package set

type MapSet map[string]struct{}

func (s MapSet) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s MapSet) Add(key string) {
	s[key] = struct{}{}
}

func (s MapSet) Delete(key string) {
	delete(s, key)
}

func (s MapSet) Size() int {
	return len(s)
}

func (s MapSet) Clear() {
	for k := range s {
		delete(s, k)
	}
}

func (s MapSet) Values() []string {
	keys := make([]string, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}
	return keys
}

func (s MapSet) Each(f func(string)) {
	for k := range s {
		f(k)
	}
}
