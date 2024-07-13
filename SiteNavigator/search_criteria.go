package SiteNavigator

import (
	"slices"

	"golang.org/x/net/html"

	uc "github.com/PlayerR9/MyGoLib/Units/common"
	slext "github.com/PlayerR9/MyGoLib/Units/slice"
)

// AttrPair is a struct that encapsulates an attribute key-value pair and a filter function.
type AttrPair struct {
	// Attr is the attribute key to match.
	Attr string

	// FilterFunc is the filter function to apply to the attribute value.
	FilterFunc slext.PredicateFilter[string]
}

// NewAttrPair constructs a new AttrPair instance using the provided parameters.
//
// Parameters:
//   - attr: The attribute key to match.
//   - filter_func: The filter function to apply to the attribute value.
//
// Returns:
//   - *AttrPair: A new AttrPair instance. Nil if the filter function is nil.
func NewAttrPair(attr string, filter_func slext.PredicateFilter[string]) *AttrPair {
	if filter_func == nil {
		return nil
	}

	ap := &AttrPair{
		Attr:       attr,
		FilterFunc: filter_func,
	}
	return ap
}

// Match is a method of the AttrPair type that checks if the attribute key-value pair
// matches the provided attribute.
//
// Parameters:
//   - attr: The attribute key-value pair to match against.
//
// Returns:
//   - bool: True if the attribute key-value pair matches the provided attribute, false otherwise.
func (ap *AttrPair) Match(attr []html.Attribute) bool {
	f := func(a html.Attribute) bool {
		if a.Key != ap.Attr {
			return false
		}

		ok := ap.FilterFunc(a.Val)
		return ok
	}

	ok := slices.ContainsFunc(attr, f)
	return ok
}

// SearchCriteria is a struct that encapsulates the parameters for searching
// within an HTML node.
type SearchCriteria struct {
	// NodeType specifies the type of the HTML node to search for.
	NodeType html.NodeType

	// Data represents the data contained within the node.
	Data *string

	// Attrs is a slice of attribute key-value pairs to match.
	Attrs []*AttrPair
}

// NewSearchCriteria constructs a new SearchCriteria instance using the provided
// parameters.
//
// Parameters:
//   - node_type: The type of the HTML node to search for.
//
// Returns:
//   - *SearchCriteria: A new SearchCriteria instance.
func NewSearchCriteria(node_type html.NodeType) *SearchCriteria {
	sc := &SearchCriteria{
		NodeType: node_type,
	}
	return sc
}

// SetData sets the data field of the SearchCriteria instance.
//
// Parameters:
//   - data: The data to set in the SearchCriteria instance.
//
// Returns:
//   - *SearchCriteria: The SearchCriteria instance with the data field set.
func (sc *SearchCriteria) SetData(data string) *SearchCriteria {
	sc.Data = &data

	return sc
}

// AppendAttr is a method of the SearchCriteria type that appends an attribute key-value
// pair to the SearchCriteria instance.
//
// Parameters:
//   - key: The attribute key to match.
//   - val: The attribute value to match.
//
// Returns:
//   - *SearchCriteria: The SearchCriteria instance with the attribute key-value pair appended.
func (sc *SearchCriteria) AppendAttr(key string, val slext.PredicateFilter[string]) *SearchCriteria {
	if val == nil {
		val = func(s string) bool {
			return true
		}
	}

	p := NewAttrPair(key, val)
	uc.Assert(p != nil, "AppendAttr: NewAttrPair returned nil")

	sc.Attrs = append(sc.Attrs, p)

	return sc
}

// Build is a method of the SearchCriteria type that constructs a slext.PredicateFilter
// function using the search criteria.
//
// Returns:
//   - slext.PredicateFilter: A function that matches the search criteria.
func (sc *SearchCriteria) Build() slext.PredicateFilter[*html.Node] {
	attrsFunc := sc.Attrs
	nt := sc.NodeType

	var f slext.PredicateFilter[*html.Node]

	if sc.Data != nil {
		data := *sc.Data

		f = func(node *html.Node) bool {
			if node == nil {
				return false
			} else if node.Type != nt {
				return false
			} else if node.Data != data {
				return false
			}

			// Check if it matches the attribute
			for _, key := range attrsFunc {
				ok := key.Match(node.Attr)
				if !ok {
					return false
				}
			}

			return true
		}
	} else {
		f = func(node *html.Node) bool {
			if node == nil {
				return false
			} else if node.Type != nt {
				return false
			}

			for _, key := range attrsFunc {
				ok := key.Match(node.Attr)
				if !ok {
					return false
				}
			}

			return true
		}
	}

	return f
}
