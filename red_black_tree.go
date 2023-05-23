package main

import "fmt"

type Color bool

const (
	Red   = false
	Black = true
)

type Node struct {
	color  Color
	value  int
	left   *Node
	right  *Node
	parent *Node
}

type RBTree struct {
	root *Node
}

func NewNode(value int) *Node {
	return &Node{color: Red, value: value}
}

func NewTree() *RBTree {
	return &RBTree{}
}

func (tree *RBTree) Insert(value int) {
	if tree.root == nil {
		tree.root = NewNode(value)
		tree.root.color = Black
	} else {
		tree.insert(NewNode(value))
	}
}

func (tree *RBTree) insert(node *Node) {
	tree.binaryInsert(node)
	node.color = Red
	for node != tree.root && node.parent.color == Red {
		if node.parent == node.parent.parent.left {
			if node.parent.parent.right != nil && node.parent.parent.right.color == Red {
				node.parent.color = Black
				node.parent.parent.right.color = Black
				node.parent.parent.color = Red
				node = node.parent.parent
			} else {
				if node == node.parent.right {
					node = node.parent
					tree.rotateLeft(node)
				}
				node.parent.color = Black
				node.parent.parent.color = Red
				tree.rotateRight(node.parent.parent)
			}
		} else {
			if node.parent.parent.left != nil && node.parent.parent.left.color == Red {
				node.parent.color = Black
				node.parent.parent.left.color = Black
				node.parent.parent.color = Red
				node = node.parent.parent
			} else {
				if node == node.parent.left {
					node = node.parent
					tree.rotateRight(node)
				}
				node.parent.color = Black
				node.parent.parent.color = Red
				tree.rotateLeft(node.parent.parent)
			}
		}
	}
	tree.root.color = Black
}

func (tree *RBTree) binaryInsert(node *Node) {
	current := tree.root
	for {
		if node.value < current.value {
			if current.left == nil {
				current.left = node
				node.parent = current
				break
			} else {
				current = current.left
			}
		} else {
			if current.right == nil {
				current.right = node
				node.parent = current
				break
			} else {
				current = current.right
			}
		}
	}
}

func (tree *RBTree) rotateLeft(node *Node) {
	y := node.right
	node.right = y.left
	if y.left != nil {
		y.left.parent = node
	}
	y.parent = node.parent
	if node.parent == nil {
		tree.root = y
	} else if node == node.parent.left {
		node.parent.left = y
	} else {
		node.parent.right = y
	}
	y.left = node
	node.parent = y
}

func (tree *RBTree) rotateRight(node *Node) {
	y := node.left
	node.left = y.right
	if y.right != nil {
		y.right.parent = node
	}
	y.parent = node.parent
	if node.parent == nil {
		tree.root = y
	} else if node == node.parent.right {
		node.parent.right = y
	} else {
		node.parent.left = y
	}
	y.right = node
	node.parent = y
}

func (tree *RBTree) search(node *Node, value int) bool {
	if node == nil {
		return false
	}
	// without recursion
	for node != nil {
		if value < node.value {
			node = node.left
		} else if value > node.value {
			node = node.right
		} else {
			return true
		}
	}
	return false
}

func (tree *RBTree) Search(value int) bool {
	return tree.search(tree.root, value)
}

func (tree *RBTree) InOrderTraversal(node *Node) {
	if node != nil {
		tree.InOrderTraversal(node.left)
		fmt.Println(node.value)
		tree.InOrderTraversal(node.right)
	}
}

func (tree *RBTree) PrintInOrder() {
	tree.InOrderTraversal(tree.root)
}

func main_test1() {
	tree := NewTree()
	arr := []int{7, 6, 5, 4, 3, 2, 1}
	for _, v := range arr {
		tree.Insert(v)
	}
	tree.PrintInOrder()
}
