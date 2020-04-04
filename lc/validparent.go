package lc

type Stack struct {
	store []string
	head  int
	tail  int
	size  int
}

func MakeStack() Stack {
	var stack Stack
	stack.store = make([]string, 10)
	stack.head = 0
	stack.tail = 0
	stack.size = 0
	return stack
}

func (stack *Stack) put(p string) {
	stack.store = append(stack.store, p)
	stack.head = len(stack.store)
	stack.size++
}

func (stack *Stack) pop() (string, bool) {
	if stack.size == 0 {
		return "", false
	}
	elem := stack.store[stack.head]
	stack.store = stack.store[:stack.head]
	return elem, true
}
func isOpeningSymbol(char rune) bool {
	if char == '{' || char == '[' || char == '(' {
		return true
	}
	return false
}

func isClosingSymbol(char rune) bool {
	if char == '}' || char == ']' || char == ')' {
		return true
	}
	return false
}

func IsValid(s string) bool {
	stack := MakeStack()
	for _, ch := range s {
		if isOpeningSymbol(ch) == true {
			stack.put(string(ch))
		} else if isClosingSymbol(ch) {
			if len(stack.store) == 0 {
				return false
			}
			if stack.store[stack.head] != string(ch) {
				return false
			}
			stack.pop()
		}
	}
	if len(stack.store) == 0 {
		return true
	} else {
		return false
	}
}
