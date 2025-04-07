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
