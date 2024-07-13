package SiteNavigator

import (
	"fmt"

	lls "github.com/PlayerR9/MyGoLib/ListLike/Stacker"
	uc "github.com/PlayerR9/MyGoLib/Units/common"
	us "github.com/PlayerR9/MyGoLib/Units/slice"
	"golang.org/x/net/html"
)

// FilterErrFunc is a function that returns an error if a condition is not met.
//
// Parameters:
//   - node: The node to check.
//
// Returns:
//   - error: An error if the condition is not met.
type FilterErrFunc func(node *html.Node) error

// FilterDataNode returns an FilterErrFunc that checks if the node has the specified data.
//
// Parameters:
//   - data: The data to check for.
//
// Returns:
//   - FilterErrFunc: The FilterErrFunc that checks if the node has the specified data.
func FilterDataNode(data string) FilterErrFunc {
	f := func(node *html.Node) error {
		fn := NewSearchCriteria(html.ElementNode).SetData(data).Build()

		ok := fn(node)
		if !ok {
			err := NewErrNoDataNodeFound(data)
			return err
		}

		return nil
	}

	return f
}

// FilterTextNode returns an FilterErrFunc that checks if the node is a text node.
//
// Parameters:
//   - checkFirstChild: If true, the function checks if the first child is a text node.
//     Otherwise, it checks if the node itself is a text node.
//
// Returns:
//   - FilterErrFunc: The FilterErrFunc that checks if the node is a text node.
func FilterTextNode(checkFirstChild bool) FilterErrFunc {
	var f FilterErrFunc

	if checkFirstChild {
		f = func(node *html.Node) error {
			ok := IsTextNodeSearch(node.FirstChild)
			if !ok {
				err := NewErrNoTextNodeFound(true)
				return err
			}

			return nil
		}
	} else {
		f = func(node *html.Node) error {
			ok := IsTextNodeSearch(node)
			if !ok {
				err := NewErrNoTextNodeFound(false)
				return err
			}

			return nil
		}
	}

	return f
}

// filterValidNodes is a helper function that filters the valid nodes from a list.
//
// Parameters:
//   - list: The list of nodes to filter.
//   - filter: The function to check if a node is valid.
//
// Returns:
//   - []*html.Node: The list of valid nodes.
//   - error: An error if no valid nodes are found.
//
// Behaviors:
//   - If no valid nodes are found, the function returns the first error encountered.
//   - If list is empty, the function returns nil, nil.
//
// Assumptions:
//   - list has at least one element.
//   - filter is not nil.
func filterValidNodes(list []*html.Node, filter FilterErrFunc) ([]*html.Node, error) {
	var el uc.ErrOrSol[*html.Node]

	for i := 0; i < len(list); i++ {
		err := filter(list[i])
		if err == nil {
			el.AddSol(list[i], 0)
		} else {
			el.AddErr(err, 0)
		}
	}

	ok := el.HasError()

	var sol []*html.Node
	var err error

	if ok {
		errs := el.GetErrors()
		err = errs[0]
	} else {
		sol = el.GetSolutions()
	}

	return sol, err
}

// FilterValidNodes filters the valid nodes from a list.
//
// Parameters:
//   - list: The list of nodes to filter.
//   - filters: The functions to check if a node is valid.
//
// Returns:
//   - []*html.Node: The list of valid nodes.
//   - error: An error if no valid nodes are found.
//
// Behaviors:
//   - If no valid nodes are found, the function returns the first error encountered.
//   - If list is empty or filters is empty, the function returns list, nil.
func FilterValidNodes(list []*html.Node, filters []FilterErrFunc) ([]*html.Node, error) {
	if len(filters) == 0 {
		return list, nil
	}

	var el uc.ErrOrSol[*html.Node]

	for level, filter := range filters {
		if len(list) == 0 {
			break
		}

		for i := 0; i < len(list); i++ {
			err := filter(list[i])
			if err == nil {
				el.AddSol(list[i], level)
			} else {
				el.AddErr(err, level)
			}
		}

		ok := el.HasError()

		if ok {
			errs := el.GetErrors()
			err := errs[0]

			return nil, err
		}

		list = el.GetSolutions()
	}

	return list, nil
}

// NodeListParser is a function that parses a list of nodes.
//
// Parameters:
//   - list: The list of nodes to parse.
//
// Returns:
//   - T: The parsed value.
//   - error: An error if the parsing fails.
type NodeListParser[T any] func(list []*html.Node) (T, error)

