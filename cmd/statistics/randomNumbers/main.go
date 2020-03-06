// Don't judge me. I am not young and I need the money :D
package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sort"
)

func oneOutOf(max int) (int, error) {
	n := big.NewInt(
		int64(max),
	)
	randomInt, err := rand.Int(rand.Reader, n)
	if err != nil {
		return 0, err
	}

	result := int(randomInt.Int64()) + 1
	return result, nil
}

func outOf(max int, times int) ([]int, error) {
	resultMap := make(map[int]bool)

	for times > 0 {
		v, err := oneOutOf(max)
		if err != nil {
			return []int{}, err
		}

		if _, ok := resultMap[v]; !ok {
			resultMap[v] = true
			times -= 1
		}
	}

	resultList := make([]int, times)
	for k, _ := range resultMap {
		resultList = append(resultList, k)
	}

	sort.IntSlice(resultList).Sort()

	return resultList, nil
}

func main() {
	fiveOutOfFifty, err := outOf(50, 5)
	if err != nil {
		panic(err)
	}

	twoOutOfTen, err := outOf(10, 2)
	if err != nil {
		panic(err)
	}

	fmt.Println(fiveOutOfFifty, twoOutOfTen)
}
