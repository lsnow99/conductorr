package csl

type CSLEvalError struct {
	Message  string
	Internal bool
}

var (
	ErrMismatchOperandTypes = &CSLEvalError{
		Message: "incorrect operand type for + operation",
	}
	ErrNumOperands = &CSLEvalError{
		Message: "incorrect number of operands",
	}
	ErrBadType = &CSLEvalError{
		Message: "failed to cast type",
		Internal: true,
	}
	ErrNoSuchFunction = &CSLEvalError{
		Message: "no such function",
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

var builtins = map[atomType]func(args... interface{})(interface{}, error){
	addAtom: func(args... interface{}) (interface{}, error) {
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
	subAtom: func(args... interface{}) (interface{}, error) {
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
	multAtom: func(args... interface{}) (interface{}, error) {
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
	divAtom: func(args... interface{}) (interface{}, error) {
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
}

func Eval(sexpr *SExpr, env map[string]interface{}) (interface{}, error) {
	if sexpr == nil {
		return nil, nil
	}
	if !sexpr.isAtom() {
		if sexpr.L != nil && sexpr.L.isAtom() {
			atom := sexpr.L
			switch typ := sexpr.L.Type(); typ {
			case numberAtom:
				x, ok := atom.numberVal()
				if !ok {
					return nil, ErrBadType
				}
				return x, nil
			case stringAtom:
				x, ok := atom.stringVal()
				if !ok {
					return nil, ErrBadType
				}
				return x, nil
			default:
				op, ok := builtins[typ]
				if !ok {
					return nil, ErrNoSuchFunction
				}
				args := []interface{}{}
				operand := sexpr.R
				for operand != nil {
					arg, err := Eval(operand, env)
					if err != nil {
						return nil, err
					}
					args = append(args, arg)
					operand = operand.R
				}
				return op(args...)
			}
		} else if sexpr.L != nil {
			return Eval(sexpr.L, env)
		}
	}
	return nil, nil
}
