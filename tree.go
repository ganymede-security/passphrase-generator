package passphrasegenerator

import (
	"log"
	"sync"
)

type node struct {
	left  *node
	value int32
	right *node
}

// tree is a node with levels
type tree struct {
	root *node
	// The value that is trying to be achieved
	sum int32

	level int
	// The key is the index of each item in the memoization array
	//lock sync.Mutex
}

func walk(t *tree, ch chan int32) {
	defer close(ch)
	var walkNode func(n *node)
	walkNode = func(n *node) {
		if n == nil {
			return
		}
		walkNode(n.left)
		ch<-n.value
		walkNode(n.right)
	}
	walkNode(t.root)
}

func sumBranch(t *tree, ch chan int32) {
	defer close(ch)
	var walkNode func(n *node)
	walkNode = func(n *node) {
		if n == nil {
			return
		}
		walkNode(n.left)
		ch<-n.value
	}
	walkNode(t.root)
}

func insertWithLevel(t *tree, ch chan int32) {
	defer close(ch)
	
}

// func (t *tree) insert(data int32) {
// 	t.lock.Lock()
// 	defer t.lock.Unlock()
// 	n := &node{
// 		left:  nil,
// 		value: data,
// 		right: nil,
// 	}
// 	if t.root == nil {
// 		t.root = n
// 	} else {
// 		insertNext(t.root, n)
// 	}
// }

// // helper function to include the number
// func insertNext(node1, node2 *node) {
// 	switch {
// 	case node1.right == nil:
// 		node1.right = node2
// 	case node1.left == nil:
// 		node1.left = node2
// 	case node1.right != nil:
// 		insertNext(node1.right, node2)
// 	case node1.left != nil:
// 		insertNext(node1.left, node2)
// 	default:
// 		log.Print("Fallthrough")
// 	}
// }

// func (n *node) getTreeNodeNum() int {
// 	if n == nil {
// 		return 0
// 	} else {
// 		return n.left.getTreeNodeNum() + n.right.getTreeNodeNum() + 1
// 	}
// }

// func (root *node) PostOrder() int32 {
// 	if root != nil {
// 	  // print left tree
// 	  //root.left.PostOrder()
// 	  // print right tree
// 	  //root.right.PostOrder()

// 	  // print root
// 	  return root.value + root.left.PostOrder() + root.right.PostOrder()
// 	}
// 	return 0
//   }

// Function dfsRoot should only be run on the root
// node, it maps the values of
// func (t *tree) dfsRoot(opts Options) {
// 	var wg sync.WaitGroup
// 	ctx := context.Background()
// 	subset := make([]int32, 0)

// 	ctx = context.WithValue(ctx, "expected", opts.withoutModifiers())
// 	ctx = context.WithValue(ctx, "remaining", opts.withoutModifiers())
// 	n := t.root

// 	if n == nil {
// 		return
// 	}

// 	wg.Add(1)
// 	go n.left.returnValue()

// 	wg.Add(1)

// }

// func (n *node) returnValue(ctx context.Context, subset []int32, target int32) (final []int32) {
// 	switch {
// 	case target == n.value:
// 		subset = append(subset, n.value)
// 		return subset
// 		ctx.Done()
// 	case target > n.value:
// 		target -= n.value
// 		subset = append(subset, n.value)
// 	}

// 	return n.value

// }
