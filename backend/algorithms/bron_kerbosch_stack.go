package algorithms

type Stack []StackEntry
type StackEntry []Clique

func NewStack() Stack {
	s := make([]StackEntry, 0)
	return s
}

func (s Stack) Push(a Clique, b Clique, c Clique) Stack {
	entry := make([]Clique, 0)
	entry = append(entry, a)
	entry = append(entry, b)
	entry = append(entry, c)
	s = append(s, entry)
	return s
}

func (s Stack) Pop() (Stack, StackEntry) {
	if len(s) == 0 {
		return nil, nil
	} else {
		entry := s[len(s)-1]
		s = s[:len(s)-1]
		return s, entry
	}
}

func (s Stack) IsEmpty() bool {
	if len(s) == 0 {
		return true
	} else {
		return false
	}
}
