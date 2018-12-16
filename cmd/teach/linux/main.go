// Dirty code. "Business value"-driven.
// Value: teach children linux
package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// i18n
const (
	left  = "links"
	mid   = "mitte"
	right = "rechts"
)

type node struct {
	left, mid, right *node
	empty            bool
	value            bool
}

func createTree(maxDepth int, probability int, emptyRef []*node) *node {
	n := node{}
	nextP := probability + 1

	if maxDepth <= probability {
		fmt.Println(" maxDepth <= probability")
		n.empty = true
		emptyRef = append(emptyRef, &n)

		return &n
	}

	// left path
	if probability < rand.Intn(maxDepth) {
		n.left = createTree(maxDepth, nextP, emptyRef)
	} else {
		n.left = &node{empty: true}
		emptyRef = append(emptyRef, n.left)
		fmt.Printf("%+v\n", emptyRef)
	}

	// mid path
	if probability < rand.Intn(maxDepth) {
		n.mid = createTree(maxDepth, nextP, emptyRef)
	} else {
		n.mid = &node{empty: true}
		emptyRef = append(emptyRef, n.left)
	}

	// right path
	if probability < rand.Intn(maxDepth) {
		n.right = createTree(maxDepth, nextP, emptyRef)
	} else {
		n.right = &node{empty: true}
		emptyRef = append(emptyRef, n.left)
	}

	return &n
}

func printTree(n *node, indent string) {
	if n.left != nil {
		fmt.Println(indent, "left")
		printTree(n.left, fmt.Sprintf("%s%s", indent, "\t"))
	}

	if n.mid != nil {
		fmt.Println(indent, "mid")
		printTree(n.mid, fmt.Sprintf("%s%s", indent, "\t"))
	}

	if n.right != nil {
		fmt.Println(indent, "right")
		printTree(n.right, fmt.Sprintf("%s%s", indent, "\t"))
	}
}

func createDirs(n *node) {

}

func main() {
	maxDepth := 3

	rand.Seed(time.Now().Unix())

	b := make([]byte, 10)
	if _, err := rand.Read(b); err != nil {
		fmt.Println(err)
		return
	}

	encoding := base64.URLEncoding.WithPadding(base64.NoPadding)
	dirName := encoding.EncodeToString(b)
	mainDir := fmt.Sprintf("%s/%s", os.TempDir(), dirName)

	if err := os.Mkdir(mainDir, 0777); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(mainDir, " directory created")

	endNodes := make([]*node, 0)
	_ = createTree(maxDepth, 0, endNodes)

	//printTree(tree, "")

	for _, endNode := range endNodes {
		fmt.Printf("%+v\n", endNode)
	}
}
