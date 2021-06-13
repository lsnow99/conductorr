package csl

import (
	"fmt"
	"regexp"
	"strconv"
)

type atomType uint8

type Atom struct {
	v   interface{}
	typ atomType
}

type SExpr struct {
	Atom
	L *SExpr
	R *SExpr
}

type Tokens []*Token

type tokenType uint8

type Token struct {
	typ tokenType
	val string
}

type Pattern struct {
	typ    tokenType
	regexp *regexp.Regexp
}

const (
	whitespaceToken tokenType = iota
	commentToken
	stringToken
	numberToken
	openToken
	closeToken
	symbolToken
	nilToken
)

const (
	stringAtom = iota + 1
	numberAtom
	varAtom
	addAtom
	subAtom
	multAtom
	divAtom
)

type CSLParserError struct {
	Message  string
	Internal bool
}

var (
	ErrSyntaxEOF = &CSLParserError{
		Message: "unexpected EOF while reading",
	}
	ErrSyntaxCloseParens = &CSLParserError{
		Message: "unexpected )",
	}
	ErrParserStrLen = &CSLParserError{
		Message:  "token type detected as string but value failed to parse",
		Internal: true,
	}
	ErrParseNumber = &CSLParserError{
		Message:  "token type detected as number but value failed to parse",
		Internal: true,
	}
)

func (c CSLParserError) Error() string {
	str := ""
	if c.Internal {
		str += "internal parser error: "
	}
	str += c.Message
	return str
}

func (a Atom) isAtom() bool {
	return a.typ > 0
}

func (a Atom) numberVal() (int64, bool) {
	x, ok := a.v.(int64)
	return x, (ok && a.typ == numberAtom)
}

func (a Atom) stringVal() (string, bool) {
	x, ok := a.v.(string)
	return x, (ok && a.typ == stringAtom)
}

func (a Atom) varName() (string, bool) {
	x, ok := a.v.(string)
	return x, (ok && a.typ == varAtom)
}

func (a Atom) Type() (atomType) {
	return a.typ
}

func (se *SExpr) String() string {
	level := 0
	exprs := se.getExprAtLevel(level)
	str := ""
	for exprs != nil && len(exprs) > 0 {
		str += fmt.Sprintf("\nLevel %d: ", level)
		for _, se := range exprs {
			leftStr, rightStr := "NIL", "NIL"
			if se.L != nil {
				leftStr = fmt.Sprintf("%v", se.L.Atom)
			}
			if se.R != nil {
				rightStr = fmt.Sprintf("%v", se.R.Atom)
			}
			str += fmt.Sprintf("(%v left:%v right:%v)", se.Atom, leftStr, rightStr)
		}
		level++
		exprs = se.getExprAtLevel(level)
	}
	return str
}

// patterns returns the array of registered regex patterns
func patterns() []Pattern {
	return []Pattern{
		{whitespaceToken, regexp.MustCompile(`^\s+`)},
		{commentToken, regexp.MustCompile(`^;.*`)},
		{stringToken, regexp.MustCompile(`^("(\\.|[^"])*")`)},
		{numberToken, regexp.MustCompile(`^((([0-9]+)?\.)?[0-9]+)`)},
		{openToken, regexp.MustCompile(`^(\()`)},
		{closeToken, regexp.MustCompile(`^(\))`)},
		{symbolToken, regexp.MustCompile(`^('|[^\s();]+)`)},
	}
}

/*
NewTokens takes a string for the program and performs regular-expression matching
to return a slice of parsed tokens. Code is adapted from 
[Jan Andersson](https://github.com/janne)'s 
[go-lisp](https://github.com/janne/go-lisp) project.
*/
func NewTokens(program string) (tokens Tokens) {
	for pos := 0; pos < len(program); {
		for _, pattern := range patterns() {
			if matches := pattern.regexp.FindStringSubmatch(program[pos:]); matches != nil {
				if len(matches) > 1 {
					tokens = append(tokens, &Token{pattern.typ, matches[1]})
				}
				pos = pos + len(matches[0])
				break
			}
		}
	}
	return
}

