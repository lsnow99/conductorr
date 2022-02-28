package csllib

import (
	"fmt"
	"path/filepath"
)

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

// DFS conducts a depth first search to and from given nodes, returning true
// if it is possible to reach the to node from the from node, and false
// otherwise.
func (d *DepGraph) DFS(to *Node, from *Node) bool {
	// Base case if we have reached our destination
	if to == from {
		return true
	}
	children := d.Edges[from]
	for _, child := range children {
		if d.DFS(to, child.RefNode) {
			// We only want to return true if there is a valid path from this
			// child. We do not want to return false yet, because it is
			// possible not all children have been checked.
			return true
		}
	}
	// All children have been checked recursively for a path to the to node,
	// and no such path exists.
	return false
}

// ResolveDependencies is a method on a node within a given dependency graph that resolves
// all scripts imported by the node recursively to build depdency graph. Each node receives
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
				parentGS, parentIsGS := parentImportableScript.(GitScript)
				parentFS, parentIsFS := parentImportableScript.(FileScript)
				curFS, curIsFS := curImportableScript.(FileScript)
				if parentIsGS && curIsFS {
					// If the currently-being-imported script is a file script and the parent
					// script is a git script, then treat the child script as a git script
					// that inherits the context from the parent script.
					is := GitScript{
						host: parentGS.host,
						repo: parentGS.repo,
						filePath: curFS.filePath,
						version: parentGS.version,
					}
					importPath = is.CanonicalizedImportPath()
				} else if parentIsFS && curIsFS && filepath.IsAbs(curFS.filePath) {
					// If the currently-being-imported script is a file script and the parent
					// script is a file script, AND if the current filepath is not absolute,
					// then rewrite the filepath to be absolute using the parent as the base
					// directory.
					parentDir := filepath.Dir(parentFS.filePath)
					curScriptFilepath := filepath.Join(parentDir, curFS.filePath)
					is := FileScript{
						filePath: curScriptFilepath,
					}
					importPath = is.CanonicalizedImportPath()
				}

				var foundNode *Node
				// Check for existing dependency
				for _, node := range deps.Nodes {
					if node.ImportPath == importPath {
						foundNode = node
					}
				}
				if foundNode == nil {
					// If we are adding a brand new node to our graph, we do not need to worry
					// about cyclic dependencies, since it is impossible for a cycle to form
					// when we add a new leaf.
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
					// Since we are referencing an existing node in our graph, we add it in a way
					// that checks to make sure no cycle has formed.
					if !deps.addEdgeCheckCyclic(n, NodeReference{ImportedAs: fnName, RefNode: foundNode}) {
						return fmt.Errorf("circular dependency detected! Importing %s from %s", importPath, n.ImportPath)
					}
				}
				depCSL := csl.CopyWithBuiltins()
				err = foundNode.ResolveDependencies(deps, depCSL, cslpm)
				if err != nil {
					return err
				}
				csl.RegisterFunction(fnName, true, false, nil, func(env map[string]interface{}, args []*SExpr, trace Trace) (interface{}, Trace) {
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
					res, tr := depCSL.Eval(sexprs, depEnv)
					return res, tr
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
	return rootNode.ResolveDependencies(deps, csl, cslpm)
}
