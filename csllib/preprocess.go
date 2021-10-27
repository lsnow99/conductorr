package csllib

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

const MaximumImports = 100

// addEdgeCheckCyclic attempts to add the given edge from node `n` to `edge.RefNode`.
// Before adding the edge, a depth-first search is performed from the node referenced
// by `edge`, to the node `n`. If it is possible to reach node `n` from `edge.RefNode`,
// then adding the edge would cause a cycle in the graph. The idea is that so long as
// the graph already has no cycles, this will detect any new ones. The edge is added
// so long as the check passes, and a boolean signalling success is returned.
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

func (n *Node) ResolveDependencies(deps DepGraph, csl *CSL, cslpm *CSLPackageManager) error {
	if len(deps.Nodes) > MaximumImports {
		return fmt.Errorf("number of dependencies exceeds maximum of %d", MaximumImports)
	}
	sexprs, err := csl.Parse(n.Script)
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
				parentImportableScript, err := cslpm.parseImport(n.ImportPath)
				if err != nil {
					return err
				}
				curImportableScript, err := cslpm.parseImport(importPath)
				if err != nil {
					return err
				}
				parentGS, pOK := parentImportableScript.(GitScript)
				curFS, cOK := curImportableScript.(FileScript)
				if pOK && cOK {
					is := GitScript{
						host: parentGS.host,
						repo: parentGS.repo,
						filePath: curFS.filePath,
						version: parentGS.version,
					}
					importPath = is.CanonicalizedImportPath()
				}
				// Check for existing dependency
				var foundNode *Node
				for _, node := range deps.Nodes {
					if node.ImportPath == importPath {
						foundNode = node
					}
				}
				if foundNode == nil {
					script, err := cslpm.Resolve(importPath)
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
						return fmt.Errorf("circular dependency detected! Importing %s from %s", importPath, n.ImportPath)
					}
				}
				depCSL := NewCSL(true)
				foundNode.ResolveDependencies(deps, depCSL, cslpm)
				csl.RegisterFunction(fnName, true, nil, func(env map[string]interface{}, args []*SExpr, trace Trace) (interface{}, Trace) {
					argList := List{}
					for _, arg := range args {
						res, trace := csl.EvalSExpr(arg, env, trace)
						if trace.Err != nil {
							return nil, trace
						}
						argList = append(argList, res)
					}
					sexprs, err := depCSL.Parse(foundNode.Script)
					if err != nil {
						trace.Err = err
						return nil, trace
					}
					depEnv := make(map[string]interface{})
					depEnv["args"] = argList
					return depCSL.Eval(sexprs, depEnv)
				})
			} else {
				return fmt.Errorf("import called with no arguments")
			}
		}
	}
	return nil
}

func (csl *CSL) PreprocessScript(script string, rootImportPath string, cslpm *CSLPackageManager) error {
	deps := DepGraph{
		Nodes: make([]*Node, 0, 1),
		Edges: make(map[*Node][]NodeReference),
	}
	rootNode := &Node{
		ImportPath: rootImportPath,
		Script:     script,
	}
	deps.Nodes = append(deps.Nodes, rootNode)
	// TODO option
	return rootNode.ResolveDependencies(deps, csl, cslpm)
}
