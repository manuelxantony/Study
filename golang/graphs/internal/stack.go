package internal

type Stack []string

// add to end
func (s *Stack) Push(val ...string) {
	*s = append(*s, val...)
}

// remove from end
func (s *Stack) Pop() string {
	temp := *s
	l := len(temp) - 1
	poped := temp[l]
	*s = temp[:l]
	return poped
}

// true if empty
func (s Stack) IsEmpty() bool {
	return len(s) == 0
}
