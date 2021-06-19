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

func checkResult(t *testing.T, expected interface{}, res interface{}, err error) {
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(res, expected) {
		t.Fatalf("got %v expected %v", res, expected)
	}
}
