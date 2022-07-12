package calculator

import (
	"errors"
	"math"
	"strconv"
	"unicode"

	"github.com/anasteizha/VK_GOLANG/queue"
	"github.com/anasteizha/VK_GOLANG/stack"
)

func ToPostfix(expression string) (*queue.Queue, error) {
	var err error
	q := new(queue.Queue)
	s := new(stack.Stack)
	for _, el := range expression {
		if unicode.IsDigit(el) {
			q.Push(el)
		} else if el == '+' || el == '-' || el == '*' || el == '/' {
			if s.IsEmpty() || s.Top() == '(' || getPriority(el) >= getPriority(s.Top().(rune)) {
				s.Push(el)
			} else if getPriority(el) < getPriority(s.Top().(rune)) {
				for !s.IsEmpty() && (getPriority(el) <= getPriority(s.Top().(rune)) || s.Top() != '(') {
					q.Push(s.Top())
					s.Pop()
				}
				s.Push(el)
			}
		} else if el == '(' {
			s.Push(el)
		} else if el == ')' {
			for s.Top() != '(' {
				q.Push(s.Top())
				s.Pop()
			}
			s.Pop()
		} else {
			err = errors.New("incorrect expression")
			return q, err
		}
	}

	for !s.IsEmpty() {
		q.Push(s.Top())
		s.Pop()
	}

	return q, err
}

func Calculate(expr *queue.Queue) (resValue float64, err error) {
	s := new(stack.Stack)
	for !expr.IsEmpty() && err == nil {
		el := expr.Front()
		expr.Pop()
		if unicode.IsDigit(el.(rune)) {
			var digit float64
			digit, err = strconv.ParseFloat(string(el.(rune)), 64)
			s.Push(digit)
		} else if el == '+' || el == '-' || el == '*' || el == '/' {
			rightOperand := s.Top()
			s.Pop()
			leftOperand := s.Top()
			s.Pop()
			var res float64
			res, err = calcOperation(el.(rune), leftOperand.(float64), rightOperand.(float64))
			s.Push(res)
		} else {
			err = errors.New("incorrect postfix expression")
		}
	}

	resValue = s.Top().(float64)
	return
}

func calcOperation(op rune, left, right float64) (float64, error) {
	const EPS = 1e-07
	switch op {
	case '+':
		return left + right, nil
	case '-':
		return left - right, nil
	case '*':
		return left * right, nil
	case '/':
		if math.Abs(right) < EPS {
			return 0, errors.New("division by zero")
		}
		return left / right, nil
	}

	return 0, errors.New("incorrect operation")
}

func getPriority(op rune) int {
	const (
		PRIORITY_ADD_DIFF = 1
		PRIORITY_MUL_DIV = 2
		DEFAULT_PRIORITY = -1
	)
	switch op {
	case '+':
		fallthrough
	case '-':
		return PRIORITY_ADD_DIFF
	case '*':
		fallthrough
	case '/':
		return PRIORITY_MUL_DIV
	}
	return DEFAULT_PRIORITY
}

