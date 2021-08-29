package csl

import (
	"fmt"
	"reflect"
	"strings"
)

var builtins = map[string]operation{}

/*
RegisterFunction registers a builtin function to CSL that can either be eager or lazy, as specified by
the `lazy` flag. The function name must meet the same criteria as a variable name to be considered valid.
Invalid function names will cause RegisterFunction to return an error. If the provided function name is
already in use, the newly registered function will override and replace it. If `lazy` is true, then
lazyFn must be supplied, else eagerFn must be supplied, else a non-nil error will be returned. Lazy
functions must handle their own evaluation of SExpr pointers.
*/
func RegisterFunction(name string, lazy bool, eagerFn eagerFnSig, lazyFn lazyFnSig) error {
	// Check to make sure this name doesn't conflict with any of the symbols in the language definition
	// (i.e. a name can not be a '(' or '"', or a number, etc))
	// If the token parses as a variable name that means that it is a valid function name
	if strings.TrimSpace(name) != name {
		return fmt.Errorf("function name %s has leading or trailing whitespace", name)
	}
	if strings.Contains(name, "\"") {
		return fmt.Errorf("function name %s contains illegal character '\"'", name)
	}

	toks := NewTokens(name)
	if len(toks) < 1 {
		return fmt.Errorf("did not parse any tokens, expected one for function name %s", name)
	}
	if len(toks) > 1 {
		return fmt.Errorf("parsed multiple tokens, expected one for function name %s", name)
	}
	if toks[0].typ != symbolToken {
		return fmt.Errorf("did not parse function name %s as a symbol token", name)
	}

	if lazy && lazyFn == nil {
		return fmt.Errorf("requested to register lazy function but no lazy function supplied")
	}
	if !lazy && eagerFn == nil {
		return fmt.Errorf("requested to register eager function but no eager function supplied")
	}

	builtins[name] = operation{
		fn:     eagerFn,
		lazy:   lazy,
		lazyFn: lazyFn,
	}
	return nil
}

func init() {
	RegisterFunction("+", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	}, nil)
	RegisterFunction("-", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	}, nil)
	RegisterFunction("*", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	}, nil)
	RegisterFunction("/", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	}, nil)
	RegisterFunction("define", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, ErrNumOperands
		}
		x, ok := args[0].(Atom)
		if !ok {
			return nil, ErrMismatchOperandTypes
		}
		varName, ok := x.varName()
		if !ok {
			return nil, ErrBadType
		}
		env[varName] = args[1]
		return nil, nil
	}, nil)
	RegisterFunction("in", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	}, nil)
	RegisterFunction(">", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	}, nil)
	RegisterFunction("<", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	}, nil)
	RegisterFunction(">=", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	}, nil)
	RegisterFunction("<=", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	}, nil)
	RegisterFunction("<<=", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	}, nil)
	RegisterFunction(">>=", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	}, nil)
	RegisterFunction(">>", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	}, nil)
	RegisterFunction("<<", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	}, nil)
	RegisterFunction("nth", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, ErrNumOperands
		}
		i, ok := args[0].(int64)
		if !ok {
			return nil, ErrMismatchOperandTypes
		}
		l, ok := args[1].(List)
		if !ok {
			l.Elems = []interface{}{args[1]}
		}
		if i < 0 || i >= int64(len(l.Elems)) {
			return nil, ErrOutOfBounds
		}
		return l.Elems[i], nil
	}, nil)
	RegisterFunction("nths", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, ErrNumOperands
		}
		i, ok := args[0].(int64)
		if !ok {
			return nil, ErrMismatchOperandTypes
		}
		s, ok := args[1].(string)
		if !ok {
			return nil, ErrMismatchOperandTypes
		}
		if i < 0 || i >= int64(len(s)) {
			return nil, ErrOutOfBounds
		}
		return string(s[i]), nil
	}, nil)
	RegisterFunction("eq", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	}, nil)
	RegisterFunction("len", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, ErrNumOperands
		}
		if l, ok := args[0].(List); ok {
			return int64(len(l.Elems)), nil
		}
		return int64(1), nil
	}, nil)
	RegisterFunction("lens", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, ErrNumOperands
		}
		if s, ok := args[0].(string); ok {
			return int64(len(s)), nil
		}
		return nil, ErrMismatchOperandTypes
	}, nil)
	RegisterFunction("append", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, ErrNumOperands
		}
		if l, ok := args[0].(List); ok {
			l.Elems = append(l.Elems, args[1:]...)
			return l, nil
		}
		return nil, ErrMismatchOperandTypes
	}, nil)
	RegisterFunction("appendleft", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, ErrNumOperands
		}
		if l, ok := args[0].(List); ok {
			l.Elems = append(args[1:], l.Elems...)
			return l, nil
		}
		return nil, ErrMismatchOperandTypes
	}, nil)
	RegisterFunction("pop", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, ErrNumOperands
		}
		if l, ok := args[0].(List); ok {
			l.Elems = l.Elems[:len(l.Elems)-1]
			return l, nil
		}
		return nil, ErrMismatchOperandTypes
	}, nil)
	RegisterFunction("popleft", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, ErrNumOperands
		}
		if l, ok := args[0].(List); ok {
			l.Elems = l.Elems[1:]
			return l, nil
		}
		return nil, ErrMismatchOperandTypes
	}, nil)
	RegisterFunction("peek", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, ErrNumOperands
		}
		if l, ok := args[0].(List); ok {
			return l.Elems[len(l.Elems)-1], nil
		}
		return nil, ErrMismatchOperandTypes
	}, nil)
	RegisterFunction("peekleft", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, ErrNumOperands
		}
		if l, ok := args[0].(List); ok {
			return l.Elems[0], nil
		}
		return nil, ErrMismatchOperandTypes
	}, nil)
	RegisterFunction("if", true, nil, func(env map[string]interface{}, args []*SExpr, trace Trace) (interface{}, Trace) {
		if len(args) != 3 {
			trace.Err = ErrNumOperands
			return nil, trace
		}
		cond, trace := EvalSExpr(args[0], env, trace)
		p, ok := cond.(bool)
		if !ok {
			trace.Err = ErrMismatchOperandTypes
			return nil, trace
		}

		if p {
			return EvalSExpr(args[1], env, trace)
		}
		return EvalSExpr(args[2], env, trace)
	})
	RegisterFunction("and", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	}, nil)
	RegisterFunction("or", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	}, nil)
	RegisterFunction("not", false, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, ErrNumOperands
		}
		p, ok := args[0].(bool)
		if !ok {
			return nil, ErrMismatchOperandTypes
		}

		return !p, nil
	}, nil)
}
