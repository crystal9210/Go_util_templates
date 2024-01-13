package main

import (
	"fmt"
	"math"
	"strings"
)

type Item interface {
	Eq(Item) bool
	Less(Item) bool
}

type Int int

func (n Int) Eq(m Item) bool {
	return m != nil && n == m.(Int)
}

func (n Int) Less(m Item) bool {
	return m != nil && n < m.(Int)
}

type Node struct {
	item        Item
	left, right *Node
}

func newNode(x Item) *Node {
	return &Node{item: x}
}

type Tree struct {
	root *Node
}

func newTree() *Tree {
	return &Tree{}
}

func (t *Tree) insertTree(x Item) {
	t.root = insertNode(t.root, x)
}

func insertNode(node *Node, x Item) *Node {
	if node == nil {
		return newNode(x)
	}
	if x.Less(node.item) {
		node.left = insertNode(node.left, x)
	} else if x.Eq(node.item) {
		// Handle duplicate values as needed
	} else {
		node.right = insertNode(node.right, x)
	}
	return node
}

func (t *Tree) searchTree(x Item) bool {
	return searchNode(t.root, x)
}

func searchNode(node *Node, x Item) bool {
	if node == nil {
		return false
	}
	if x.Eq(node.item) {
		return true
	}
	if x.Less(node.item) {
		return searchNode(node.left, x)
	}
	return searchNode(node.right, x)
}

func (t *Tree) deleteTree(x Item) {
	t.root = deleteNode(t.root, x)
}

func deleteNode(node *Node, x Item) *Node {
	if node == nil {
		return nil
	}
	if x.Eq(node.item) {
		// Node with only one child or no child
		if node.left == nil {
			return node.right
		} else if node.right == nil {
			return node.left
		}
		// Node with two children: Get the inorder successor (smallest in the right subtree)
		node.item = minValue(node.right)
		// Delete the inorder successor
		node.right = deleteNode(node.right, node.item)
	} else if x.Less(node.item) {
		node.left = deleteNode(node.left, x)
	} else {
		node.right = deleteNode(node.right, x)
	}
	return node
}

func minValue(node *Node) Item {
	current := node
	for current.left != nil {
		current = current.left
	}
	return current.item
}

// ツリーの深さを計算する関数
func treeDepth(node *Node) int {
	if node == nil {
		return 0
	}
	leftDepth := treeDepth(node.left)
	rightDepth := treeDepth(node.right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

// 特定の深さのノードを表示する関数
func printLevel(node *Node, level int, indentSize int) string {
	if node == nil {
		return strings.Repeat(" ", indentSize) // 空のスペースで埋める
	}
	if level == 1 {
		return fmt.Sprintf("%s%d%s", strings.Repeat(" ", indentSize/2), node.item, strings.Repeat(" ", indentSize/2))
	} else if level > 1 {
		left := printLevel(node.left, level-1, indentSize/2)
		right := printLevel(node.right, level-1, indentSize/2)
		return left + right
	}
	return ""
}

// ツリー全体を表示する関数
func (t *Tree) printTree() {
	depth := treeDepth(t.root)
	indentSize := int(math.Pow(2, float64(depth))) // インデントのサイズ
	for i := 1; i <= depth; i++ {
		fmt.Println(printLevel(t.root, i, indentSize))
	}
}

// // ツリーを視覚的に表示する関数
// func printNode(node *Node, prefix string, isLeft bool) {
// 	if node != nil {
// 			fmt.Println(prefix + getDirection(isLeft) + fmt.Sprint(node.item))
// 			newPrefix := prefix + getBranch(isLeft)
// 			printNode(node.left, newPrefix, true)
// 			printNode(node.right, newPrefix, false)
// 	}
// }

// func getDirection(isLeft bool) string {
// 	if isLeft {
// 			return "├── "
// 	}
// 	return "└── "
// }

// func getBranch(isLeft bool) string {
// 	if isLeft {
// 			return "│   "
// 	}
// 	return "    "
// }

// // ツリー全体を表示する関数
//
//	func (t *Tree) printTree() {
//		printNode(t.root, "", false)
//	}
func main() {
	a := newTree()
	for _, v := range []int{5, 6, 4, 7, 3, 8, 2, 9, 1} {
		a.insertTree(Int(v))
	}
	a.printTree()

	fmt.Println("Searching for 6:", a.searchTree(Int(6)))
	fmt.Println("Searching for 10:", a.searchTree(Int(10)))

	a.deleteTree(Int(5))
	a.printTree()
}
