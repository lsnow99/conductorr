package csl

import "testing"

func TestExpandNoNesting(t *testing.T) {
	script := `(1 2 3)`
	expected := `( 1 ( 2 ( 3 NIL ) ) )`
	t.Logf("Expecting: %s", expected)
	t.Log(Expand(NewTokens(script)))
}

func TestExpandOneNested(t *testing.T) {
	script := `(1 (2 3))`
	expected := `( 1 ( 2 3 NIL) NIL )`
	t.Logf("Expecting: %s", expected)
	t.Log(Expand(NewTokens(script)))
}

func TestNextElemBasic(t *testing.T) {
	script := `(1 (2 3))`
	t.Log(NewTokens(script).NextElem())
}

func TestExpandComplex1(t *testing.T) {
	script := `(((1 0) 3) 4)`
	expected := `( ( ( 1 0 NIL ) 3 NIL ) 4 NIL )`
	t.Logf("Expecting: %s", expected)
	t.Log(Expand(NewTokens(script)))
}

func TestExpandComplex2(t *testing.T) {
	script := `(+ 1 4 5 (3 4 (1 0)) 8)`
	expected := `( + ( ( ( 1 ( 4 ( 5 ( ( 3 ( 4 ( 1 ( 0 NIL ) ) ) NIL ) ) ) ) ) 8 ) NIL ) )`
	t.Logf("Expecting: %s", expected)
	t.Log(Expand(NewTokens(script)))
}

func TestExpandComplex3(t *testing.T) {
	script := `(- (+ 1 2) 3)`
	expected := `( - ( ( + ( 1 ( 2 NIL ) ) ) ( 3 NIL ) ) )`
	t.Logf("Expecting: %s", expected)
	t.Log(Expand(NewTokens(script)))
}

func TestParseBasic1(t *testing.T) {
	csl := NewCSL(true)
	script := `(1 2 3)`
	t.Log(csl.Parse(script))
}

func TestParseBasic2(t *testing.T) {
	csl := NewCSL(true)
	script := `(1 (2 3))`
	t.Log(csl.Parse(script))
}

func TestParseBasic3(t *testing.T) {
	csl := NewCSL(true)
	script := `(((1 0) 3) 4)`
	t.Log(csl.Parse(script))
}

func TestParseArith(t *testing.T) {
	csl := NewCSL(true)
	script := `(* 2 (+ 3 (/ 12 6) (* 1 2)))`
	t.Log(csl.Parse(script))
}

func TestParseImport(t *testing.T) {
	csl := NewCSL(true)
	script := `(import "github.com/lsnow99/csl/myscript.csl")`
	sexprs, err := csl.Parse(script)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(sexprs)
}