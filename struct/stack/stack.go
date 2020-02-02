package stack

import "fmt"

type stack struct {
	limit int
	s     []int
	top   int
}

func NewStack(l int) *stack {
	st := new(stack)
	st.limit = l
	st.top = 0
	st.s = make([]int, st.limit)
	return st
}
func (s *stack) Push(v int) bool {
	if s.top == s.limit {
		return false
	} else {
		s.s[s.top] = v
		s.top++
		return true
	}
}
func (s *stack) Pop() (int, bool) {
	if s.top == 0 {
		return 0, false
	} else {
		s.top--
		return s.s[s.top], true
	}
}
func (s *stack) String() (str string) {
	for i, val := range s.s {
		if i == s.top {
			break
		} else {
			str += fmt.Sprintf("[%v:%v] ", i, val)
		}
	}
	return
}
