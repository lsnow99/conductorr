package csl

import (
	"reflect"
	"strings"
)

type CSLEvalError struct {
	Message  string
	Internal bool
}

type Trace struct {
	ExprTree []*SExpr
	err      error
}

type List struct {
	Elems []interface{}
}

type operation func(env map[string]interface{}, args ...interface{}) (interface{}, error)

var (
	ErrMismatchOperandTypes = &CSLEvalError{
		Message: "incorrect operand types",
	}
	ErrNumOperands = &CSLEvalError{
		Message: "incorrect number of operands",
	}
	ErrBadType = &CSLEvalError{
		Message:  "failed to cast type",
		Internal: true,
	}
	ErrNoSuchFunction = &CSLEvalError{
		Message: "no such function",
	}
	ErrUnexpectedFunction = &CSLEvalError{
		Message: "unexpected function call when parsing list",
	}
	ErrUndefinedVar = &CSLEvalError{
		Message: "undefined variable",
	}
	ErrOutOfBounds = &CSLEvalError{
		Message: "index out of bounds",
	}
)

func (c CSLEvalError) Error() string {
	str := ""
	if c.Internal {
		str += "internal eval error: "
	}
	str += c.Message
	return str
}

var builtins = map[atomType]operation{
	addAtom: func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	},
	subAtom: func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	},
	multAtom: func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	},
	divAtom: func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	},
	defineAtom: func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	},
	inAtom: func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	},
	greaterAtom: func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	},
	lessAtom: func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	},
	greaterEqualAtom: func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	},
	lessEqualAtom: func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	},
	allLessEqualAtom: func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	},
	allGreaterEqualAtom: func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	},
	allGreaterAtom: func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	},
	allLessAtom: func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	},
	nthAtom: func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	},
	eqAtom: func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	},
	lengthAtom: func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
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
	},
}

func Eval(sexprs []*SExpr, env map[string]interface{}) (interface{}, Trace) {
	var res interface{}
	trace := Trace{
		ExprTree: make([]*SExpr, 0),
		err:      nil,
	}
	for _, tree := range sexprs {
		res, trace = EvalSExpr(tree, env, trace)
		if trace.err != nil {
			return res, trace
		}
	}
	return res, trace
}

func EvalSExpr(sexpr *SExpr, env map[string]interface{}, trace Trace) (interface{}, Trace) {
	list := List{}
	var op operation
	var firstElem bool = true
	var isDefine bool
	var argNum int
	if sexpr == nil {
		return nil, trace
	}
	if !sexpr.isAtom() {
		operand := sexpr
		for operand != nil {
			trace.ExprTree = append(trace.ExprTree, operand)
			if operand.L != nil && operand.L.isAtom() {
				atom := operand.L
				switch typ := operand.L.Type(); typ {
				case numberAtom:
					x, ok := atom.numberVal()
					if !ok {
						trace.err = ErrBadType
						return nil, trace
					}
					list.Elems = append(list.Elems, x)
				case stringAtom:
					x, ok := atom.stringVal()
					if !ok {
						trace.err = ErrBadType
						return nil, trace
					}
					list.Elems = append(list.Elems, x)
				case varAtom:
					x, ok := atom.varName()
					if !ok {
						trace.err = ErrBadType
						return nil, trace
					}
					if isDefine && argNum == 0 {
						list.Elems = append(list.Elems, atom.Atom)
					} else {
						val, ok := env[x]
						if !ok {
							trace.err = ErrUndefinedVar
							return nil, trace
						}
						list.Elems = append(list.Elems, val)
					}
				default:
					if !firstElem {
						trace.err = ErrUnexpectedFunction
						return nil, trace
					}
					var ok bool
					op, ok = builtins[typ]
					if typ == defineAtom {
						isDefine = true
					}
					if !ok {
						trace.err = ErrNoSuchFunction
						return nil, trace
					}
					argNum--
				}
			} else if operand.L != nil {
				val, tr := EvalSExpr(operand.L, env, trace)
				trace = tr
				if trace.err != nil {
					return nil, trace
				}
				list.Elems = append(list.Elems, val)
			}
			argNum++
			operand = operand.R
			firstElem = false
		}
	}

	if len(list.Elems) == 1 && op == nil {
		return list.Elems[0], trace
	}
	if op != nil {
		resp, err := op(env, list.Elems...)
		if err != nil {
			trace.err = err
			return nil, trace
		}
		return resp, trace
	} else {
		return list, trace
	}
}
