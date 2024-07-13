package SiteNavigator

import "strings"

// ErrNoNodesFound is an error that is returned when no nodes are found.
type ErrNoNodesFound struct{}

// Error implements the error interface.
//
// It returns the error message: "no nodes found".
func (e *ErrNoNodesFound) Error() string {
	return "no nodes found"
}

// NewErrNoNodesFound creates a new ErrNoNodesFound error.
//
// Returns:
//   - *ErrNoNodesFound: The new error.
func NewErrNoNodesFound() *ErrNoNodesFound {
	e := &ErrNoNodesFound{}
	return e
}

// ErrNoDataNodeFound is an error that is returned when no data nodes are found.
type ErrNoDataNodeFound struct {
	// Data is the data that was not found.
	Data string
}

// Error implements the error interface.
//
// It returns the error message: "no <data> tags found".
func (e *ErrNoDataNodeFound) Error() string {
	var builder strings.Builder

	builder.WriteString("no <")
	builder.WriteString(e.Data)
	builder.WriteString("> tags found")

	str := builder.String()
	return str
}

// NewErrNoDataNodeFound creates a new ErrNoDataNodeFound error.
//
// Parameters:
//   - data: The data that was not found.
//
// Returns:
//   - *ErrNoDataNodeFound: The new error.
func NewErrNoDataNodeFound(data string) *ErrNoDataNodeFound {
	e := &ErrNoDataNodeFound{
		Data: data,
	}
	return e
}

// ErrNoTextNodeFound is an error that is returned when no text nodes are found.
type ErrNoTextNodeFound struct {
	// IsFirstChild is whether the first child is not a text node.
	IsFirstChild bool
}

// Error implements the error interface.
//
// It returns the error message: "node is not a text node".
// However, if IsFirstChild is true, it returns the error message: "first child is not a text node".
func (e *ErrNoTextNodeFound) Error() string {
	if e.IsFirstChild {
		return "first child is not a text node"
	} else {
		return "node is not a text node"
	}
}

// NewErrNoTextNodeFound creates a new ErrNoTextNodeFound error.
//
// Parameters:
//   - isFirstChild: Whether the first child is not a text node.
//
// Returns:
//   - *ErrNoTextNodeFound: The new error.
func NewErrNoTextNodeFound(isFirstChild bool) *ErrNoTextNodeFound {
	e := &ErrNoTextNodeFound{
		IsFirstChild: isFirstChild,
	}
	return e
}
