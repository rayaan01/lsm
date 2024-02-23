package main

import "memtable/internal"

func main() {
	root := internal.Insert(nil, 10)
	root = internal.Insert(root, 5)
	root = internal.Insert(root, 8)
	internal.Visualize(root)
}
