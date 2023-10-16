package main

import (
	"fmt"
)

type Node struct{
	Left *Node
	Value int
	Right *Node
}

type BinaryTree struct {
	Root *Node
}

func (bt *BinaryTree) insert(val int) *BinaryTree{
	if bt.Root != nil{
		bt.Root.insert(val)
	} else {
		rootNode := &Node{Value: val}
		bt.Root = rootNode
	}
	return bt
}


func (n *Node) insert(val int){
	if n == nil{
		return
	}
	if val < n.Value{
		if n.Left == nil{
			newNode := &Node{Value: val}
			n.Left = newNode
		} else {
			n.Left.insert(val)
		}
	} else {
		if n.Right == nil {
			newNode := &Node{Value: val}
			n.Right = newNode
		} else {
			n.Right.insert(val)
		}
	}
}


func (bt *BinaryTree) levelOrder() []int{
	if bt.Root == nil {
		return []int{}
	}

	out := []int{}
	queue := []Node{}

	queue = append(queue, *bt.Root)

	for len(queue) != 0{
		out = append(out, queue[0].Value) 
		if queue[0].Left != nil{
			queue = append(queue, *queue[0].Left)
		}
		if queue[0].Right != nil {
			queue = append(queue, *queue[0].Right)
		}
		queue = queue[1:] // dequeuing
	}

	return out
}

func (bt *BinaryTree) preOrder(node *Node) {
	// D L R
	if node == nil {
		return
	}

	fmt.Print (node.Value, " ")	
	
	// the following will not work
	// node = node.Left
	// tree.preOrder(node)
	// node = node.Right
	// tree.preOrder(node)

	// the following will work 
	// n := node.Left
	// tree.preOrder(n)
	// n = node.Right
	// tree.preOrder(n)

	// because 
	// in the first case we are using node itself, 
	// changing the same node to diffrent node in refernce  line 83 and 85

	bt.preOrder(node.Left)
	bt.preOrder(node.Right)
	
}

func (bt *BinaryTree) inOrder(node *Node){
	if node == nil {
		return
	}
	bt.inOrder(node.Left)
	fmt.Print(node.Value, " ")
	bt.inOrder(node.Right)
}

func (bt *BinaryTree) postOrder(node *Node){
	if node == nil {
		return
	}
	bt.postOrder(node.Left)
	bt.postOrder(node.Right)
	fmt.Print(node.Value, " ")
}

func (n *Node) delete(node *Node, value int) *Node {
	if node == nil {
		return nil
	} else if value < node.Value { // lesser
		node.Left = n.delete(node.Left, value)
	} else if value > node.Value{ // greater
		node.Right = n.delete(node.Right, value)
	} else{ // node found
		if node.Left == nil && node.Right == nil {
			node = nil
		} else if node.Left == nil { // case 2, 1 child
			node = node.Right
		} else if node.Right == nil{
			node = node.Left
		} else{ // case  3, 2 children
			minNode := n.FindMin(node.Right)
			node.Value = minNode.Value
			node.Right = n.delete(node.Right, minNode.Value)
		}
	}
	return node
}


func (n *Node) FindMin(node *Node) *Node{
	if node == nil {
		return nil
	}
	for node.Left != nil{
		node = node.Left
	}
	return node
}

func main() {
	tree := &BinaryTree{}
	tree.insert(20).insert(9).insert(49).insert(5).insert(12).insert(15).insert(23).insert(56).insert(53).insert(54)
	// fmt.Println("Level-Order")
	// fmt.Println(tree.levelOrder())	
	// fmt.Println("Pre-Order")
	// tree.preOrder(tree.Root)
	// fmt.Println("\nIn-Order")
	// tree.inOrder(tree.Root)
	// fmt.Println("\nPost-Order")
	// tree.postOrder(tree.Root)
	tree.inOrder(tree.Root)
	fmt.Println()
	tree.Root.delete(tree.Root, 12)
	tree.inOrder(tree.Root)
	fmt.Println()


}