package finder

import "fmt"

// TOD0: Add docstrings
// TOD0: Add errorhandling
// TOD0: Add logging

type Stack struct {
	content []Folder
	count   int
}

func (s *Stack) Push(element Folder) {
	s.content = append(s.content, element)
	s.count = len(s.content)
}

func (s *Stack) IsEmpty() bool {
	return s.count == 0
}

func (s *Stack) Pop() (Folder, error) {

	if s.IsEmpty() {
		return Folder{}, fmt.Errorf("[ERROR]: Stack is empty. Nothing to pop.")
	}

	// get last element
	el := s.content[s.count-1]

	if s.count > 0 {
		s.content = s.content[:s.count-1]
		s.count = len(s.content)

	} else {
		s.content = []Folder{}
	}
	return el, nil
}
