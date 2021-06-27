package csl

import (
	"reflect"
	"strings"
)

var builtins = map[string]operation{}

func RegisterFunction(pattern string, fn func(env map[string]interface{}, args ...interface{}) (interface{}, error)) {
	builtins[pattern] = fn
}

func init() {
	RegisterFunction("+", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, ErrNumOperands
		}
		var sum int64
		for _, arg := range args {
			x, ok := arg.(int64)
			if !ok {
				return nil, ErrMismatchOperandTypes
			}
			sum += x
		}
		return sum, nil
	})
	RegisterFunction("-", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, ErrNumOperands
		}
		difference, ok := args[0].(int64)
		if !ok {
			return nil, ErrMismatchOperandTypes
		}
		for _, arg := range args[1:] {
			x, ok := arg.(int64)
			if !ok {
				return nil, ErrMismatchOperandTypes
			}
			difference -= x
		}
		return difference, nil
	})
	RegisterFunction("*", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, ErrNumOperands
		}
		var product int64 = 1
		for _, arg := range args {
			x, ok := arg.(int64)
			if !ok {
				return nil, ErrMismatchOperandTypes
			}
			product *= x
		}
		return product, nil
	})
	RegisterFunction("/", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, ErrNumOperands
		}
		quotient, ok := args[0].(int64)
		if !ok {
			return nil, ErrMismatchOperandTypes
		}
		for _, arg := range args[1:] {
			x, ok := arg.(int64)
			if !ok {
				return nil, ErrMismatchOperandTypes
			}
			quotient /= x
		}
		return quotient, nil
	})
	RegisterFunction("define", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, ErrNumOperands
		}
		x, ok := args[0].(Atom)
		if !ok {
			return nil, ErrMismatchOperandTypes
		}
		varName, ok := x.varName()
		env[varName] = args[1]
		return nil, nil
	})
	RegisterFunction("in", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, ErrNumOperands
		}
		x := args[0]
		l, ok := args[1].(List)
		if !ok {
			if str, ok := args[1].(string); ok {
				if substr, ok := x.(string); ok {
					return strings.Contains(str, substr), nil
				} else {
					return nil, ErrMismatchOperandTypes
				}
			}
			return reflect.DeepEqual(x, args[1]), nil
		}
		for _, elem := range l.Elems {
			if reflect.DeepEqual(elem, x) {
				return true, nil
			}
		}
		return false, nil
	})
	RegisterFunction(">", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, ErrNumOperands
		}
		first, ok := args[0].(int64)
		if !ok {
			return nil, ErrMismatchOperandTypes
		}
		for _, arg := range args[1:] {
			x, ok := arg.(int64)
			if !ok {
				return nil, ErrMismatchOperandTypes
			}
			if !(first > x) {
				return false, nil
			}
		}
		return true, nil
	})
	RegisterFunction("<", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, ErrNumOperands
		}
		first, ok := args[0].(int64)
		if !ok {
			return nil, ErrMismatchOperandTypes
		}
		for _, arg := range args[1:] {
			x, ok := arg.(int64)
			if !ok {
				return nil, ErrMismatchOperandTypes
			}
			if !(first < x) {
				return false, nil
			}
		}
		return true, nil
	})
	RegisterFunction(">=", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, ErrNumOperands
		}
		first, ok := args[0].(int64)
		if !ok {
			return nil, ErrMismatchOperandTypes
		}
		for _, arg := range args[1:] {
			x, ok := arg.(int64)
			if !ok {
				return nil, ErrMismatchOperandTypes
			}
			if !(first >= x) {
				return false, nil
			}
		}
		return true, nil
	})
	RegisterFunction("<=", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, ErrNumOperands
		}
		first, ok := args[0].(int64)
		if !ok {
			return nil, ErrMismatchOperandTypes
		}
		for _, arg := range args[1:] {
			x, ok := arg.(int64)
			if !ok {
				return nil, ErrMismatchOperandTypes
			}
			if !(first <= x) {
				return false, nil
			}
		}
		return true, nil
	})
	RegisterFunction("<<=", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, ErrNumOperands
		}
		prev, ok := args[0].(int64)
		if !ok {
			return nil, ErrMismatchOperandTypes
		}
		for _, arg := range args[1:] {
			x, ok := arg.(int64)
			if !ok {
				return nil, ErrMismatchOperandTypes
			}
			if !(prev <= x) {
				return false, nil
			}
			prev = x
		}
		return true, nil
	})
	RegisterFunction(">>=", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, ErrNumOperands
		}
		prev, ok := args[0].(int64)
		if !ok {
			return nil, ErrMismatchOperandTypes
		}
		for _, arg := range args[1:] {
			x, ok := arg.(int64)
			if !ok {
				return nil, ErrMismatchOperandTypes
			}
			if !(prev >= x) {
				return false, nil
			}
			prev = x
		}
		return true, nil
	})
	RegisterFunction(">>", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, ErrNumOperands
		}
		prev, ok := args[0].(int64)
		if !ok {
			return nil, ErrMismatchOperandTypes
		}
		for _, arg := range args[1:] {
			x, ok := arg.(int64)
			if !ok {
				return nil, ErrMismatchOperandTypes
			}
			if !(prev > x) {
				return false, nil
			}
			prev = x
		}
		return true, nil
	})
	RegisterFunction("<<", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, ErrNumOperands
		}
		prev, ok := args[0].(int64)
		if !ok {
			return nil, ErrMismatchOperandTypes
		}
		for _, arg := range args[1:] {
			x, ok := arg.(int64)
			if !ok {
				return nil, ErrMismatchOperandTypes
			}
			if !(prev < x) {
				return false, nil
			}
			prev = x
		}
		return true, nil
	})
	RegisterFunction("nth", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, ErrNumOperands
		}
		i, ok := args[0].(int64)
		if !ok {
			return nil, ErrMismatchOperandTypes
		}
		l, ok := args[1].(List)
		if !ok {
			if s, ok := args[1].(string); ok {
				if i < 0 || i >= int64(len(s)) {
					return nil, ErrOutOfBounds
				}
				return string(s[i]), nil
			}
			l = List{
				Elems: []interface{}{args[1]},
			}
		}
		if i < 0 || i >= int64(len(l.Elems)) {
			return nil, ErrOutOfBounds
		}
		return l.Elems[i], nil
	})
	RegisterFunction("eq", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, ErrNumOperands
		}
		first := args[0]
		for _, arg := range args[1:] {
			if !reflect.DeepEqual(first, arg) {
				return false, nil
			}
		}
		return true, nil
	})
	RegisterFunction("len", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, ErrNumOperands
		}
		if s, ok := args[0].(string); ok {
			return int64(len(s)), nil
		}
		if l, ok := args[0].(List); ok {
			return int64(len(l.Elems)), nil
		}
		return nil, ErrMismatchOperandTypes
	})
	RegisterFunction("if", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 3 {
			return nil, ErrNumOperands
		}
		p, ok := args[0].(bool)
		if !ok {
			return nil, ErrMismatchOperandTypes
		}

		if p {
			return args[1], nil
		}
		return args[2], nil
	})
	RegisterFunction("and", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 1 {
			return nil, ErrNumOperands
		}
		p, ok := args[0].(bool)
		if !ok {
			return nil, ErrMismatchOperandTypes
		}

		for _, arg := range args {
			q, ok := arg.(bool)
			if !ok {
				return nil, ErrMismatchOperandTypes
			}
			p = p && q
		}
		return p, nil
	})
	RegisterFunction("or", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 1 {
			return nil, ErrNumOperands
		}
		p, ok := args[0].(bool)
		if !ok {
			return nil, ErrMismatchOperandTypes
		}

		for _, arg := range args {
			q, ok := arg.(bool)
			if !ok {
				return nil, ErrMismatchOperandTypes
			}
			p = p || q
		}
		return p, nil
	})
	RegisterFunction("not", func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, ErrNumOperands
		}
		p, ok := args[0].(bool)
		if !ok {
			return nil, ErrMismatchOperandTypes
		}

		return !p, nil
	})
}