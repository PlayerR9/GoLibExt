package site_navigator

import (
	"golang.org/x/net/html"

	uc "github.com/PlayerR9/lib_units/common"
	luint "github.com/PlayerR9/lib_units/ints"
	slext "github.com/PlayerR9/lib_units/slices"
	tr "github.com/PlayerR9/tree/tree"
)

var (
	// IsTextNodeSearch is a search criteria that matches text nodes.
	IsTextNodeSearch slext.PredicateFilter[*html.Node]

	// GetChildrenFunc is a function that returns the children of an HTML node.
	GetChildrenFunc tr.NextsFunc[*TreeNode]
)

func init() {
	IsTextNodeSearch = NewSearchCriteria(html.TextNode).Build()

	GetChildrenFunc = func(n *TreeNode, info tr.Infoer) ([]*TreeNode, error) {
		if n == nil {
			err := uc.NewErrNilParameter("n")
			return nil, err
		}

		if n.Data == nil {
			err := uc.NewErrNilParameter("n.Data")
			return nil, err
		}

		var children []*TreeNode

		for c := n.Data.FirstChild; c != nil; c = c.NextSibling {
			new_n := NewTreeNode(c)

			children = append(children, new_n)
		}

		return children, nil
	}
}

// GetDirectChildren returns a slice of the direct children of the provided node.
//
// Parameters:
//   - node: The HTML node to extract the children from.
//
// Returns:
//   - []*html.Node: A slice containing the direct children of the node.
func GetDirectChildren(node *html.Node) []*html.Node {
	if node == nil {
		return nil
	}

	var children []*html.Node
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		children = append(children, c)
	}
	return children
}

// HtmlTree is a struct that represents an HTML tree.
type HtmlTree struct {
	// The tree constructed from the HTML node.
	tree *tr.Tree[*TreeNode]
}

// NewHtmlTree constructs a tree from an HTML node.
//
// Parameters:
//   - root: The root HTML node.
//
// Returns:
//   - *HtmlTree: The tree constructed from the HTML node.
//   - error: An error if the tree construction fails.
//
// Errors:
//   - *uc.ErrNilValue: If any html.Node is nil.
func NewHtmlTree(root *html.Node) (*HtmlTree, error) {
	var builder tr.Builder[*TreeNode]

	builder.SetNextFunc(GetChildrenFunc)

	n := NewTreeNode(root)

	tree, err := builder.Build(n)
	if err != nil {
		return nil, err
	}

	ht := &HtmlTree{
		tree: tree,
	}

	return ht, nil
}

// ExtractSpecificNode finds all nodes that match the given search criteria and
// that are direct children of the provided node.
//
// Parameters:
//   - matchFun: The search criteria to apply to each node.
//
// Returns:
//   - []*html.Node: A slice containing all nodes that match the search criteria.
//   - error: An error if the search fails.
//
// Behavior:
//   - If no criteria is provided, then any node will match.
func (t *HtmlTree) ExtractSpecificNode(matchFun slext.PredicateFilter[*html.Node]) ([]*html.Node, error) {
	if matchFun == nil {
		panic("Case not handled: matchFun is nil")
	}

	children := t.tree.GetDirectChildren()
	if len(children) == 0 {
		return nil, nil
	}

	S := make([]*html.Node, 0, len(children))

	for _, child := range children {
		S = append(S, child.Data)
	}

	S = slext.SliceFilter(S, matchFun)
	return S, nil
}

// MatchNodes performs a breadth-first search on an HTML section returning a
// slice of nodes that match the provided search criteria.
//
// Parameters:
//   - matchFun: The search criteria to apply to each node.
//
// Returns:
//   - []*html.Node: A slice containing all nodes that match the search criteria.
//
// Behavior:
//   - It does not search the children of the nodes that match the criteria.
//   - If no criteria is provided, then the first node will match.
func (t *HtmlTree) MatchNodes(matchFun slext.PredicateFilter[*html.Node]) ([]*html.Node, error) {
	if matchFun == nil {
		panic("Case not handled: matchFun is nil")
	}

	var solution []*html.Node

	f := func(node *TreeNode, info tr.Infoer) (bool, error) {
		if node == nil {
			err := uc.NewErrNilParameter("node")
			return false, err
		}

		ok := matchFun(node.Data)
		if !ok {
			return true, nil
		}

		solution = append(solution, node.Data)
		return false, nil
	}

	err := tr.BFS(t.tree, nil, f)
	if err != nil {
		return nil, err
	}

	return solution, nil
}

// ExtractContentFromDocument performs a depth-first search on an HTML document,
// finding the first node that matches the provided search criteria.
//
// Parameters:
//   - matchFun: The search criteria to apply to each node.
//
// Returns:
//   - *html.Node: The first node that matches the search criteria, nil if no
//     matching node is found.
func (t *HtmlTree) ExtractContentFromDocument(matchFun slext.PredicateFilter[*html.Node]) (*html.Node, error) {
	if matchFun == nil {
		panic("Case not handled: matchFun is nil")
	}

	var solution *html.Node

	f := func(node *TreeNode, info tr.Infoer) (bool, error) {
		if node == nil {
			err := uc.NewErrNilParameter("node")
			return false, err
		}

		ok := matchFun(node.Data)
		if !ok {
			return true, nil
		}

		solution = node.Data
		return false, nil
	}

	err := tr.DFS(t.tree, nil, f)
	if err != nil {
		return nil, err
	}

	return solution, nil
}

// ExtractNodes performs a breadth-first search on an HTML section returning a
// slice of nodes that match the provided search criteria.
//
// Parameters:
//   - criterias: A list of search criteria to apply to each node.
//
// Returns:
//   - []*html.Node: A slice containing all nodes that match the search criteria.
//
// Behavior:
//   - If no criteria is provided, then any node will match.
func (t *HtmlTree) ExtractNodes(criterias ...slext.PredicateFilter[*html.Node]) ([]*html.Node, error) {
	criterias = slext.FilterNilPredicates(criterias)
	if len(criterias) == 0 {
		return nil, nil
	}

	todo := []*HtmlTree{t}

	for i, criteria := range criterias {
		var new_todo []*html.Node

		for _, tree := range todo {
			result, err := tree.MatchNodes(criteria)
			if err != nil {
				err := luint.NewErrWhileAt("applying", i+1, "criteria", err)
				return nil, err
			}

			if len(result) != 0 {
				new_todo = append(new_todo, result...)
			}
		}

		if len(new_todo) == 0 {
			return nil, nil
		}

		for i, node := range new_todo {
			new_tree, err := NewHtmlTree(node)
			if err != nil {
				err := luint.NewErrWhileAt("adding", i+1, "tree", err)
				return nil, err
			}

			todo = append(todo, new_tree)
		}
	}

	var solution []*html.Node

	for _, t := range todo {
		root := t.tree.Root()

		solution = append(solution, root.Data)
	}

	return solution, nil
}
