package site_navigator

import (
	us "github.com/PlayerR9/lib_units/slices"
)

var (
	// FilterNilFEFuncs is a predicate filter that filters out nil FilterErrFuncs.
	FilterNilFEFuncs us.PredicateFilter[FilterErrFunc]
)

func init() {
	FilterNilFEFuncs = func(fef FilterErrFunc) bool {
		return fef != nil
	}
}
