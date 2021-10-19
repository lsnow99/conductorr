package csl

import "fmt"

type DepGraph struct {
	Nodes []*Node
	Edges map[*Node][]NodeReference
}

type NodeReference struct {
	ImportedAs string
	RefNode    *Node
}

type Node struct {
	ImportPath string
	Script     string
}


func (d *DepGraph) addEdgeCheckCyclic(n *Node, edge NodeReference) (ok bool) {
	if !d.DFS(n, edge.RefNode) {
		ok = true
		d.Edges[n] = append(d.Edges[n], edge)
	}
	return
}

func (d *DepGraph) DFS(to *Node, from *Node) bool {
	children := d.Edges[from]
	if to == from {
		return true
	}
	for _, child := range children {
		if d.DFS(to, child.RefNode) {
			return true
		}
	}
	return false
}

func (n *Node) ResolveDependencies(deps DepGraph) error {
	sexprs, err := Parse(n.Script)
	if err != nil {
		return err
	}
	for _, sexpr := range sexprs {
		L := sexpr.L
		R := sexpr.R
		if L != nil && L.Type() == importAtom {
			if R != nil {
				RL := R.L
				if RL == nil {
					return fmt.Errorf("unexpected nil branch on s-expression tree")
				}
				RR := R.R
				if RR == nil {
					return fmt.Errorf("import requires two arguments, received one")
				}
				RRL := RR.L
				if RRL == nil {
					return fmt.Errorf("unexpected nil branch on s-expression tree")
				}
				importPath, ok := RL.stringVal()
				if !ok {
					return fmt.Errorf("first argument to import was not string literal")
				}
				fnName, ok := RRL.varName()
				if !ok {
					return fmt.Errorf("second argument to import was not a valid symbol")
				}
				// Check for existing dependency
				var foundNode *Node
				for _, node := range deps.Nodes {
					if node.ImportPath == importPath {
						foundNode = node
					}
				}
				if foundNode == nil {
					is, err := ParseImport(importPath)
					if err != nil {
						return err
					}
					script, err := is.Fetch()
					if err != nil {
						return err
					}
					foundNode = &Node{
						ImportPath: importPath,
						Script:     script,
					}
					deps.Nodes = append(deps.Nodes, foundNode)
					deps.Edges[n] = append(deps.Edges[n], NodeReference{
						ImportedAs: fnName,
						RefNode:    foundNode,
					})
				} else {
					if !deps.addEdgeCheckCyclic(n, NodeReference{ImportedAs: fnName, RefNode: foundNode}) {

					}
				}
			} else {
				return fmt.Errorf("import called with no arguments")
			}
		}
	}
	return nil
}

func PreprocessScript(script string) DepGraph {
	deps := DepGraph{
		Nodes: make([]*Node, 0, 1),
		Edges: make(map[*Node][]NodeReference),
	}
	rootNode := &Node{
		ImportPath: "",
		Script:     script,
	}
	deps.Nodes = append(deps.Nodes, rootNode)
	rootNode.ResolveDependencies(deps)
	return deps
}
