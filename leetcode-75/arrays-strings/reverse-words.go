package main

import "strings"

type Stack []string

func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

func (s *Stack) Pop() string {
	index := len(*s) - 1
	str := (*s)[index]
	*s = (*s)[:index]
	return str
}

func (s *Stack) Peek() string {
	return (*s)[len(*s)-1]
}

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func reverseWords(s string) string {
	whitespace := map[rune]bool{
		' ':  true, // space
		'\t': true, // tab
		'\n': true, // newline (line feed)
		'\r': true, // carriage return
		'\f': true, // form feed
		'\v': true, // vertical tab
	}

	var words Stack
	var word strings.Builder

	for _, v := range s {
		if whitespace[v] {
			if word.Len() > 0 {
				words.Push(word.String())
				word.Reset()
			}
		} else {
			word.WriteRune(v)
		}
	}

	if word.Len() > 0 {
		words.Push(word.String())
	}

	if words.IsEmpty() {
		return ""
	}

	reversed := words.Pop()

	for !words.IsEmpty() {
		reversed += " " + words.Pop()
	}

	return reversed
}

