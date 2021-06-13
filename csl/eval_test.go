package csl

import "testing"

func TestEvalBasic1(t *testing.T) {
	expr, err := Parse(`
	;wtf dfas
	(* 2 (+ 3 (/ 12 6) (* 1 2)))
	`)
	if err != nil {
		t.Fatal(err)
	}
	res, err := Eval(expr, nil)
	if err != nil {
		t.Fatal(err)
	}
	if int64(14) != res {
		t.Logf("%v != 14", res)
	}
}
