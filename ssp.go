package passphrasegenerator

import (
	"fmt"
)

func combinations(L []int32, r int32, target int32) [][]int32 {
	if r == 1 {
		temp := make([][]int32, 0)
		count := int32(0)
		for _, rr := range L {
			if rr + count < target {
				t := make([]int32, 0)
				t = append(t, rr)
				temp = append(temp, [][]int32{t}...)
			}
		}
		return temp
	} else {
		res := make([][]int32, 0)
		for i := 0; i < len(L); i++ {
			perms := make([]int32, 0)
			perms = append(perms, L[:i]...)
			for _, x := range combinations(perms, r-1, target) {
				count := int32(0)
				if L[i] + count < target {
				t := append(x, L[i])
				res = append(res, [][]int32{t}...)
				}
			}
		}
		return res
	}
}

func powerSet(L []int32, target int32) [][]int32 {
	res := make([][]int32, len(L))

	for i := int32(0); i <= int32(len(L)); i++ {
		x := combinations(L, i, target)
		res = append(res, x...)
	}
	return res
}

func SubSetSum(set []int32, target int32) (subset []int32, err error) {
	allSets := powerSet(set, target)
	for _, v := range allSets {
		if sumArr(v) == target {
			return v, nil
		}
	}
	err = fmt.Errorf("ssp: Error finding subset that equals given sum")
	return nil, err
}
