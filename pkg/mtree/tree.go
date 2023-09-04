package mtree

import (
	"fmt"
	"io"
	"strings"
)

type Tree map[string]Tree
type BoxType int

const (
	Regular BoxType = iota
	Last
	AfterLast
	Between
)

func (tree Tree) Add(path string) {
	frags := strings.Split(path, "/")
	tree.add(frags)
}

func (tree Tree) add(frags []string) {
	if len(frags) == 0 {
		return
	}

	nextTree, ok := tree[frags[0]]
	if !ok {
		nextTree = Tree{}
		tree[frags[0]] = nextTree
	}

	nextTree.add(frags[1:])
}

func (tree Tree) Fprint(w io.Writer, root bool, padding string) {
	if tree == nil {
		return
	}

	index := 0
	for k, v := range tree {
		fmt.Fprintf(w, "%s%s\n", padding+getPadding(root, getBoxType(index, len(tree))), k)
		v.Fprint(w, false, padding+getPadding(root, getBoxTypeExternal(index, len(tree))))
		index++
	}
}

func (boxType BoxType) String() string {
	switch boxType {
	case Regular:
		return "\u251c\u2500"
	case Last:
		return "\u2514\u2500"
	case AfterLast:
		return "  "
	case Between:
		return "\u2502"
	default:
		panic("invalid box type")
	}
}

func getBoxType(index int, length int) BoxType {
	if index+1 == length {
		return Last
	} else if index+1 > length {
		return AfterLast
	}
	return Regular
}

func getBoxTypeExternal(index int, length int) BoxType {
	if index+1 == length {
		return AfterLast
	}
	return Between
}

func getPadding(root bool, boxType BoxType) string {
	if root {
		return ""
	}

	return boxType.String() + " "
}
