package csl

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
		Message: "incorrect operand type for + operation",
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
				if tr.err != nil {
					return nil, trace
				}
				list.Elems = append(list.Elems, val)
			}
			argNum++
			operand = operand.R
			firstElem = false
		}
	}

	if len(list.Elems) == 1 {
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
