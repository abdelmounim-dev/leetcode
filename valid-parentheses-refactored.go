package main

type Stack struct {
    data []byte
}

func NewStack() *Stack {
    return &Stack{
        data: []byte{},
    }
}

func (s *Stack) Push(val byte) {
    s.data = append(s.data, val)
}

func (s *Stack) Top() byte {
    if len(s.data) == 0 {
        return 0
    }
    return s.data[len(s.data)-1]
}

func (s *Stack) Pop() byte {
    if len(s.data) == 0 {
        return 0
    }
    val := s.data[len(s.data)-1]
    s.data = s.data[:len(s.data)-1]
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
    return len(parenStack.data) == 0
}