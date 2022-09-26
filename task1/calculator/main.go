package main

import (
	"flag"
	"fmt"

	"github.com/kuzkuss/VK_GOLANG/calc"
)

func main() {
	flag.Parse()
	expression := flag.Arg(0)

	res, err := calculator.Calculate(expression)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

