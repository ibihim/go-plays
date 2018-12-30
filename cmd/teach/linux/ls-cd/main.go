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

func main() {
	rand.Seed(time.Now().Unix())

	createTmpDir()

	prob := createDecreasingProbability(2)
	tree := createTree(prob)

	printTree(tree, "")
}

func createTmpDir() {
	buf := make([]byte, 10)
	if _, err := rand.Read(buf); err != nil {
		fmt.Println(err)
		return
	}

	enc := base64.URLEncoding.WithPadding(base64.NoPadding)
	dirName := enc.EncodeToString(buf)
	mainDir := fmt.Sprintf("%s/%s", os.TempDir(), dirName)
	if err := os.Mkdir(mainDir, 0777); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(mainDir, " directory created")
}
