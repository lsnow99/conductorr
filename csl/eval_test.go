package csl

import (
	"reflect"
	"testing"
)

func TestEvalBasic1(t *testing.T) {
	expr, err := Parse(`
	;wtf dfas
	(* 2 (+ 3 (/ 12 6) (* 1 2)))
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace := Eval(expr, nil)
	checkResult(t, int64(14), res, trace.err)
}

func TestEvalBasic2(t *testing.T) {
	expr, err := Parse(`
	(* 1 1 1 4 3)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace := Eval(expr, nil)
	checkResult(t, int64(12), res, trace.err)
}

func TestEvalList(t *testing.T) {
	expr, err := Parse(`
	(1 1 1 4 3)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace := Eval(expr, nil)
	expected := List{
		Elems: []interface{}{int64(1), int64(1), int64(1), int64(4), int64(3)},
	}
	checkResult(t, expected, res, trace.err)
}

func TestDefineVar(t *testing.T) {
	expr, err := Parse(`
	(define x 10)
	`)
	if err != nil {
		t.Fatal(err)
	}
	env := make(map[string]interface{})
	_, trace := Eval(expr, env)
	if trace.err != nil {
		t.Fatal(trace.err)
	}
	expected := int64(10)
	if x, ok := env["x"]; ok {
		checkResult(t, expected, x, trace.err)
	} else {
		t.Fatal("x not defined in environment")
	}
}

func TestDefineVarWithExpr(t *testing.T) {
	expr, err := Parse(`
	(define x (* 3 4))
	`)
	if err != nil {
		t.Fatal(err)
	}
	env := make(map[string]interface{})
	_, trace := Eval(expr, env)
	expected := int64(12)
	if x, ok := env["x"]; ok {
		checkResult(t, expected, x, trace.err)
	} else {
		t.Fatal("x not defined in environment")
	}
}

func TestDefineVarWithEnv(t *testing.T) {
	expr, err := Parse(`
	(define x y)
	`)
	if err != nil {
		t.Fatal(err)
	}
	env := make(map[string]interface{})
	expected := int64(17)
	env["y"] = expected
	_, trace := Eval(expr, env)
	if x, ok := env["x"]; ok {
		checkResult(t, expected, x, trace.err)
	} else {
		t.Fatal("x not defined in environment")
	}
}

func TestDefineMultVar(t *testing.T) {
	expr, err := Parse(`
	(define x 17)

	(define y 13)

	(define z (* x y))
	`)
	if err != nil {
		t.Fatal(err)
	}
	env := make(map[string]interface{})
	_, trace := Eval(expr, env)
	if x, ok := env["x"]; ok {
		checkResult(t, int64(17), x, trace.err)
	} else {
		t.Fatal("x not defined in environment")
	}

	if y, ok := env["y"]; ok {
		checkResult(t, int64(13), y, trace.err)
	} else {
		t.Fatal("x not defined in environment")
	}

	if z, ok := env["z"]; ok {
		checkResult(t, int64(17 * 13), z, trace.err)
	} else {
		t.Fatal("x not defined in environment")
	}
}

func TestDefineListVar(t *testing.T) {
	expr, err := Parse(`
	(define x (1 5 3 1 8))
	`)
	if err != nil {
		t.Fatal(err)
	}
	env := make(map[string]interface{})
	_, trace := Eval(expr, env)
	if x, ok := env["x"]; ok {
		expected := List{
			Elems: []interface{}{int64(1), int64(5), int64(3), int64(1), int64(8)},
		}
		checkResult(t, expected, x, trace.err)
	} else {
		t.Fatal("x not defined in environment")
	}
}

func checkResult(t *testing.T, expected interface{}, res interface{}, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("got %v expected %v", res, expected)
	}
}
