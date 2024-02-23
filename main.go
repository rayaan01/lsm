package main

import "fmt"

type BinaryTree struct {
	value  int
	left   *BinaryTree
	right  *BinaryTree
	height int
}

func visualize(node *BinaryTree) {
	if node == nil {
		return
	}
	visualize(node.left)
	fmt.Println("Node", node)
	fmt.Println("Left", node.left)
	fmt.Println("Right", node.right)
	visualize(node.right)
}

func createNode(value int) *BinaryTree {
	return &BinaryTree{
		value:  value,
		left:   nil,
		right:  nil,
		height: 1,
	}
}

func getHeight(node *BinaryTree) int {
	if node == nil {
		return 0
	} else {
		return node.height
	}
}

func updateHeight(node *BinaryTree) int {
	var (
		heightLeftSubtree  int
		heightRightSubtree int
	)

	if node.left == nil {
		heightLeftSubtree = 0
	} else {
		heightLeftSubtree = node.left.height
	}

	if node.right == nil {
		heightRightSubtree = 0
	} else {
		heightRightSubtree = node.right.height
	}

	return 1 + max(heightLeftSubtree, heightRightSubtree)
}

func getBalanceFactor(node *BinaryTree) int {
	return getHeight(node.left) - getHeight(node.right)
}

func insert(node *BinaryTree, value int) *BinaryTree {
	if node == nil {
		return createNode(value)
	}
	if value < node.value {
		node.left = insert(node.left, value)
	} else {
		node.right = insert(node.right, value)
	}

	node.height = updateHeight(node)
	bf := getBalanceFactor(node)

	if bf < -1 && value > node.right.value {
		return rotateLeft(node)
	}

	if bf > 1 && value < node.left.value {
		return rotateRight(node)
	}

	if bf < -1 && value < node.right.value {
		node.right = rotateRight(node.right)
		return rotateLeft(node)
	}

	if bf > 1 && value > node.left.value {
		node.left = rotateLeft(node.right)
		return rotateRight(node)
	}

	return node
}

func rotateLeft(node *BinaryTree) *BinaryTree {
	root := node.right
	rootLeftSubtree := root.left
	root.left = node
	node.right = rootLeftSubtree

	node.height = 1 + max(getHeight(node.left), getHeight(node.right))
	root.height = 1 + max(getHeight(root.left), getHeight(root.right))
	return root
}

func rotateRight(node *BinaryTree) *BinaryTree {
	root := node.left
	rootRightSubtree := root.right
	root.right = node
	node.left = rootRightSubtree

	node.height = 1 + max(getHeight(node.left), getHeight(node.right))
	root.height = 1 + max(getHeight(root.left), getHeight(root.right))
	return root
}

func main() {
	root := insert(nil, 10)
	root = insert(root, 5)
	root = insert(root, 8)
	visualize(root)
}
