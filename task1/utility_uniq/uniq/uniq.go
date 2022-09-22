package uniq

import (
	"errors"
	"fmt"
	"strings"
	"strconv"
)

func Usage() {
	fmt.Println("Usage: uniq [-c | -d | -u] [-i] [-f num] [-s chars] [input_file [output_file]]")
}

func min(l int, r int) int {
	if l > r {
		return r
	} else {
		return l
	}
}

func cut(opts Options, str string) string {
	words := strings.Split(str, " ")
	words = words[min(opts.fields, len(words)):]
	resStr := strings.Join(words, " ")
	return resStr[min(opts.chars, len(resStr)):]
}

func addStr(result []string, opts Options, count int, prevStr string) []string{
	if opts.count {
		prevStr = strconv.Itoa(count) + " " + prevStr
		result = append(result, prevStr)
	} else if opts.double && count > 1 {
		result = append(result, prevStr)
	} else if opts.unique && count == 1{
		result = append(result, prevStr)
	} else if !opts.double && !opts.unique && !opts.count {
		result = append(result, prevStr)
	}
	return result
}

func Uniq(strs []string, opts Options) (result []string, err error) {
	if (opts.count && opts.double || opts.count && opts.unique || opts.double && opts.unique) {
		err = errors.New("error: incompatible flags")
		return
	}
	count := 1
	prevString := strs[0]
	for idx := 1; idx < len(strs); idx++ {
		curStr := cut(opts, strs[idx])
		cutPrev := cut(opts, prevString)
		if curStr == cutPrev {
			count++
		} else if opts.insensitive {
			if !strings.EqualFold(curStr, cutPrev) {
				result = addStr(result, opts, count, prevString)
				count = 1
				prevString = strs[idx]
			}
		} else {
			result = addStr(result, opts, count, prevString)
			count = 1
			prevString = strs[idx]
		}
	}
	result = addStr(result, opts, count, prevString)

	return
}

