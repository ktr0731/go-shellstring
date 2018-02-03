package main

import (
	"errors"
	"fmt"
	"log"
)

func Parse(in string) ([]string, error) {
	stack := []rune{}
	out := []string{}

	var needCloseDoubleQuote, needCloseSingleQuote bool
	s := make([]rune, 0, len(in))
	for _, r := range in {
		log.Println(string(r))
		switch r {
		case '\'':
			if needCloseSingleQuote {
				if stack[len(stack)-1] != '"' {
					return nil, errors.New("open single quote missing")
				}
				stack = stack[:len(stack)-1]
				out = append(out, string(s))
				s = []rune{}
				needCloseSingleQuote = false
			} else {
				stack = append(stack, '"')
				needCloseSingleQuote = true
			}
		case '"':
			if needCloseDoubleQuote {
				if stack[len(stack)-1] != '"' {
					return nil, errors.New("open double quote missing")
				}
				stack = stack[:len(stack)-1]
				out = append(out, string(s))
				s = []rune{}
				needCloseDoubleQuote = false
			} else {
				stack = append(stack, '"')
				needCloseDoubleQuote = true
			}
		case ' ':
			// out of quote
			if len(stack) == 0 {
				if len(s) != 0 {
					out = append(out, string(s))
					s = []rune{}
				}
			} else {
				s = append(s, r)
			}
		default:
			s = append(s, r)
		}
	}
	if len(stack) != 0 {
		return nil, fmt.Errorf("corrensponding symbol missing: %q", stack)
	}
	if len(s) != 0 {
		out = append(out, string(s))
	}
	return out, nil
}
