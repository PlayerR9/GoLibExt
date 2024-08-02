package GoLibExt

import (
	_ "github.com/PlayerR9/lib_units/generator"
	_ "github.com/PlayerR9/tree"
)

//go:generate go run github.com/PlayerR9/tree/cmd -name=TreeNode -fields=Data/*html.Node -o=Site_Navigator/treenode.go
