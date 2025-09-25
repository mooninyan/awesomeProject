package main

// Вариант2 :
//
// Напишите функцию isValid, которая принимает строку s, состоящую только из символов '(', ')', '{', '}', '[' и ']',
// и возвращает true, если скобки расставлены правильно, и false в противном случае.
//
// Примеры:
// Вход: "()"
// Выход: true
//
// Вход: "([{}])"
// Выход: true

func main() {
	str1 := "([{}])"
	checkCorrect(str1)
}

func checkCorrect(str1 string) bool {
	m := make(map[string]string)
	m[")"] = "("
	m["}"] = "{"
	m["]"] = "["
	m["["] = "]"
	m["{"] = "}"
	m["("] = ")"

	open := make(map[string]bool)
	open["("] = true
	open["["] = true
	open["{"] = true

	closes := make(map[string]bool)
	closes[")"] = true
	closes["]"] = true
	closes["}"] = true

	stack := Stack{}
	for _, it := range str1 {
		itv := string(it)
		if open[itv] {
			stack.Push(itv)
		}
		if closes[itv] {
			pop := stack.Pop()
			if pop != itv {
				return false
			}
		}
	}
	return true
}

type Stack struct {
	data []string
}

func (s *Stack) Push(v string) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() string {
	if len(s.data) == 0 {
		return ""
	}
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}
