package SiteNavigator

import (
	us "github.com/PlayerR9/MyGoLib/Units/slice"
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
