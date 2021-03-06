package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var str = "" // global variable for print
var route1 []int
var route2 []int

// a node consists of its value and its children (Left & Right)
type Node struct {
	left  *Node
	right *Node
	value int
}

//Creates a tree structure of nodes
type Tree struct {
	root *Node
}

func find(t *Tree, v1 int, v2 int) (int, error) {
	route1 = nil
	route2 = nil
	return ancestor(t.root, v1, v2)
}

//returns the common ancestor of the two nodes
func ancestor(root *Node, v1 int, v2 int) (int, error) {

	if !findPath(root, v1, true) || !findPath(root, v2, false) { // checks if route exists
		return -1, errors.New("No Path Exists!")
	}

	if (len(route1) == 1) || (len(route2) == 1) {
		if route1[0] == route2[0] {
			return route1[0], errors.New("EEE") //if arrays length is 1 return first position
		} else {
			return -1, errors.New("No Paths Exist!")
		}

	} else if len(route1) > len(route2) {
		for i, _ := range route2 {
			if route1[i] != route2[i] {
				return route1[i-1], nil
			}
		}
	} else {
		for i, _ := range route1 {
			if route1[i] != route2[i] {
				return route1[i-1], nil
			}
		}
	}
	return -1, errors.New("No Paths Exist!")
}

// returns route1 array
func returnRoute1() []int {
	return route1
}

// returns route2 array
func returnRoute2() []int {
	return route2
}

// convers an int[] to string for testing
func arrayToString(a []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(a), " ", delim, -1), "[]")
}

//finds the track from the route and populates an array,
//returns true if a path exists, else false
func findPath(root *Node, n int, path bool) bool {

	if root == nil {
		return false
	}

	if path {
		route1 = append(route1, root.value)
		if root.value == n {
			return true
		}
		if root.left != nil && findPath(root.left, n, path) {
			return true
		}
		if root.right != nil && findPath(root.right, n, path) {
			return true
		}
		route1 = route1[:len(route1)-1]
	} else {
		route2 = append(route2, root.value)
		if root.value == n {
			return true
		}
		if root.left != nil && findPath(root.left, n, path) {
			return true
		}
		if root.right != nil && findPath(root.right, n, path) {
			return true
		}
		route2 = route2[:len(route2)-1]
	}
	return false
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
		} else {
			n.left.insert(val)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value: val, left: nil, right: nil}
		} else {
			n.right.insert(val)
		}
	}
}

func printTree(node *Node) string {
	str = ""
	if node == nil {
		return ""
	}
	return print(node)
}

//prints content of tree
func print(node *Node) string {
	if node == nil {
		return " "
	}
	str = str + strconv.Itoa(node.value) + " "
	print(node.left)
	print(node.right)
	return str
}

func main() {
}
