package main

import (
	"fmt"
	"flag"

	"github.com/anasteizha/VK_GOLANG/calc"
)

func main() {
	flag.Parse()
	expression := flag.Arg(0)

	postfixExpression, err := calculator.ToPostfix(expression)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := calculator.Calculate(postfixExpression)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

