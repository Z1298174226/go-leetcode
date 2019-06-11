package main

import (
	"container/list"
	"fmt"
)

type Tree struct {
	left  *Tree
	right *Tree
	val   int
}

func preReverse(root *Tree) {
	l := list.New()
	currentNode := root
	for {
		if currentNode == nil && l.Len() == 0 {
			break
		} else {
			for {
				if currentNode == nil {
					break
				} else {
					fmt.Println(currentNode.val)
					l.PushBack(currentNode)
					currentNode = currentNode.left
				}
			}
			if l.Len() > 0 {
				N := l.Back()
				Node, ok := N.Value.(*Tree)
				if ok {
					l.Remove(N)
					currentNode = Node.right
				} else {
					fmt.Errorf("类型不匹配")
				}
			}
		}
	}
}

func postReverse(root *Tree) {
	currentNode := root
	var lastNode *Tree
	l := list.New()
	for {
		if currentNode == nil {
			break
		} else {
			l.PushBack(currentNode)
			currentNode = currentNode.left
		}
	}
	for {
		if l.Len() == 0 {
			break
		} else {
			Node := l.Back()
			currentNode = Node.Value.(*Tree)
			if currentNode.right != nil && currentNode.right != lastNode {
				currentNode = currentNode.right
				for {
					if currentNode == nil {
						break
					} else {
						l.PushBack(currentNode)
						currentNode = currentNode.left
					}

				}
			} else {
				l.Remove(Node)
				fmt.Println(currentNode.val)
				lastNode = currentNode
			}
		}
	}
}

func main() {
	root := Tree{left: &Tree{left: nil, right: &Tree{left: nil, right: nil, val: 3}, val: 2}, right: nil, val: 1}
	postReverse(&root)
}
