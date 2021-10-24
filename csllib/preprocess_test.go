package csllib

import (
	"fmt"
	"testing"
)

func TestPreprocessScript1(t *testing.T) {
	csl := NewCSL(false)
	cslpm := NewCSLPackageManager(DefaultFetcher, false)
	err := csl.PreprocessScript(`(define x 7)`, "", cslpm)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPreprocessScript2(t *testing.T) {
	csl := NewCSL(false)
	cslpm := NewCSLPackageManager(DefaultFetcher, false)
	deps := csl.PreprocessScript(`(import "github.com/lsnow99/myscripts:main.csl" main)`, "", cslpm)
	fmt.Println(deps)
}

func TestCyclicDep(t *testing.T) {
	
}

func TestDFS(t *testing.T) {
	deps := DepGraph{}
	deps.Edges = make(map[*Node][]NodeReference)
	deps.Nodes = make([]*Node, 0)
	a := new(Node)
	b := new(Node)
	c := new(Node)
	a2b := NodeReference{
		RefNode: b,
	}
	b2c := NodeReference{
		RefNode: c,
	}
	c2a := NodeReference{
		RefNode: a,
	}
	deps.Nodes = append(deps.Nodes, a, b, c)
	deps.Edges[a] = append(deps.Edges[a], a2b)
	deps.Edges[b] = append(deps.Edges[b], b2c)
	deps.Edges[c] = append(deps.Edges[c], c2a)
	if !deps.DFS(a, b) {
		t.Fatalf("Expected to find a path from b to a but did not")
	}
}

