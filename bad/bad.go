// package main

import "fmt"

// 節・葉に持たせる機能のインターフェース
type Item interface {
	Eq(Item) bool
	Less(Item) bool
}

// インターフェースの実装
type Int int

func (n Int) Eq(m Item) bool {
	return n == m.(Int)
}
func (n Int) Less(m Item) bool {
	return n < m.(Int)
}

// 節・葉の構造
type Node struct {
	item        Item
	left, right *Node
}

// 節・葉の作成
func newNode(x Item) *Node {
	p := new(Node)
	p.item = x
	return p
}

// 二分木
type Tree struct {
	root *Node
}

// 二分木の生成
func newTree() *Tree {
	return new(Tree)
}

// データの探索
func (t *Tree) searchTree(x Item) bool {
	return searchNode(t.root, x)
}

// 木への挿入
func (t *Tree) insertTree(x Item) {
	t.root = insertNode(t.root, x)
}

// 節・葉の挿入
func insertNode(node *Node, x Item) *Node {
	switch {
	case node == nil:
		return newNode(x)
	case x.Eq(node.item):
		return node
	case x.Less(node.item):
		node.left = insertNode(node.left, x)
	default:
		node.right = insertNode(node.right, x)
	}
	return node
}

// 節・葉の探索
func searchNode(node *Node, x Item) bool {
	for node != nil {
		switch {
		case x.Eq(node.item):
			return true
		case x.Less(node.item):
			node = node.left
		default:
			node = node.right
		}
	}
	return false
}

// 木のダンプ
func (t *Tree) printTree() {
	t.foreachTree(func(x Item) { fmt.Print(x, " ") })
	fmt.Println("")
}

// 木からの削除
func (t *Tree) deleteTree(x Item) {
	t.root = deleteNode(t.root, x)
}

// 節・葉の削除
func deleteNode(node *Node, x Item) *Node {
	if node != nil {
		if x.Eq(node.item) {
			if node.left == nil {
				return node.right
			} else if node.right == nil {
				return node.left
			} else {
				node.item = searchMin(node.right)
				node.right = deleteMin(node.right)
			}
		} else if x.Less(node.item) {
			node.left = deleteNode(node.left, x)
		} else {
			node.right = deleteNode(node.right, x)
		}
	}
	return node
}

// 最小値を求める
func searchMin(node *Node) Item {
	if node.left == nil {
		return node.item
	}
	return searchMin(node.left)
}

// 最小値を削除する
func deleteMin(node *Node) *Node {
	if node.left == nil {
		return node.right
	}
	node.left = deleteMin(node.left)
	return node
}

// 巡回
func foreachNode(f func(Item), node *Node) {
	if node != nil {
		foreachNode(f, node.left)
		f(node.item)
		foreachNode(f, node.right)
	}
}
func (t *Tree) foreachTree(f func(Item)) {
	foreachNode(f, t.root)
}

func main() {
	a := newTree()
	for _, v := range []int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0} {
		a.insertTree(Int(v))
	}
	a.printTree()
	for i := 0; i < 10; i++ {
		fmt.Println(a.searchTree(Int(i)))
	}
	for i := 0; i < 10; i++ {
		a.deleteTree(Int(i))
		a.printTree()
	}
}
