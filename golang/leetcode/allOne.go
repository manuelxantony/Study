package main

// type Node struct {
// 	Value int
// 	Next *Node
// 	Prev *Node
// 	Keys map[string]int // ["leet", "hello"]
// }

// type List struct {
// 	head *Node
// 	tail *Node
// }

// type AllOne struct {
// 	allone map[string]*Node
// 	list *List
// 	valueNode map[int]*Node
// }

// func Constructor() AllOne {
//     return AllOne{
// 		allone: make(map[string]*Node),
// 	}
// }

// func (this *AllOne) Inc(key string)  {
// 	if this.list.head == nil || this.list.tail == nil {
// 		newNode := &Node{}
// 		this.list.head = newNode
// 		this.list.tail = newNode
// 	}

// 	if node, ok := this.allone[key]; ok {
// 		value := node.Value
// 		value += 1

// 		if node.Next.Next != nil{
// 			if node.Next.Value == value {
// 				// adding to keys
// 				node.Next.Keys[key] = 0
// 				this.allone[key] = node.Next
// 				delete(node.Keys, key)
// 			}else{
// 				keyMap := make(map[string]int)
// 				keyMap[key] = 0
// 				newNode := &Node{
// 					Value: value,
// 					Keys: keyMap,
// 				}
// 				next := node.Next
// 				node.Next = newNode
// 				newNode.Next = next
// 			}
// 		} else {
// 			keyMap := make(map[string]int)
// 			keyMap[key] = 0 // here value can be anything
// 			newNode := &Node{
// 				Value: value,
// 				Keys: keyMap,
// 			}
// 			node.Next = newNode
// 			this.allone[key] = node.Next
// 		}
// 	} else {
// 		// adding first time
// 		// checking node with value 1 is present
// 		if valueNode, okValue := this.valueNode[1]; okValue{
// 			valueNode.Keys[key] = 0

// 			return
// 		}

// 		keyMap := make(map[string]int)
// 		keyMap[key] = 0 // here value can be anything
// 		newNode := &Node{
// 			Value: 1,
// 			Keys: keyMap,
// 		}
// 		this.list.head.Next = newNode
// 		this.allone[key] = newNode
// 	}
// }

// func (this *AllOne) Dec(key string)  {

// }

// func (this *AllOne) GetMaxKey() string {

// }

// func (this *AllOne) GetMinKey() string {

// }
