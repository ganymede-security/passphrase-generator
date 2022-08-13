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

// Oracle for ssp
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
