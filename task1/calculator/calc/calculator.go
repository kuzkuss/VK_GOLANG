package calculator

import (
	"errors"
	"math"
	"strconv"
	"unicode"

	"github.com/kuzkuss/VK_GOLANG/queue"
	"github.com/kuzkuss/VK_GOLANG/stack"
)

func ToPostfix(expression string) (queue.Queue, error) {
	var err error
	q := queue.Queue{}
	s := stack.Stack{}
	numRes := ""
	for _, el := range expression {
		if unicode.IsDigit(el) {
			numRes += string(el)
		} else {
			if numRes != "" {
				q.Push(numRes)
			}
			switch {
			case el == '+' || el == '-' || el == '*' || el == '/':
				if s.IsEmpty() || s.Top() == '(' || getPriority(el) >= getPriority(s.Top().(rune)) {
					s.Push(el)
				} else if getPriority(el) < getPriority(s.Top().(rune)) {
					for !s.IsEmpty() && (s.Top() != '(' || getPriority(el) <= getPriority(s.Top().(rune))) {
						q.Push(s.Top())
						s.Pop()
					}
					s.Push(el)
				}
			case el == '(':
				s.Push(el)
			case el == ')':
				for s.Top() != '(' {
					q.Push(s.Top())
					s.Pop()
				}
				s.Pop()
			default:
				err = errors.New("incorrect expression")
				return q, err
			}
			numRes = ""
		}
	}

	if numRes != "" {
		q.Push(numRes)
	}

	for !s.IsEmpty() {
		q.Push(s.Top())
		s.Pop()
	}

	return q, err
}

func Calculate(expression string) (resValue float64, err error) {
	postfixExpression, err := ToPostfix(expression)
	if err != nil {
		return
	}

	s := stack.Stack{}
	for !postfixExpression.IsEmpty() && err == nil {
		el := postfixExpression.Front()
		postfixExpression.Pop()
		if strEl, ok := el.(string); ok {
			var num float64
			if num, err = strconv.ParseFloat(strEl, 64); err == nil {
				s.Push(num)
			}
		} else if el == '+' || el == '-' || el == '*' || el == '/' {
			rightOperand, ok := s.Top().(float64)
			if !ok {
				err = errors.New("incorrect postfix expression")
				continue
			}
			s.Pop()
			leftOperand, ok := s.Top().(float64)
			if !ok {
				err = errors.New("incorrect postfix expression")
				continue
			}
			s.Pop()
			var res float64
			if res, err = calcOperation(el.(rune), leftOperand, rightOperand); err == nil {
				s.Push(res)
			}
		} else {
			err = errors.New("incorrect postfix expression")
		}
	}

	if err == nil {
		resValue = s.Top().(float64)
	}
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