// CreateExtractor creates a NodeListParser from the given parameters.
//
// Parameters:
//   - parse: The function that parses the list of nodes.
//   - filters: The functions that filter the list of nodes.
//
// Returns:
//   - NodeListParser: The created NodeListParser.
//
// Behaviors:
//   - If parse is nil, the function returns a NodeListParser that returns
//     the error *errors.ErrInvalidParameter.
//   - Nil functions in filters are ignored.
func CreateExtractor[T any](parse NodeListParser[T], filters ...FilterErrFunc) NodeListParser[T] {
	if parse == nil {
		f := func(list []*html.Node) (T, error) {
			err := uc.NewErrNilParameter("parse")
			return *new(T), err
		}

		return f
	}

	filters = us.SliceFilter(filters, FilterNilFEFuncs)
	if len(filters) == 0 {
		return parse
	}

	f := func(list []*html.Node) (T, error) {
		var err error

		for _, extract := range filters {
			if len(list) == 0 {
				break
			}

			list, err = filterValidNodes(list, extract)
			if err != nil {
				return *new(T), err
			}
		}

		res, err := parse(list)
		if err != nil {
			return *new(T), err
		}

		return res, nil
	}

	return f
}

// ActionType is an enumeration of the different actions that can be performed on a node.
type ActionType int8

const (
	// OnlyDirectChildren is an action that extracts only the direct children of a node.
	OnlyDirectChildren ActionType = iota

	// DFSOne is an action that extracts only one node using depth-first search.
	DFSOne

	// BFSMany is an action that extracts multiple nodes using breadth-first search.
	BFSMany
)

// GTEFunc is a function that extracts nodes from a tree.
//
// Parameters:
//   - tree: The tree to extract nodes from.
//
// Returns:
//   - []*html.Node: The list of nodes extracted from the tree.
//   - error: An error if the extraction fails.
type GTEFunc func(tree *HtmlTree) ([]*html.Node, error)

// GenericTreeExtraction creates a GTEFunc from the given parameters.
//
// Parameters:
//   - search: The search criteria to use.
//   - action: The action to perform on the node.
//
// Returns:
//   - GTEFunc: The created GTEFunc.
//
// Behaviors:
//   - If search is nil, the function uses a nil filter.
//   - If action is not recognized, the function returns a GTEFunc that returns nil.
func GenericTreeExtraction(search *SearchCriteria, action ActionType) GTEFunc {
	var filter us.PredicateFilter[*html.Node]

	if search != nil {
		filter = search.Build()
	} else {
		filter = nil
	}

	var f GTEFunc

	switch action {
	case OnlyDirectChildren:
		f = func(tree *HtmlTree) ([]*html.Node, error) {
			sol, err := tree.ExtractSpecificNode(filter)

			return sol, err
		}
	case BFSMany:
		f = func(tree *HtmlTree) ([]*html.Node, error) {
			sol, err := tree.ExtractNodes(filter)
			return sol, err
		}
	case DFSOne:
		f = func(tree *HtmlTree) ([]*html.Node, error) {
			node, err := tree.ExtractContentFromDocument(filter)
			if err != nil {
				return nil, err
			}

			var sol []*html.Node

			if node != nil {
				sol = append(sol, node)
			}

			return sol, nil
		}
	default:
		f = func(tree *HtmlTree) ([]*html.Node, error) {
			return nil, fmt.Errorf("invalid action type: %v", action)
		}
	}

	return f
}

// CEWithSearch creates a NodeListParser from the given parameters.
//
// Parameters:
//   - search: The search criteria to use.
//   - action: The action to perform on the node.
//   - parse: The function that parses the list of nodes.
//   - filters: The functions that filter the list of nodes.
//
// Returns:
//   - NodeListParser: The created NodeListParser.
//
// Behaviors:
//   - If parse is nil, the function returns a NodeListParser that returns
//     the error *errors.ErrInvalidParameter.
//   - Nil functions in filters are ignored.
//   - Uses a Stack to traverse the tree and GTEFunc to extract nodes.
//   - It terminates as soon as a valid result is found.
func CEWithSearch[T any](search *SearchCriteria, action ActionType, parse NodeListParser[T], filters ...FilterErrFunc) NodeListParser[T] {
	if parse == nil {
		f := func(list []*html.Node) (T, error) {
			err := uc.NewErrNilParameter("parse")
			return *new(T), err
		}

		return f
	}

	filters = us.SliceFilter(filters, FilterNilFEFuncs)
	searchFunc := GenericTreeExtraction(search, action)

	f := func(list []*html.Node) (T, error) {
		if len(list) == 0 {
			err := NewErrNoNodesFound()
			return *new(T), err
		}

		S := lls.NewArrayStack(list...)

		var el uc.ErrOrSol[*html.Node]

		for {
			node, ok := S.Pop()
			if !ok {
				break
			}

			tree, err := NewHtmlTree(node)
			if err != nil {
				el.AddErr(err, 0)
				continue
			}

			newList, err := searchFunc(tree)
			if err != nil {
				el.AddErr(err, 1)
				continue
			}

			if len(newList) != 0 {
				newList, err = FilterValidNodes(newList, filters)
				if err != nil {
					el.AddErr(err, 1)
					continue
				}
			}

			res, err := parse(newList)
			if err == nil {
				return res, nil
			}

			el.AddErr(err, 1+len(filters))
		}

		errList := el.GetErrors()

		reason := NewErrNoNodesFound()
		err := uc.NewErrPossibleError(reason, errList[0])

		return *new(T), err
	}

	return f
}
