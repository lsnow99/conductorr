package csl

import (
	"fmt"
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
	checkResult(t, int64(14), res, trace)
}

func TestEvalBasic2(t *testing.T) {
	expr, err := Parse(`
	(* 1 1 1 4 3)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace := Eval(expr, nil)
	checkResult(t, int64(12), res, trace)
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
	checkResult(t, expected, res, trace)
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
		t.Fatal(trace)
	}
	expected := int64(10)
	if x, ok := env["x"]; ok {
		checkResult(t, expected, x, trace)
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
		checkResult(t, expected, x, trace)
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
		checkResult(t, expected, x, trace)
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
		checkResult(t, int64(17), x, trace)
	} else {
		t.Fatal("x not defined in environment")
	}

	if y, ok := env["y"]; ok {
		checkResult(t, int64(13), y, trace)
	} else {
		t.Fatal("x not defined in environment")
	}

	if z, ok := env["z"]; ok {
		checkResult(t, int64(17*13), z, trace)
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
		checkResult(t, expected, x, trace)
	} else {
		t.Fatal("x not defined in environment")
	}
}

func TestInList(t *testing.T) {
	expr, err := Parse(`
	(in 3 (4 8 1 9))
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace := Eval(expr, nil)
	checkResult(t, false, res, trace)

	expr, err = Parse(`
	(in 3 (4 8 1 3 9))
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace = Eval(expr, nil)
	checkResult(t, true, res, trace)

	expr, err = Parse(`
	(in "hello world" ("hi" "world"))
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace = Eval(expr, nil)
	checkResult(t, false, res, trace)

	expr, err = Parse(`
	(in "hello world" ("hello world"))
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace = Eval(expr, nil)
	checkResult(t, true, res, trace)
}

func TestInString(t *testing.T) {
	expr, err := Parse(`
	(in "hello" "world")
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace := Eval(expr, nil)
	checkResult(t, false, res, trace)

	expr, err = Parse(`
	(in "hello" "hello world")
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace = Eval(expr, nil)
	checkResult(t, true, res, trace)
}

func TestGreaterThan(t *testing.T) {
	expr, err := Parse(`
	(> 3 1 2)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace := Eval(expr, nil)
	checkResult(t, true, res, trace)

	expr, err = Parse(`
	(> 3 1 3)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace = Eval(expr, nil)
	checkResult(t, false, res, trace)
}

func TestAllGreaterThan(t *testing.T) {
	expr, err := Parse(`
	(>> 3 2 1)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace := Eval(expr, nil)
	checkResult(t, true, res, trace)

	expr, err = Parse(`
	(>> 3 1 2)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace = Eval(expr, nil)
	checkResult(t, false, res, trace)
}

func TestLessThan(t *testing.T) {
	expr, err := Parse(`
	(< 1 2 2)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace := Eval(expr, nil)
	checkResult(t, true, res, trace)

	expr, err = Parse(`
	(< 1 1 2)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace = Eval(expr, nil)
	checkResult(t, false, res, trace)
}

func TestAllLessThan(t *testing.T) {
	expr, err := Parse(`
	(<< 1 2 3)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace := Eval(expr, nil)
	checkResult(t, true, res, trace)

	expr, err = Parse(`
	(<< 1 1 2)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace = Eval(expr, nil)
	checkResult(t, false, res, trace)
}

func TestGreaterThanEqual(t *testing.T) {
	expr, err := Parse(`
	(>= 3 1 3)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace := Eval(expr, nil)
	checkResult(t, true, res, trace)

	expr, err = Parse(`
	(>= 3 1 4)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace = Eval(expr, nil)
	checkResult(t, false, res, trace)
}

func TestAllGreaterThanEqual(t *testing.T) {
	expr, err := Parse(`
	(>>= 3 2 2)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace := Eval(expr, nil)
	checkResult(t, true, res, trace)

	expr, err = Parse(`
	(>>= 3 2 3)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace = Eval(expr, nil)
	checkResult(t, false, res, trace)
}

func TestLessThanEqual(t *testing.T) {
	expr, err := Parse(`
	(<= 1 2 1)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace := Eval(expr, nil)
	checkResult(t, true, res, trace)

	expr, err = Parse(`
	(<= 2 2 1)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace = Eval(expr, nil)
	checkResult(t, false, res, trace)
}

func TestAllLessThanEqual(t *testing.T) {
	expr, err := Parse(`
	(<<= 1 2 2)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace := Eval(expr, nil)
	checkResult(t, true, res, trace)

	expr, err = Parse(`
	(<<= 1 2 1)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace = Eval(expr, nil)
	checkResult(t, false, res, trace)
}

func TestNthElem(t *testing.T) {
	expr, err := Parse(`
	(nth 1 (1 7 3 9))
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace := Eval(expr, nil)
	checkResult(t, int64(7), res, trace)

	expr, err = Parse(`
	(nth 0 (99))
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace = Eval(expr, nil)
	checkResult(t, int64(99), res, trace)
}

func TestEqInt(t *testing.T) {
	expr, err := Parse(`
	(eq 17 17)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace := Eval(expr, nil)
	checkResult(t, true, res, trace)

	expr, err = Parse(`
	(eq 17 39)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace = Eval(expr, nil)
	checkResult(t, false, res, trace)
}

func TestEqIntVar(t *testing.T) {
	env := make(map[string]interface{})
	env["x"] = int64(18)
	expr, err := Parse(`
	(eq x 18)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace := Eval(expr, env)
	checkResult(t, true, res, trace)

	expr, err = Parse(`
	(eq 17 x)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace = Eval(expr, env)
	checkResult(t, false, res, trace)
}

func TestEqList(t *testing.T) {
	expr, err := Parse(`
	(eq (1 2 4 3 25 6) (1 2 4 3 25 6))
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace := Eval(expr, nil)
	checkResult(t, true, res, trace)

	expr, err = Parse(`
	(eq (1 25 4 3 2 6) (1 2 4 3 25 6))
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace = Eval(expr, nil)
	checkResult(t, false, res, trace)
}

func TestEqStr(t *testing.T) {
	expr, err := Parse(`
	(eq "hello" "hello")
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace := Eval(expr, nil)
	checkResult(t, true, res, trace)

	expr, err = Parse(`
	(eq "hello" "world")
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace = Eval(expr, nil)
	checkResult(t, false, res, trace)
}

func TestNegativeInt(t *testing.T) {
	expr, err := Parse(`
	(-1)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace := Eval(expr, nil)
	checkResult(t, int64(-1), res, trace)
}

func TestGiMiKi(t *testing.T) {
	expr, err := Parse(`
	(-1Gi)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace := Eval(expr, nil)
	checkResult(t, int64(-1*(1<<30)), res, trace)

	expr, err = Parse(`
	(20Mi)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace = Eval(expr, nil)
	checkResult(t, int64(20*(1<<20)), res, trace)

	expr, err = Parse(`
	(17Ki)
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, trace = Eval(expr, nil)
	checkResult(t, int64(17*(1<<10)), res, trace)
}

func printTrace(trace Trace) {
	fmt.Println("Program execution steps: ")
	for _, step := range trace.ExprTree {
		fmt.Println(step)
	}
}

func checkResult(t *testing.T, expected interface{}, res interface{}, trace Trace) {
	if trace.err != nil {
		printTrace(trace)
		t.Fatal(trace.err)
	}
	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("got %v expected %v", res, expected)
	}
}
