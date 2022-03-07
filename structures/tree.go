package main

import "fmt"

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Value int
}

var DefaultValue int = -1024

func InsertNodeToTree(tree *TreeNode, node *TreeNode) {
	if tree == nil {
		return
	}
	if tree.Value == DefaultValue {
		tree.Value = node.Value
		return
	}
	if node.Value > tree.Value {
		if tree.Right == nil {
			tree.Right = &TreeNode{Value: DefaultValue}
		}
		InsertNodeToTree(tree.Right, node)
	}
	if node.Value < tree.Value {
		if tree.Left == nil {
			tree.Left = &TreeNode{Value: DefaultValue}
		}
		InsertNodeToTree(tree.Left, node)
	}
}

func InitTree(values ...int) (root *TreeNode) {
	rootNode := TreeNode{Value: DefaultValue, Right: nil, Left: nil}
	for _, value := range values {
		node := TreeNode{Value: value}
		InsertNodeToTree(&rootNode, &node)
	}
	return &rootNode
}

func InOrderTraversal(tree *TreeNode) {
	if tree != nil {
		InOrderTraversal(tree.Left)
		fmt.Println(tree.Value)
		InOrderTraversal(tree.Right)
	}
}

func main() {
	tree := InitTree(1, 2, 3, 4, 5, 0, -1, -2)
	InOrderTraversal(tree)

}
