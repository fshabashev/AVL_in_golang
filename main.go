package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type TreeNode struct {
	val    int
	left   *TreeNode
	right  *TreeNode
	height int
}

type AVLTree struct {
	root *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	return &TreeNode{
		val:    val,
		height: 1,
	}
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func (t *AVLTree) Insert(val int) {
	t.root = t.insert(t.root, val)
}

func (t *AVLTree) insert(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return NewTreeNode(val)
	}
	if val < node.val {
		node.left = t.insert(node.left, val)
	} else {
		node.right = t.insert(node.right, val)
	}
	node.height = max(t.getHeight(node.left), t.getHeight(node.right)) + 1
	return t.balance(node)
}

func (t *AVLTree) balance(node *TreeNode) *TreeNode {
	if t.getBalance(node) > 1 {
		if t.getBalance(node.left) < 0 {
			node.left = t.leftRotate(node.left)
		}
		node = t.rightRotate(node)
	} else if t.getBalance(node) < -1 {
		if t.getBalance(node.right) > 0 {
			node.right = t.rightRotate(node.right)
		}
		node = t.leftRotate(node)
	}
	return node
}

func (t *AVLTree) leftRotate(node *TreeNode) *TreeNode {
	newRoot := node.right
	node.right = newRoot.left
	newRoot.left = node
	node.height = max(t.getHeight(node.left), t.getHeight(node.right)) + 1
	newRoot.height = max(t.getHeight(newRoot.left), t.getHeight(newRoot.right)) + 1
	return newRoot
}

func (t *AVLTree) rightRotate(node *TreeNode) *TreeNode {
	newRoot := node.left
	node.left = newRoot.right
	newRoot.right = node
	node.height = max(t.getHeight(node.left), t.getHeight(node.right)) + 1
	newRoot.height = max(t.getHeight(newRoot.left), t.getHeight(newRoot.right)) + 1
	return newRoot
}

func (t *AVLTree) getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return node.height
}

func (t *AVLTree) getBalance(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return t.getHeight(node.left) - t.getHeight(node.right)
}

func (t *AVLTree) InorderTraverse(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	t.InorderTraverse(node.left, res)
	*res = append(*res, node.val)
	t.InorderTraverse(node.right, res)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func TestAVLBySortingRandomNumbers() {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, 1000)
	for i := range arr {
		arr[i] = rand.Intn(1000)
	}
	tree := NewAVLTree()
	for _, num := range arr {
		tree.Insert(num)
	}
	res := []int{}
	tree.InorderTraverse(tree.root, &res)
	sort.Ints(arr)
	for i := range arr {
		if arr[i] != res[i] {
			fmt.Println("test failed")
			return
		}
	}
	fmt.Println("test passed")
}

func CheckAVLForBalance() {
	tree := NewAVLTree()
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(30)
	tree.Insert(40)
	tree.Insert(50)
	tree.Insert(25)
	if tree.root.val != 30 || tree.root.left.val != 20 || tree.root.right.val != 40 ||
		tree.root.left.left.val != 10 || tree.root.left.right.val != 25 || tree.root.right.right.val != 50 {
		fmt.Println("test failed")
		return
	}
	fmt.Println("test passed")
}

func main() {
	TestAVLBySortingRandomNumbers()
	CheckAVLForBalance()
}

