package main

import "errors"

type Stack struct {
	Elem []rune
}

func (s *Stack) add(elm ...rune) {
	s.Elem = append(s.Elem, elm...)
}

func (s *Stack) pop() {
	if s.isEmpty() {
		return
	}
	s.Elem = s.Elem[:len(s.Elem)-1]
}

func (s *Stack) top() (rune, error) {
	if s.isEmpty() {
		return 0, errors.New("is nil")
	}
	return s.Elem[len(s.Elem)-1], nil
}

func (s *Stack) isEmpty() bool {
	if len(s.Elem) == 0 {
		return true
	}
	return false
}

func isValid(s string) bool {
	st := Stack{}
	for _, r := range s {
		switch r {
		case '(':
			st.add(r)
		case '[':
			st.add(r)
		case '{':
			st.add(r)
		case ')':
			v, err := st.top()
			if err != nil {
				return false
			}
			if v != '(' {
				return false
			}
			st.pop()
		case ']':
			v, err := st.top()
			if err != nil {
				return false
			}
			if v != '[' {
				return false
			}
			st.pop()
		case '}':
			v, err := st.top()
			if err != nil {
				return false
			}
			if v != '{' {
				return false
			}
			st.pop()
		default:
			return false
		}
	}
	return st.isEmpty()
}