// getExprAtLevel debugging function to get all sexprs at a given depth
func (se *SExpr) getExprAtLevel(level int) []*SExpr {
	if se == nil {
		return nil
	}

	if level < 0 {
		return nil
	} else if level == 0 {
		return []*SExpr{se}
	} else if level > 0 {
		leftExprs := se.L.getExprAtLevel(level - 1)
		rightExprs := se.R.getExprAtLevel(level - 1)
		return append(leftExprs, rightExprs...)
	}

	return nil
}

/*
Parse is the main client function to parse a script into an S-Expression
tree. Returns a non-nil error if a parsing error is encountered.
*/
func Parse(script string) (*SExpr, error) {
	tokens := NewTokens(script)
	return ParseTokens(tokens)
}

/*
ParseAtomicToken given an atomic token, (a symbol in Lisp), parse it into
an atomic S-Expression that is ready to be evaluated.
*/
func ParseAtomicToken(tok *Token) (*SExpr, error) {
	sexpr := &SExpr{}
	var err error
	switch tok.typ {
	case stringToken:
		if len(tok.val) < 2 {
			err = ErrParserStrLen
			break
		}
		sexpr.v = tok.val[1 : len(tok.val)-1]
		sexpr.typ = stringAtom
	case numberToken:
		sexpr.v, err = strconv.ParseInt(tok.val, 10, 64)
		if err != nil {
			err = ErrParseNumber
			break
		}
		sexpr.typ = numberAtom
	case symbolToken:
		switch tok.val {
		case "+":
			sexpr.typ = addAtom
		case "-":
			sexpr.typ = subAtom
		case "*":
			sexpr.typ = multAtom
		case "/":
			sexpr.typ = divAtom
		default:
			sexpr.typ = varAtom
			sexpr.v = tok.val
		}
	}
	return sexpr, err
}

/*
ParseTokens main parsing logic that builds an S-Expression tree from a slice
of parsed tokens.
*/
func ParseTokens(tokens Tokens) (*SExpr, error) {
	root := &SExpr{}
	var err error
	for index := 0; index < len(tokens); index++ {
		tok := tokens[index]
		switch tok.typ {
		case openToken:
			elem := tokens[index+1:].NextElem()
			curRoot := root
			for len(elem) > 0 && elem[0].typ != closeToken {
				index += len(elem)
				if len(elem) > 1 {
					curRoot.L, err = ParseTokens(elem)
				} else if len(elem) == 1 {
					curRoot.L, err = ParseAtomicToken(elem[0])
				}
				if err != nil {
					return nil, err
				}
				elem = tokens[index+1:].NextElem()
				if len(elem) > 0 && elem[0].typ != closeToken {
					newRoot := &SExpr{}
					curRoot.R = newRoot
					curRoot = newRoot
				} else {
					break
				}
			}
		case closeToken:
		default:
			root.L, err = ParseAtomicToken(tok)
		}
	}
	return root, err
}

/*
NextElem takes a slice of tokens and returns the next "element", defined to be
the next element in a list. For example, the second element in the expression
```
(1 (+ 3 4) 7)
```
is `(+ 3 4)`
*/
func (tokens Tokens) NextElem() Tokens {
	elem := Tokens{}
	opened := 0
	for i := 0; i < len(tokens); i++ {
		tok := tokens[i]
		elem = append(elem, tok)
		if tok.typ == openToken {
			opened++
		}
		if tok.typ == closeToken {
			opened--
		}
		if opened <= 0 {
			break
		}
	}
	return elem
}

/*
Expand takes a standard format lisp expression and returns its expanded form.
This function is not in use, but was helpful in building the main parser logic.
*/
func Expand(tokens Tokens) Tokens {
	L := Tokens{}
	for index := 0; index < len(tokens); index++ {
		tok := tokens[index]
		switch tok.typ {
		case openToken:
			elem := tokens[index+1:].NextElem()
			opened := 0
			for len(elem) > 0 && elem[0].typ != closeToken {
				index += len(elem)
				if len(elem) > 1 {
					elem = Expand(elem)
				}
				L = append(L, &Token{
					typ: openToken,
					val: "(",
				})
				opened++
				L = append(L, elem...)
				elem = tokens[index+1:].NextElem()
			}
			L = append(L,
				&Token{
					typ: nilToken,
					val: "NIL",
				})
			for i := 0; i < opened; i++ {
				L = append(L,
					&Token{
						typ: closeToken,
						val: ")",
					})
			}
		case closeToken:
		default:
			L = append(L, tok)
		}
	}
	return L
}
