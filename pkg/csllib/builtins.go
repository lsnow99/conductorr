package csllib

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

/*
RegisterFunction registers a builtin function to CSL that can either be eager or lazy, as specified by
the `lazy` flag. The function name must meet the same criteria as a variable name to be considered valid.
Invalid function names will cause RegisterFunction to return an error. If the provided function name is
already in use, the newly registered function will override and replace it. If `lazy` is true, then
lazyFn must be supplied, else eagerFn must be supplied, else a non-nil error will be returned. Lazy
functions must handle their own evaluation of SExpr pointers. The builtin flag specifies that the function
is builtin to the CSL instance, and is not a module.
*/
func (csl *CSL) RegisterFunction(name string, lazy, builtin bool, eagerFn eagerFnSig, lazyFn lazyFnSig) error {
	// Check to make sure this name doesn't conflict with any of the symbols in the language definition
	// (i.e. a name can not be a '(' or '"', or a number, etc))
	// If the token parses as a variable name that means that it is a valid function name
	if strings.TrimSpace(name) != name {
		return fmt.Errorf("function name %s has leading or trailing whitespace", name)
	}
	if strings.Contains(name, "\"") {
		return fmt.Errorf("function name %s contains illegal character '\"'", name)
	}

	fnNameErr := checkValidFunctionName(name)
	if fnNameErr != nil {
		return fnNameErr
	}

	if lazy && lazyFn == nil {
		return fmt.Errorf("requested to register lazy function but no lazy function supplied")
	}
	if !lazy && eagerFn == nil {
		return fmt.Errorf("requested to register eager function but no eager function supplied")
	}

	csl.builtins[name] = operation{
		fn:      eagerFn,
		builtin: builtin,
		lazy:    lazy,
		lazyFn:  lazyFn,
	}
	return nil
}

// CopyWithBuiltins returns a new CSL instance with the builtin functions from the original
// csl instance pre-registered.
func (csl *CSL) CopyWithBuiltins() *CSL {
	cslNew := NewCSL(false)
	for k, v := range csl.builtins {
		if v.builtin {
			cslNew.builtins[k] = operation{
				fn:      v.fn,
				builtin: v.builtin,
				lazy:    v.lazy,
				lazyFn:  v.lazyFn,
			}
		}
	}
	return cslNew
}

// AliasFunction creates an alias to an existing function.
// The alias parameter must be a valid function name. If the function
// name is already registered to a different function, it will be
// remapped. The source parameter must be a currently-registered
// function.
func (csl *CSL) AliasFunction(alias, source string) error {
	src, ok := csl.builtins[alias]
	if !ok {
		return fmt.Errorf("attempted to register alias %s to nonexistant function %s", alias, source)
	}

	fnNameErr := checkValidFunctionName(alias)
	if fnNameErr != nil {
		return fnNameErr
	}

	csl.builtins[alias] = src

	return nil
}

