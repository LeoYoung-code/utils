package utils

type SetStruct map[string]struct{}

func (s SetStruct) Has(key string) bool {
	_, ok := s[key]
	return ok
}

func (s SetStruct) Add(key string) {
	s[key] = struct{}{}
}

func (s SetStruct) Delete(key string) {
	delete(s, key)
}
