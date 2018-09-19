package main

import (
    "fmt" 
    "strconv" 
)
  
  var str = ""
type Node struct { 
    left  *Node
    right *Node
    value  int
}
  
type Tree struct {
    root *Node
}
  
func (t *Tree) insert(val int) *Tree {
    if t.root == nil {
        t.root = &Node{value: val, left: nil, right: nil}
    } else {
        t.root.insert(val)
    }
    return t
}
  
func (n *Node) insert(val int) {
    if n == nil {
        return
    } else if val <= n.value {
        if n.left == nil {
            n.left = &Node{value: val, left: nil, right: nil}
        }else {
            n.left.insert(val)
        }
    }else {
        if n.right == nil {
            n.right = &Node{value: val, left: nil, right: nil}
        }else {
            n.right.insert(val)
        }
    }   
}
//prints content of tree
func print( node *Node) string{
    if node == nil {
        return " "
    }
    str =  str + strconv.Itoa(node.value) + " "
    print(node.left)
    print(node.right)
    return str
}


  
func main() {
    tree := &Tree{}
    tree.insert(10).
        insert(5).
        insert(-50).
        insert(-75).
        insert(80).
        insert(60).
        insert(30).
        insert(55).
        insert(85).
        insert(15).
        insert(75).
        insert(-10)
    fmt.Printf(print(tree.root))
}