package stack

type Stack struct {
	data []interface{}
}

func (s * Stack) IsEmpty() bool {
	return len(s.data) == 0
}

func (s * Stack) Pop() {
	if !s.IsEmpty() {
		s.data = s.data[:len(s.data) - 1]
	}
}

func (s * Stack) Push(value interface{}) {
	s.data = append(s.data, value)
}

func (s * Stack) Top() (res interface{}) {
	if !s.IsEmpty() {
		res = s.data[len(s.data) - 1]
	}
	return
}

