package main

type Stack struct {
	data []byte
	head int
}

func NewStack() Stack {
	return Stack{
		data: []byte{},
		head: -1,
	}
}

func (s *Stack) Push(val byte) {
	s.head++
	s.data = append(s.data, val)
}

func (s *Stack) Top() byte {
	if s.head < 0 {
		return 0
	}
	return s.data[s.head]
}

func (s *Stack) Pop() byte {
	if s.head < 0 {
		return 0
	}
	val := s.data[s.head]
	s.head--
	if s.head < 0 {
		s.data = []byte{}
	}
	s.data = s.data[:s.head]
	return val
}

func isValid(s string) bool {
	inverseParenMap := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	parenMap := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	parenStack := NewStack()
	for _, v := range s {
		char := byte(v)

		if _, ok := parenMap[char]; ok {
			parenStack.Push(char)
		} else if opening, ok := inverseParenMap[char]; ok {
			if parenStack.Top() != opening {
				return false
			} else {
				parenStack.Pop()
			}
		}
	}
	return parenStack.head < 0
}
