package internal

import "fmt"

type AVLTree struct {
	key    string
	value  string
	left   *AVLTree
	right  *AVLTree
	height int16
}

func Insert(node *AVLTree, key string, value string) *AVLTree {
	if node == nil {
		return createNode(key, value)
	}

	if key < node.key {
		node.left = Insert(node.left, key, value)
	} else {
		node.right = Insert(node.right, key, value)
	}

	node.height = updateHeight(node)
	bf := getHeight(node.left) - getHeight(node.right)

	if bf < -1 && key > node.right.key {
		return rotateLeft(node)
	}

	if bf > 1 && key < node.left.key {
		return rotateRight(node)
	}

	if bf < -1 && key < node.right.key {
		node.right = rotateRight(node.right)
		return rotateLeft(node)
	}

	if bf > 1 && key > node.left.key {
		node.left = rotateLeft(node.right)
		return rotateRight(node)
	}

	return node
}

func Get(node *AVLTree, key string) string {
	if node == nil {
		return "(nil)"
	}

	if node.key == key {
		return node.value
	}

	if key < node.key {
		return Get(node.left, key)
	} else {
		return Get(node.right, key)
	}
}

func createNode(key string, value string) *AVLTree {
	return &AVLTree{
		key:    key,
		value:  value,
		left:   nil,
		right:  nil,
		height: 1,
	}
}

func getHeight(node *AVLTree) int16 {
	if node == nil {
		return 0
	} else {
		return node.height
	}
}

func updateHeight(node *AVLTree) int16 {
	var (
		heightLeftSubtree  int16
		heightRightSubtree int16
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

func rotateLeft(node *AVLTree) *AVLTree {
	root := node.right
	rootLeftSubtree := root.left
	root.left = node
	node.right = rootLeftSubtree

	node.height = 1 + max(getHeight(node.left), getHeight(node.right))
	root.height = 1 + max(getHeight(root.left), getHeight(root.right))

	return root
}

func rotateRight(node *AVLTree) *AVLTree {
	root := node.left
	rootRightSubtree := root.right
	root.right = node
	node.left = rootRightSubtree

	node.height = 1 + max(getHeight(node.left), getHeight(node.right))
	root.height = 1 + max(getHeight(root.left), getHeight(root.right))

	return root
}

func Visualize(node *AVLTree) {
	if node == nil {
		return
	}
	Visualize(node.left)
	fmt.Println("Node", node)
	fmt.Println("Left", node.left)
	fmt.Println("Right", node.right)
	Visualize(node.right)
}
