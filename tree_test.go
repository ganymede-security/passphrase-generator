package passphrasegenerator

import (
	"log"
	"testing"
)

func TestInsert(t *testing.T) {
	arr, _ := memoize(opts)
	log.Print(arr)
	var root tree

	root.sum = opts.withoutModifiers()

	for _, v := range arr {
		var root tree
		for j := 0; j < k+1; j++ {
			
		}
	}

	ch := make(chan int32)
	go sumBranch(&root, ch)

	for i := range ch {

		if i != 0 {
		log.Print(i)
		continue
		}
	}

	log.Print("Items in original array ", len(arr))
	log.Print("Number of nodes in tree ", root.root.getTreeNodeNum())
	log.Print("Number: ", root.root.PostOrder())
}
