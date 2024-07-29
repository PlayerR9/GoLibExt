package GoLibExt

import (
	_ "github.com/PlayerR9/go_generator/Generator"
	_ "github.com/PlayerR9/tree"
)

//go:generate go run github.com/PlayerR9/tree/cmd -name=TreeNode -fields=Data/*html.Node -o=SiteNavigator/treenode.go
