package csl

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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
	boolToken
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
	defineAtom
	inAtom
	greaterAtom
	allGreaterAtom
	lessAtom
	allLessAtom
	greaterEqualAtom
	allGreaterEqualAtom
	lessEqualAtom
	allLessEqualAtom
	nthAtom
	eqAtom
	lengthAtom
	ifAtom
	boolAtom
	andAtom
	orAtom
	notAtom
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

func (a Atom) boolVal() (bool, bool) {
	x, ok := a.v.(bool)
	return x, (ok && a.typ == boolAtom)
}

func (a Atom) varName() (string, bool) {
	x, ok := a.v.(string)
	return x, (ok && a.typ == varAtom)
}

func (a Atom) Type() atomType {
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

var numberPattern = regexp.MustCompile(`^(((\-)?([0-9]+)(?i)(Gi|G|Mi|M|Ki|K)?))`)

// patterns returns the array of registered regex patterns
func patterns() []Pattern {
	return []Pattern{
		{whitespaceToken, regexp.MustCompile(`^\s+`)},
		{commentToken, regexp.MustCompile(`^;.*`)},
		{stringToken, regexp.MustCompile(`^("(\\.|[^"])*")`)},
		{boolToken, regexp.MustCompile(`^(true|false)`)},
		{numberToken, numberPattern},
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
func Parse(script string) ([]*SExpr, error) {
	var opened, startIndex int
	tokens := NewTokens(script)
	trees := make([]*SExpr, 0)
	for index, tok := range tokens {
		if tok.typ == openToken {
			opened++
		} else if tok.typ == closeToken {
			opened--
			if opened == 0 {
				tree, err := ParseTokens(tokens[startIndex:index+1])
				if err != nil {
					return nil, err
				}
				startIndex = index + 1
				trees = append(trees, tree)
			} else if opened < 0 {
				return nil, ErrSyntaxCloseParens
			}
		}
	}
	if opened != 0 {
		return nil, ErrSyntaxEOF
	}
	return trees, nil
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
	case boolToken:
		p, err := strconv.ParseBool(tok.val)
		if err != nil {
			return nil, err
		}
		sexpr.v = p
		sexpr.typ = boolAtom
	case numberToken:
		matches := numberPattern.FindAllStringSubmatch(tok.val, -1)
		if matches == nil {
			return nil, ErrParseNumber
		}
		if len(matches) != 1 {
			return nil, ErrParseNumber
		}
		if len(matches[0]) != 6 {
			return nil, ErrParseNumber
		}
		// Parse base number
		num, err := strconv.ParseInt(matches[0][4], 10, 64)
		if err != nil {
			err = ErrParseNumber
			break
		}
		// Parse multiplier
		switch strings.ToUpper(matches[0][5]) {
		case "G":
			num *= int64(1000000000)
		case "M":
			num *= int64(1000000)
		case "K":
			num *= int64(1000)
		case "GI":
			num *= (1 << 30)
		case "MI":
			num *= (1 << 20)
		case "KI":
			num *= (1 << 10)
		}
		// Flip sign if necessary
		if matches[0][3] == "-" {
			num *= -1
		}
		sexpr.v = num
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
		case "define":
			sexpr.typ = defineAtom
		case "in":
			sexpr.typ = inAtom
		case ">":
			sexpr.typ = greaterAtom
		case ">>":
			sexpr.typ = allGreaterAtom
		case "<":
			sexpr.typ = lessAtom
		case "<<":
			sexpr.typ = allLessAtom
		case ">=":
			sexpr.typ = greaterEqualAtom
		case ">>=":
			sexpr.typ = allGreaterEqualAtom
		case "<=":
			sexpr.typ = lessEqualAtom
		case "<<=":
			sexpr.typ = allLessEqualAtom
		case "nth":
			sexpr.typ = nthAtom
		case "eq":
			sexpr.typ = eqAtom
		case "len":
			sexpr.typ = lengthAtom
		case "if":
			sexpr.typ = ifAtom
		case "and":
			sexpr.typ = andAtom
		case "or":
			sexpr.typ = orAtom
		case "not":
			sexpr.typ = notAtom
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
