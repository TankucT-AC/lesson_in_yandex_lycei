package main

import (
	"errors"
	"strconv"
)

var (
	InputError = errors.New("uncorrect data")
)

func Priority(elem string) int {
	res := 0
	switch {
	case elem == "(" || elem == ")":
		res = 1
	case elem == "+" || elem == "-":
		res = 2
	case elem == "*" || elem == "/":
		res = 3
	}
	return res
}

func Functions(x, y float64, s string) float64 {
	var res float64
	switch s {
	case "+":
		res = x + y
	case "-":
		res = x - y
	case "*":
		res = x * y
	case "/":
		res = x / y
	}
	return res
}

func Calc(expression string) (float64, error) {

	if expression == "" {
		return 0, InputError
	}

	out := []string{}
	steck := []string{}
	for _, r := range expression {
		s := string(r)
		if _, err := strconv.Atoi(s); err == nil {
			out = append(out, s)
		} else if s != "(" && s != ")" {
			for {
				if !(len(steck) > 0 && Priority(steck[len(steck)-1]) >= Priority(s)) {
					break
				}
				out = append(out, steck[len(steck)-1])
				steck = steck[:len(steck)-1]
			}
			steck = append(steck, s)
		} else if s == "(" {
			steck = append(steck, s)
		} else {
			for {
				if steck[len(steck)-1] == "(" {
					break
				}
				out = append(out, steck[len(steck)-1])
				steck = steck[:len(steck)-1]
			}
			steck = steck[:len(steck)-1]
		}
	}

	for i := len(steck)-1; i > -1; i-- {
		out = append(out, steck[i])
	}

	new_stack := []float64{}
	for _, elem := range out {
		if i, err := strconv.Atoi(elem); err == nil {
			new_stack = append(new_stack, float64(i))
		} else {
			if len(new_stack) < 2 {
				return 0, InputError
			}
			x := new_stack[len(new_stack)-2]
			y := new_stack[len(new_stack)-1]
			new_stack = new_stack[:len(new_stack)-2]
			new_stack = append(new_stack, Functions(x, y, elem))
		}
	}

	return new_stack[0], nil
}
