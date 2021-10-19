package csl

import "fmt"

type DepGraph struct {
	nodes []*Node
	edges map[Node][]NodeReference
}

type NodeReference struct {
	importedAs string
	refNode    *Node
}

type Node struct {
	importPath string
	script     string
}


func (d *DepGraph) addEdgeCheckCyclic(n *Node, edge NodeReference) (ok bool) {
	if !d.dfs(n, edge.refNode) {
		ok = true
		d.edges[*n] = append(d.edges[*n], edge)
	}
	return
}

func (d *DepGraph) dfs(to *Node, from *Node) bool {
	children := d.edges[*from]
	if to == from {
		return true
	}
	for _, child := range children {
		if d.dfs(to, child.refNode) {
			return true
		}
	}
	return false
}

func (n *Node) ResolveDependencies(deps DepGraph) error {
	sexprs, err := Parse(n.script)
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
				for _, node := range deps.nodes {
					if node.importPath == importPath {
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
						importPath: importPath,
						script:     script,
					}
					deps.nodes = append(deps.nodes, foundNode)
					deps.edges[*n] = append(deps.edges[*n], NodeReference{
						importedAs: fnName,
						refNode:    foundNode,
					})
				} else {
					if !deps.addEdgeCheckCyclic(n, NodeReference{importedAs: fnName, refNode: foundNode}) {

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
		nodes: make([]*Node, 0, 1),
		edges: make(map[Node][]NodeReference),
	}
	rootNode := &Node{
		importPath: "",
		script:     script,
	}
	deps.nodes = append(deps.nodes, rootNode)
	rootNode.ResolveDependencies(deps)
	return deps
}
