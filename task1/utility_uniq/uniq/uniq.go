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

func cut(opts Options, str string) string {
	words := strings.Split(str, " ")

	if opts.fields >= len(words) {
		return ""
	}

	resStr := strings.Join(words[opts.fields:], " ")

	if opts.chars >= len(resStr) {
		return ""
	}

	return resStr[opts.chars:]
}

func addStr(result []string, opts Options, count int, prevStr string) []string{
	switch {
	case opts.count:
		prevStr = strconv.Itoa(count) + " " + prevStr
		result = append(result, prevStr)
	case opts.double && count > 1:
		result = append(result, prevStr)
	case opts.unique && count == 1:
		result = append(result, prevStr)
	case !opts.double && !opts.unique && !opts.count:
		result = append(result, prevStr)
	}
	return result
}

func Uniq(strs []string, opts Options) ([]string, error) {
	result := make([]string, 0)
	if (opts.count && opts.double || opts.count && opts.unique || opts.double && opts.unique) {
		err := errors.New("error: incompatible flags")
		return result, err
	}

	var err error
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

	return result, err
}

