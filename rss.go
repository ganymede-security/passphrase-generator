package passphrasegenerator

import "fmt"

// Returns true if there is a subset of the target value in the array, otherwise
// returns false
func isSubsetSum(arr []int32, targetValue int32) bool {
	dp := make([][]int32, len(arr)+1)
	memo := make([]int32, len(arr))
	counter := 0

	for i := 0; i < len(dp); i++ {
		col := make([]int32, targetValue+1)
		dp[i] = col
	}

	for i := int32(0); i < int32(len(dp)); i++ {
		for j := int32(0); j < int32(len(dp[i])); j++ {

			if i == 0 && j == 0 {
				dp[i][j] = 1
				continue
			}
			if i == 0 {
				continue
			}
			if j == 0 {
				dp[i][j] = 1
				counter++
				continue
			}
			if dp[i-1][j] == 1 {
				dp[i][j] = 1
				counter++
				continue
			}

			val := arr[i-1]
			if j >= val && dp[i-1][j-val] == 1 {
				memo = append(memo, val)
				dp[i][j] = 1
			}
		}
	}

	for _, v := range dp {

		if v[len(v)-1] == 1 {

			return true
		}
	}

	return false
}

func memoize(opts Options) ([]int32, error) {
	arr := make([]int32, 0)
	minWordLen := 3
	counter := int32(0)
	target := int32(opts.PhraseLength - opts.Numbers - opts.SpecialChars)

	for i := int32(0); i != target; {
		rn, err := getRandNum(minWordLen, opts.MaxWordLength)
		if err != nil {
			return nil, err
		}
		arr = append(arr, rn)
		counter += rn
		if counter == target {
			return arr, nil
		} else if counter > target && isSubsetSum(arr, target) {
			return arr, nil
		} else {
			continue
		}
	}
	err := fmt.Errorf("rss: Unable to generate random subset")

	return nil, err
}

// Naive algorithm for finding the subset of the array that adds to the target sum.
// Algorithm is O(2^n * n) since there are 2^n subsets to check. Memory is O(n)
// func inclusionExclusion(opts Options) /*(subset []int32, err error)*/ {
// 	// Create a binary tree in which each level corresponds to an input number.
// 	// left branch corresponds to inclusion and the right branch is for exclusion
// 	arr, err := memoize(opts)
// 	if err != nil {
// 		return
// 	}
// 	tree := newNode(arr[0])
// 	rl := int32(opts.PhraseLength - opts.Numbers - opts.SpecialChars)

// 	// Fill the tree with values, values that are greater than the phrase length
// 	// length are input on the left node.
// 	for _, v := range arr {
// 		for j:= 0; j > len(arr); j++ {
// 		switch {
// 		case v + tree.value == rl:
// 			tree.right = newNode(v)
// 			rl -= v
// 		case v + tree.value > rl:
// 			tree.left = newNode(v)
// 		case v + tree.value < rl:
// 			rl -= v
// 			tree.right = newNode(v)
// 			}
// 		}
// 	}
// 	//ret := make([]int32, len(arr))

// }

/*
set[]={3, 4, 5, 2}
target=6

0    1    2    3    4    5    6

0   T    F    F    F    F    F    F

3   T    F    F    T    F    F    F

4   T    F    F    T    T    F    F

5   T    F    F    T    T    T    F

2   T    F    T    T    T    T    T

*/
