// This code reffer to https://qiita.com/oko1977/items/822c0b3168716ebfbf0c for learning purpose.

// 節・葉に持たせる機能のインターフェース(間接的;直接的なインターフェースの多様性・効率性を実現するための機能)
type Item interface {
	Eq(Item) bool
	Less(Item) bool
}

// ラッパーとしての新しい型インターフェースの実装(直接的)
type Int int

// Intに対し、nを値に持つとして関数を実装
func (n Int) Eq(m Item) bool {
	return n == m.(Int)
}
func (n Int) Less(m Item) bool {
	return n < m.(Int)
}

// Go言語の仕様、オーバーライドが存在しないことに注意→同dir.memo_btree.txt中l.23参照

// 節・葉の構造
type Node struct {
	item        Item
	left, right *Node // *：ポインタ
}

// 節・葉の作成
func newNode(x Item) *Node {
	p := new(Node)
	p.item = x
	return p
}

// 2分木
type Tree struct {
	root *Node
}

// 2分木の生成
func newtree() *Tree {
	return new(Tree)
}

// データの探索
func (t *Tree) searchTree(x Item) bool {
	return searchNode(t.root, x)
}

// 木への挿入
func (t *Tree) insertTree(x Item) bool {
	t.root = insertNode(t.root, x)
}

// 節・葉の挿入
func insertNode(node *Node, x Item) {
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
	t.root=deleteNode(t.root, x)
}
// 節・葉の削除
func deleteNode(node *Node,x Item) *Node {
	if node.left==nil
}