func (csl *CSL) RegisterDefaults() {
	// Keep track of the fact that default builtins have been registered
	csl.didRegisterDefaults = true

	csl.RegisterFunction("+", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, buildError("+", ErrNumOperands, args)
		}
		var sum int64
		for _, arg := range args {
			x, ok := arg.(int64)
			if !ok {
				return nil, buildError("+", ErrMismatchOperandTypes, args)
			}
			sum += x
		}
		return sum, nil
	}, nil)
	csl.RegisterFunction("-", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, buildError("-", ErrNumOperands, args)
		}
		difference, ok := args[0].(int64)
		if !ok {
			return nil, buildError("-", ErrMismatchOperandTypes, args)
		}
		for _, arg := range args[1:] {
			x, ok := arg.(int64)
			if !ok {
				return nil, buildError("-", ErrMismatchOperandTypes, args)
			}
			difference -= x
		}
		return difference, nil
	}, nil)
	csl.RegisterFunction("*", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, buildError("*", ErrNumOperands, args)
		}
		var product int64 = 1
		for _, arg := range args {
			x, ok := arg.(int64)
			if !ok {
				return nil, buildError("*", ErrMismatchOperandTypes, args)
			}
			product *= x
		}
		return product, nil
	}, nil)
	csl.RegisterFunction("/", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, buildError("/", ErrNumOperands, args)
		}
		quotient, ok := args[0].(int64)
		if !ok {
			return nil, buildError("/", ErrMismatchOperandTypes, args)
		}
		for _, arg := range args[1:] {
			x, ok := arg.(int64)
			if !ok {
				return nil, buildError("/", ErrMismatchOperandTypes, args)
			}
			if x == 0 {
				return nil, buildError("/", ErrDivideByZero, args)
			}
			quotient /= x
		}
		return quotient, nil
	}, nil)
	csl.RegisterFunction("define", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, buildError("define", ErrNumOperands, args)
		}
		x, ok := args[0].(Atom)
		if !ok {
			return nil, buildError("define", ErrMismatchOperandTypes, args)
		}
		varName, ok := x.varName()
		if !ok {
			return nil, buildError("define", ErrBadType, args)
		}
		env[varName] = args[1]
		return nil, nil
	}, nil)
	csl.RegisterFunction("in", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, buildError("in", ErrNumOperands, args)
		}
		x := args[0]
		l, ok := args[1].(List)
		if !ok {
			return nil, buildError("in", ErrMismatchOperandTypes, args)
		}
		for _, elem := range l {
			if reflect.DeepEqual(elem, x) {
				return true, nil
			}
		}
		return false, nil
	}, nil)
	csl.RegisterFunction("contains", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, buildError("contains", ErrNumOperands, args)
		}
		var str, substr string
		var ok bool
		if str, ok = args[0].(string); !ok {
			return nil, buildError("contains", ErrMismatchOperandTypes, args)
		}
		if substr, ok = args[1].(string); !ok {
			return nil, buildError("contains", ErrMismatchOperandTypes, args)
		}
		return strings.Contains(str, substr), nil
	}, nil)
	csl.RegisterFunction("substr", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 || len(args) > 3 {
			return nil, buildError("substr", ErrNumOperands, args)
		}
		var str string
		var a, b int64
		var ok bool
		if str, ok = args[0].(string); !ok {
			return nil, buildError("substr", ErrMismatchOperandTypes, args)
		}
		if a, ok = args[1].(int64); !ok {
			return nil, buildError("substr", ErrMismatchOperandTypes, args)
		}
		if len(args) == 3 {
			if b, ok = args[2].(int64); !ok {
				return nil, buildError("substr", ErrMismatchOperandTypes, args)
			}
		} else {
			b = int64(len(str))
		}
		return str[a:b], nil
	}, nil)
	csl.RegisterFunction(">", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, buildError(">", ErrNumOperands, args)
		}
		first, ok := args[0].(int64)
		if !ok {
			return nil, buildError(">", ErrMismatchOperandTypes, args)
		}
		for _, arg := range args[1:] {
			x, ok := arg.(int64)
			if !ok {
				return nil, buildError(">", ErrMismatchOperandTypes, args)
			}
			if !(first > x) {
				return false, nil
			}
		}
		return true, nil
	}, nil)
	csl.RegisterFunction("<", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, buildError("<", ErrNumOperands, args)
		}
		first, ok := args[0].(int64)
		if !ok {
			return nil, buildError("<", ErrMismatchOperandTypes, args)
		}
		for _, arg := range args[1:] {
			x, ok := arg.(int64)
			if !ok {
				return nil, buildError("<", ErrMismatchOperandTypes, args)
			}
			if !(first < x) {
				return false, nil
			}
		}
		return true, nil
	}, nil)
	csl.RegisterFunction(">=", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, buildError(">=", ErrNumOperands, args)
		}
		first, ok := args[0].(int64)
		if !ok {
			return nil, buildError(">=", ErrMismatchOperandTypes, args)
		}
		for _, arg := range args[1:] {
			x, ok := arg.(int64)
			if !ok {
				return nil, buildError(">=", ErrMismatchOperandTypes, args)
			}
			if !(first >= x) {
				return false, nil
			}
		}
		return true, nil
	}, nil)
	csl.RegisterFunction("<=", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, buildError("<=", ErrNumOperands, args)
		}
		first, ok := args[0].(int64)
		if !ok {
			return nil, buildError("<=", ErrMismatchOperandTypes, args)
		}
		for _, arg := range args[1:] {
			x, ok := arg.(int64)
			if !ok {
				return nil, buildError("<=", ErrMismatchOperandTypes, args)
			}
			if !(first <= x) {
				return false, nil
			}
		}
		return true, nil
	}, nil)
	csl.RegisterFunction("<<=", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, buildError("<<=", ErrNumOperands, args)
		}
		prev, ok := args[0].(int64)
		if !ok {
			return nil, buildError("<<=", ErrMismatchOperandTypes, args)
		}
		for _, arg := range args[1:] {
			x, ok := arg.(int64)
			if !ok {
				return nil, buildError("<<=", ErrMismatchOperandTypes, args)
			}
			if !(prev <= x) {
				return false, nil
			}
			prev = x
		}
		return true, nil
	}, nil)
	csl.RegisterFunction(">>=", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, buildError(">>=", ErrNumOperands, args)
		}
		prev, ok := args[0].(int64)
		if !ok {
			return nil, buildError(">>=", ErrMismatchOperandTypes, args)
		}
		for _, arg := range args[1:] {
			x, ok := arg.(int64)
			if !ok {
				return nil, buildError(">>=", ErrMismatchOperandTypes, args)
			}
			if !(prev >= x) {
				return false, nil
			}
			prev = x
		}
		return true, nil
	}, nil)
	csl.RegisterFunction(">>", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, buildError(">>", ErrNumOperands, args)
		}
		prev, ok := args[0].(int64)
		if !ok {
			return nil, buildError(">>", ErrMismatchOperandTypes, args)
		}
		for _, arg := range args[1:] {
			x, ok := arg.(int64)
			if !ok {
				return nil, buildError(">>", ErrMismatchOperandTypes, args)
			}
			if !(prev > x) {
				return false, nil
			}
			prev = x
		}
		return true, nil
	}, nil)
	csl.RegisterFunction("<<", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, buildError("<<", ErrNumOperands, args)
		}
		prev, ok := args[0].(int64)
		if !ok {
			return nil, buildError("<<", ErrMismatchOperandTypes, args)
		}
		for _, arg := range args[1:] {
			x, ok := arg.(int64)
			if !ok {
				return nil, buildError("<<", ErrMismatchOperandTypes, args)
			}
			if !(prev < x) {
				return false, nil
			}
			prev = x
		}
		return true, nil
	}, nil)
	csl.RegisterFunction("><", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 3 {
			return nil, buildError("><", ErrNumOperands, args)
		}
		x, ok := args[0].(int64)
		if !ok {
			return nil, buildError("><", ErrMismatchOperandTypes, args)
		}
		y, ok := args[1].(int64)
		if !ok {
			return nil, buildError("><", ErrMismatchOperandTypes, args)
		}
		for _, arg := range args[2:] {
			v, ok := arg.(int64)
			if !ok {
				return nil, buildError("><", ErrMismatchOperandTypes, args)
			}
			if !(v > x) || !(v < y) {
				return false, nil
			}
		}
		return true, nil
	}, nil)
	csl.RegisterFunction(">=<=", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 3 {
			return nil, buildError(">=<=", ErrNumOperands, args)
		}
		x, ok := args[0].(int64)
		if !ok {
			return nil, buildError(">=<=", ErrMismatchOperandTypes, args)
		}
		y, ok := args[1].(int64)
		if !ok {
			return nil, buildError(">=<=", ErrMismatchOperandTypes, args)
		}
		for _, arg := range args[2:] {
			v, ok := arg.(int64)
			if !ok {
				return nil, buildError(">=<=", ErrMismatchOperandTypes, args)
			}
			if !(v >= x) || !(v <= y) {
				return false, nil
			}
		}
		return true, nil
	}, nil)
	csl.RegisterFunction(">=<", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 3 {
			return nil, buildError(">=<", ErrNumOperands, args)
		}
		x, ok := args[0].(int64)
		if !ok {
			return nil, buildError(">=<", ErrMismatchOperandTypes, args)
		}
		y, ok := args[1].(int64)
		if !ok {
			return nil, buildError(">=<", ErrMismatchOperandTypes, args)
		}
		for _, arg := range args[2:] {
			v, ok := arg.(int64)
			if !ok {
				return nil, buildError(">=<", ErrMismatchOperandTypes, args)
			}
			if !(v >= x) || !(v < y) {
				return false, nil
			}
		}
		return true, nil
	}, nil)
	csl.RegisterFunction("><=", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 3 {
			return nil, buildError("><=", ErrNumOperands, args)
		}
		x, ok := args[0].(int64)
		if !ok {
			return nil, buildError("><=", ErrMismatchOperandTypes, args)
		}
		y, ok := args[1].(int64)
		if !ok {
			return nil, buildError("><=", ErrMismatchOperandTypes, args)
		}
		for _, arg := range args[2:] {
			v, ok := arg.(int64)
			if !ok {
				return nil, buildError("><=", ErrMismatchOperandTypes, args)
			}
			if !(v > x) || !(v <= y) {
				return false, nil
			}
		}
		return true, nil
	}, nil)
	csl.RegisterFunction("<>", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 3 {
			return nil, buildError("<>", ErrNumOperands, args)
		}
		x, ok := args[0].(int64)
		if !ok {
			return nil, buildError("<>", ErrMismatchOperandTypes, args)
		}
		y, ok := args[1].(int64)
		if !ok {
			return nil, buildError("<>", ErrMismatchOperandTypes, args)
		}
		for _, arg := range args[2:] {
			v, ok := arg.(int64)
			if !ok {
				return nil, buildError("<>", ErrMismatchOperandTypes, args)
			}
			if !(v < x) && !(v > y) {
				return false, nil
			}
		}
		return true, nil
	}, nil)
	csl.RegisterFunction("<=>=", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 3 {
			return nil, buildError("<=>=", ErrNumOperands, args)
		}
		x, ok := args[0].(int64)
		if !ok {
			return nil, buildError("<=>=", ErrMismatchOperandTypes, args)
		}
		y, ok := args[1].(int64)
		if !ok {
			return nil, buildError("<=>=", ErrMismatchOperandTypes, args)
		}
		for _, arg := range args[2:] {
			v, ok := arg.(int64)
			if !ok {
				return nil, buildError("<=>=", ErrMismatchOperandTypes, args)
			}
			if !(v <= x) && !(v >= y) {
				return false, nil
			}
		}
		return true, nil
	}, nil)
	csl.RegisterFunction("<>=", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 3 {
			return nil, buildError("<>=", ErrNumOperands, args)
		}
		x, ok := args[0].(int64)
		if !ok {
			return nil, buildError("<>=", ErrMismatchOperandTypes, args)
		}
		y, ok := args[1].(int64)
		if !ok {
			return nil, buildError("<>=", ErrMismatchOperandTypes, args)
		}
		for _, arg := range args[2:] {
			v, ok := arg.(int64)
			if !ok {
				return nil, buildError("<>=", ErrMismatchOperandTypes, args)
			}
			if !(v < x) && !(v >= y) {
				return false, nil
			}
		}
		return true, nil
	}, nil)
	csl.RegisterFunction("<=>", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 3 {
			return nil, buildError("<=>", ErrNumOperands, args)
		}
		x, ok := args[0].(int64)
		if !ok {
			return nil, buildError("<=>", ErrMismatchOperandTypes, args)
		}
		y, ok := args[1].(int64)
		if !ok {
			return nil, buildError("<=>", ErrMismatchOperandTypes, args)
		}
		for _, arg := range args[2:] {
			v, ok := arg.(int64)
			if !ok {
				return nil, buildError("<=>", ErrMismatchOperandTypes, args)
			}
			if !(v <= x) && !(v > y) {
				return false, nil
			}
		}
		return true, nil
	}, nil)
	csl.RegisterFunction("nth", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, buildError("nth", ErrNumOperands, args...)
		}
		i, ok := args[0].(int64)
		if !ok {
			return nil, buildError("nth", ErrMismatchOperandTypes, args...)
		}
		l, ok := args[1].(List)
		if !ok {
			l = []interface{}{args[1]}
		}
		if i < 0 || i >= int64(len(l)) {
			return nil, buildError("nth", ErrOutOfBounds, args...)
		}
		return l[i], nil
	}, nil)
	csl.RegisterFunction("nthstr", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, buildError("nthstr", ErrNumOperands, args...)
		}
		i, ok := args[0].(int64)
		if !ok {
			return nil, buildError("nthstr", ErrMismatchOperandTypes, args...)
		}
		s, ok := args[1].(string)
		if !ok {
			return nil, buildError("nthstr", ErrMismatchOperandTypes, args...)
		}
		if i < 0 || i >= int64(len(s)) {
			return nil, buildError("nthstr", ErrOutOfBounds, args...)
		}
		return string(s[i]), nil
	}, nil)
	csl.RegisterFunction("eq", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, buildError("eq", ErrNumOperands, args...)
		}
		first := args[0]
		for _, arg := range args[1:] {
			if !reflect.DeepEqual(first, arg) {
				return false, nil
			}
		}
		return true, nil
	}, nil)
	csl.RegisterFunction("len", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, buildError("len", ErrNumOperands, args...)
		}
		if l, ok := args[0].(List); ok {
			return int64(len(l)), nil
		}
		return int64(1), nil
	}, nil)
	csl.RegisterFunction("lenstr", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, buildError("lenstr", ErrNumOperands, args...)
		}
		if s, ok := args[0].(string); ok {
			return int64(len(s)), nil
		}
		return nil, buildError("lenstr", ErrMismatchOperandTypes, args...)
	}, nil)
	csl.RegisterFunction("append", true, true, nil, func(env map[string]interface{}, args []*SExpr, trace Trace) (interface{}, Trace) {
		if len(args) < 2 {
			trace.Err = buildError("append", ErrNumOperands, args)
			return nil, trace
		}
		if args[0].L == nil || !args[0].L.isAtom() {
			trace.Err = buildError("append", ErrMismatchOperandTypes, args)
			return nil, trace
		}
		atom := args[0].L
		x, ok := atom.varName()
		if !ok {
			trace.Err = buildError("append", ErrMismatchOperandTypes, args)
			return nil, trace
		}
		v, ok := env[x]
		if !ok {
			v = List{}
		}
		l, ok := v.(List)
		if !ok {
			trace.Err = buildError("append", ErrMismatchOperandTypes, args)
			return nil, trace
		}
		for _, arg := range args[1:] {
			value, trace := csl.EvalSExpr(arg, env, trace)
			if trace.Err != nil {
				return nil, trace
			}
			l = append(l, value)
		}
		env[x] = l
		return nil, trace
	})
	csl.RegisterFunction("appendleft", true, true, nil, func(env map[string]interface{}, args []*SExpr, trace Trace) (interface{}, Trace) {
		if len(args) < 2 {
			trace.Err = buildError("appendleft", ErrNumOperands, args)
			return nil, trace
		}
		if args[0].L == nil || !args[0].L.isAtom() {
			trace.Err = buildError("appendleft", ErrMismatchOperandTypes, args)
			return nil, trace
		}
		atom := args[0].L
		x, ok := atom.varName()
		if !ok {
			trace.Err = buildError("appendleft", ErrMismatchOperandTypes, args)
			return nil, trace
		}
		v, ok := env[x]
		if !ok {
			v = List{}
		}
		l, ok := v.(List)
		if !ok {
			trace.Err = buildError("appendleft", ErrMismatchOperandTypes, args)
			return nil, trace
		}
		values := make([]interface{}, 0, len(args))
		for _, arg := range args[1:] {
			value, trace := csl.EvalSExpr(arg, env, trace)
			if trace.Err != nil {
				return nil, trace
			}
			values = append(values, value)
		}
		l = append(values, l...)
		env[x] = l
		return nil, trace
	})
  csl.RegisterFunction("extend", true, true, nil,func(env map[string]interface{}, args []*SExpr, trace Trace) (interface{}, Trace) {
    // Accepts two arguments
    if len(args) != 2 {
      trace.Err = buildError("extend", ErrNumOperands, args)
      return nil, trace
    }

    if args[0].L == nil || !args[0].L.isAtom() {
      trace.Err = buildError("extend", ErrMismatchOperandTypes, args)
      return nil, trace
    }
    atom := args[0].L
    x, ok := atom.varName()
    if !ok {
      trace.Err = buildError("extend", ErrMismatchOperandTypes, args)
      return nil, trace
    }
    av, ok := env[x]
    if !ok {
      av = List{}
    }
    a, ok := av.(List)
    if !ok {
      trace.Err = buildError("extend", ErrMismatchOperandTypes, args)
      return nil, trace
    }

    bv, trace := csl.EvalSExpr(args[1], env, trace)
    if trace.Err != nil {
      return nil, trace
    }

    b, ok := bv.(List)
    if !ok {
      trace.Err = buildError("extend", ErrMismatchOperandTypes, args)
      return nil, trace
    }
    a = append(a, b...)
    env[x] = a
    return nil, trace
  })
	csl.RegisterFunction("pop", true, true, nil, func(env map[string]interface{}, args []*SExpr, trace Trace) (interface{}, Trace) {
		if len(args) != 1 {
			trace.Err = buildError("pop", ErrNumOperands, args)
			return nil, trace
		}
		if args[0].L == nil || !args[0].L.isAtom() {
			trace.Err = buildError("pop", ErrMismatchOperandTypes, args)
			return nil, trace
		}
		atom := args[0].L
		x, ok := atom.varName()
		if !ok {
			trace.Err = buildError("pop", ErrMismatchOperandTypes, args)
			return nil, trace
		}
		v, ok := env[x]
		if !ok {
			v = List{}
		}
		l, ok := v.(List)
		if !ok {
			trace.Err = buildError("pop", ErrMismatchOperandTypes, args)
			return nil, trace
		}
		if len(l) == 0 {
			return nil, trace
		}
		val := l[len(l)-1]
		l = l[:len(l)-1]
		env[x] = l
		return val, trace
	})
	csl.RegisterFunction("popleft", true, true, nil, func(env map[string]interface{}, args []*SExpr, trace Trace) (interface{}, Trace) {
		if len(args) != 1 {
			trace.Err = buildError("popleft", ErrNumOperands, args)
			return nil, trace
		}
		if args[0].L == nil || !args[0].L.isAtom() {
			trace.Err = buildError("popleft", ErrMismatchOperandTypes, args)
			return nil, trace
		}
		atom := args[0].L
		x, ok := atom.varName()
		if !ok {
			trace.Err = buildError("popleft", ErrMismatchOperandTypes, args)
			return nil, trace
		}
		v, ok := env[x]
		if !ok {
			v = List{}
		}
		l, ok := v.(List)
		if !ok {
			trace.Err = buildError("popleft", ErrMismatchOperandTypes, args)
			return nil, trace
		}
		if len(l) == 0 {
			return nil, trace
		}
		val := l[0]
		l = l[1:]
		env[x] = l
		return val, trace
	})
	csl.RegisterFunction("peek", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, buildError("peek", ErrNumOperands, args...)
		}
		if l, ok := args[0].(List); ok {
			return l[len(l)-1], nil
		}
		return nil, buildError("pop", ErrMismatchOperandTypes, args...)
	}, nil)
	csl.RegisterFunction("peekleft", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, buildError("peekleft", ErrNumOperands, args...)
		}
		if l, ok := args[0].(List); ok {
			return l[0], nil
		}
		return nil, buildError("peekleft", ErrMismatchOperandTypes, args...)
	}, nil)
	csl.RegisterFunction("if", true, true, nil, func(env map[string]interface{}, args []*SExpr, trace Trace) (interface{}, Trace) {
		if len(args) != 3 {
			trace.Err = buildError("if", ErrNumOperands, args)
			return nil, trace
		}
		cond, trace := csl.EvalSExpr(args[0], env, trace)
		p, ok := cond.(bool)
		if !ok {
			trace.Err = buildError("if", ErrMismatchOperandTypes, args)
			return nil, trace
		}

		if p {
			return csl.EvalSExpr(args[1], env, trace)
		}
		return csl.EvalSExpr(args[2], env, trace)
	})
	csl.RegisterFunction("and", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 1 {
			return nil, buildError("and", ErrNumOperands, args...)
		}
		p, ok := args[0].(bool)
		if !ok {
			return nil, buildError("and", ErrMismatchOperandTypes, args...)
		}

		for _, arg := range args {
			q, ok := arg.(bool)
			if !ok {
				return nil, buildError("and", ErrMismatchOperandTypes, args...)
			}
			p = p && q
		}
		return p, nil
	}, nil)
	csl.RegisterFunction("or", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 1 {
			return nil, buildError("or", ErrNumOperands, args...)
		}
		p, ok := args[0].(bool)
		if !ok {
			return nil, buildError("and", ErrMismatchOperandTypes, args...)
		}

		for _, arg := range args {
			q, ok := arg.(bool)
			if !ok {
				return nil, buildError("and", ErrMismatchOperandTypes, args...)
			}
			p = p || q
		}
		return p, nil
	}, nil)
	csl.RegisterFunction("not", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 1 {
			return nil, buildError("not", ErrNumOperands, args...)
		}
		p, ok := args[0].(bool)
		if !ok {
			return nil, buildError("not", ErrMismatchOperandTypes, args...)
		}

		return !p, nil
	}, nil)
	csl.RegisterFunction("join", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 {
			return nil, buildError("join", ErrNumOperands, args...)
		}
		separator, ok := args[0].(string)
		if !ok {
			return nil, buildError("join", ErrMismatchOperandTypes, args...)
		}
		str := ""
		for i, arg := range args[1:] {
			list, ok := arg.(List)
			if !ok {
				return nil, buildError("join", ErrMismatchOperandTypes, args...)
			}
			for j, elem := range list {
				if i != len(args)-1 && j != len(list)-1 {
					str += fmt.Sprintf("%v%s", elem, separator)
				} else {
					str += fmt.Sprintf("%v", elem)
				}
			}
		}
		return str, nil
	}, nil)
	csl.RegisterFunction("split", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, buildError("split", ErrNumOperands, args...)
		}
		str, ok := args[0].(string)
		if !ok {
			return nil, buildError("split", ErrMismatchOperandTypes, args...)
		}
		delimiter, ok := args[1].(string)
		if !ok {
			return nil, buildError("split", ErrMismatchOperandTypes, args...)
		}
		l := List{}
		split := strings.Split(str, delimiter)
		for _, s := range split {
			l = append(l, s)
		}
		return l, nil
	}, nil)
	csl.RegisterFunction("match", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, buildError("match", ErrNumOperands, args...)
		}
		var pattern, str string
		var ok bool
		if pattern, ok = args[0].(string); !ok {
			return nil, buildError("match", ErrMismatchOperandTypes, args...)
		}
		if str, ok = args[1].(string); !ok {
			return nil, buildError("match", ErrMismatchOperandTypes, args...)
		}

		regex, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		match := regex.MatchString(str)
		return match, nil
	}, nil)
	csl.RegisterFunction("find", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) != 2 {
			return nil, buildError("find", ErrNumOperands, args...)
		}
		var pattern, str string
		var ok bool
		if pattern, ok = args[0].(string); !ok {
			return nil, buildError("find", ErrMismatchOperandTypes, args...)
		}
		if str, ok = args[1].(string); !ok {
			return nil, buildError("find", ErrMismatchOperandTypes, args...)
		}

		regex, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		match := regex.FindString(str)
		return match, nil
	}, nil)
	csl.RegisterFunction("findall", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		// Accepts either exactly 2 or 3 arguments
		if len(args) < 2 || len(args) > 3 {
			return nil, buildError("findall", ErrNumOperands, args...)
		}

		var pattern, str string
		var ok bool
		var n int64 = -1
		if pattern, ok = args[0].(string); !ok {
			return nil, buildError("findall", ErrMismatchOperandTypes, args...)
		}
		if str, ok = args[1].(string); !ok {
			return nil, buildError("findall", ErrMismatchOperandTypes, args...)
		}
		if len(args) == 3 {
			// If a third argument was not provided, use -1 as default
			if n, ok = args[2].(int64); !ok {
				return nil, buildError("findall", ErrMismatchOperandTypes, args...)
			}
		}

		regex, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		matches := regex.FindAllString(str, int(n))

		// Convert to CSL list
		var l List
		for _, match := range matches {
			l = append(l, match)
		}

		return l, nil
	}, nil)
	csl.RegisterFunction("findallsubmatch", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		if len(args) < 2 || len(args) > 3 {
			return nil, buildError("findallsubmatch", ErrNumOperands, args...)
		}
		var pattern, str string
		var ok bool
		var n int64 = -1
		if pattern, ok = args[0].(string); !ok {
			return nil, buildError("findallsubmatch", ErrMismatchOperandTypes, args...)
		}
		if str, ok = args[1].(string); !ok {
			return nil, buildError("findallsubmatch", ErrMismatchOperandTypes, args...)
		}
		if len(args) == 3 {
			if n, ok = args[2].(int64); !ok {
				return nil, buildError("findallsubmatch", ErrMismatchOperandTypes, args...)
			}
		}

		regex, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		matches := regex.FindAllStringSubmatch(str, int(n))

		var l List
		for _, match := range matches {
			var subL List
			for _, submatch := range match {
				subL = append(subL, submatch)
			}
			l = append(l, subL)
		}

		return l, nil
	}, nil)
	csl.RegisterFunction("findindex", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		// Accepts exactly 2 arguments
		if len(args) != 2 {
			return nil, buildError("findindex", ErrNumOperands, args...)
		}

		var pattern, str string
		var ok bool
		if pattern, ok = args[0].(string); !ok {
			return nil, buildError("findindex", ErrMismatchOperandTypes, args...)
		}
		if str, ok = args[1].(string); !ok {
			return nil, buildError("findindex", ErrMismatchOperandTypes, args...)
		}

		regex, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		indexes := regex.FindStringIndex(str)

		// Convert to CSL list
		var l List
		for _, index := range indexes {
			l = append(l, index)
		}

		return l, nil
	}, nil)
	csl.RegisterFunction("findallindex", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		// Accepts either exactly 2 or 3 arguments
		if len(args) < 2 || len(args) > 3 {
			return nil, buildError("findallindex", ErrNumOperands, args...)
		}

		var pattern, str string
		var ok bool
		var n int64 = -1
		if pattern, ok = args[0].(string); !ok {
			return nil, buildError("findallindex", ErrMismatchOperandTypes, args...)
		}
		if str, ok = args[1].(string); !ok {
			return nil, buildError("findallindex", ErrMismatchOperandTypes, args...)
		}
		if len(args) == 3 {
			// If a third argument was not provided, use -1 as default
			if n, ok = args[2].(int64); !ok {
				return nil, buildError("findallindex", ErrMismatchOperandTypes, args...)
			}
		}

		regex, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		matches := regex.FindAllStringIndex(str, int(n))

		// Convert to CSL list
		var l List
		for _, match := range matches {
			var subL List
			for _, submatch := range match {
				subL = append(subL, submatch)
			}
			l = append(l, subL)
		}

		return l, nil
	}, nil)
	csl.RegisterFunction("findallsubmatchindex", false, true, func(env map[string]interface{}, args ...interface{}) (interface{}, error) {
		// Accepts either exactly 2 or 3 arguments
		if len(args) < 2 || len(args) > 3 {
			return nil, buildError("findallsubmatchindex", ErrNumOperands, args...)
		}

		var pattern, str string
		var ok bool
		var n int64 = -1
		if pattern, ok = args[0].(string); !ok {
			return nil, buildError("findallsubmatchindex", ErrMismatchOperandTypes, args...)
		}
		if str, ok = args[1].(string); !ok {
			return nil, buildError("findallsubmatchindex", ErrMismatchOperandTypes, args...)
		}
		if len(args) == 3 {
			// If a third argument was not provided, use -1 as default
			if n, ok = args[2].(int64); !ok {
				return nil, buildError("findallsubmatchindex", ErrMismatchOperandTypes, args...)
			}
		}

		regex, err := regexp.Compile(pattern)
		if err != nil {
			return nil, err
		}
		matches := regex.FindAllStringSubmatchIndex(str, int(n))

		// Convert to CSL list
		var l List
		for _, match := range matches {
			var subL List
			for _, submatch := range match {
				subL = append(subL, submatch)
			}
			l = append(l, subL)
		}

		return l, nil
	}, nil)
}

func checkValidFunctionName(name string) error {
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
	return nil
}

func buildError(fnName string, err error, args ...interface{}) error {
	return fmt.Errorf("error invoking %s with arguments %v: %w", fnName, args, err)
}
