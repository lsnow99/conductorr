package csllib

import "fmt"

type CSLEvalError struct {
	Message  string
	Internal bool
}

type Trace struct {
	ExprTree []*SExpr
	Err      error
}

type List = []interface{}

type eagerFnSig func(env map[string]interface{}, args ...interface{}) (interface{}, error)
type lazyFnSig func(env map[string]interface{}, args []*SExpr, trace Trace) (interface{}, Trace)

type operation struct {
	fn     eagerFnSig
	lazyFn lazyFnSig
	lazy   bool
}

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
	ErrDivideByZero = &CSLEvalError{
		Message: "divide by zero",
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

func (csl *CSL) Eval(sexprs []*SExpr, env map[string]interface{}) (interface{}, Trace) {
	var res interface{}
	trace := Trace{
		ExprTree: make([]*SExpr, 0),
		Err:      nil,
	}
	for _, tree := range sexprs {
		res, trace = csl.EvalSExpr(tree, env, trace)
		if trace.Err != nil {
			return res, trace
		}
	}
	return res, trace
}

func (csl *CSL) EvalSExpr(sexpr *SExpr, env map[string]interface{}, trace Trace) (interface{}, Trace) {
	list := List{}
	var op *operation
	var firstElem bool = true
	var isDefine bool
	var argNum int
	if sexpr == nil {
		return nil, trace
	}
	// TODO: is this check necessary? Maybe it should return an error if false
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
						trace.Err = ErrBadType
						return nil, trace
					}
					list = append(list, x)
				case stringAtom:
					x, ok := atom.stringVal()
					if !ok {
						trace.Err = ErrBadType
						return nil, trace
					}
					list = append(list, x)
				case boolAtom:
					x, ok := atom.boolVal()
					if !ok {
						trace.Err = ErrBadType
						return nil, trace
					}
					list = append(list, x)
				case varAtom:
					x, ok := atom.varName()
					if !ok {
						trace.Err = ErrBadType
						return nil, trace
					}
					if isDefine && argNum == 0 {
						list = append(list, atom.Atom)
					} else {
						val, ok := env[x]
						if !ok {
							trace.Err = fmt.Errorf("undefined variable: %s", x)
							return nil, trace
						}
						list = append(list, val)
					}
				case importAtom:
					// Ignore import commands
					return nil, trace
				case fnAtom:
					if !firstElem {
						trace.Err = ErrUnexpectedFunction
						return nil, trace
					}

					var ok bool
					fnName, ok := atom.fnName()
					if !ok {
						trace.Err = ErrBadType
						return nil, trace
					}
					if fnName == "define" {
						isDefine = true
					}

					op = &operation{}
					*op, ok = csl.builtins[fnName]
					if !ok {
						trace.Err = ErrNoSuchFunction
						return nil, trace
					}
					// If the function is lazily-evaluated, then call it immediately and just
					// pass the arguments as SExpr arguments, without evaluating them. Return
					// the result of the lazy function
					if op.lazy {
						lazyOps := make([]*SExpr, 0)
						operand = operand.R
						for operand != nil {
							arg := &SExpr{
								Atom: operand.Atom,
								L:    operand.L,
							}
							lazyOps = append(lazyOps, arg)
							operand = operand.R
						}
						return op.lazyFn(env, lazyOps, trace)
					}
					argNum--
				}
			} else if operand.L != nil {
				val, tr := csl.EvalSExpr(operand.L, env, trace)
				trace = tr
				if trace.Err != nil {
					return nil, trace
				}
				list = append(list, val)
			}
			argNum++
			operand = operand.R
			firstElem = false
		}
	}

	if len(list) == 1 && op == nil {
		return list[0], trace
	}
	if op != nil {
		resp, err := op.fn(env, list...)
		if err != nil {
			trace.Err = err
			return nil, trace
		}
		return resp, trace
	} else {
		return list, trace
	}
}

func (csl *CSL) Invoke(sexprs []*SExpr, args ...interface{}) (interface{}, Trace) {
	env := make(map[string]interface{})
	var argsList List
	for _, arg := range args {
		argsList = append(argsList, arg)
	}
	env["args"] = argsList
	result, trace := csl.Eval(sexprs, env)
	return result, trace
}
