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
	out   []*Channel
	in    []*Channel
	value int
}
type Channel struct {
	node *Node
}

//Creates a tree structure of nodes
type Tree struct {
	nodes []*Node
}

func addNode(t *Tree, val int) error {

	for _, i := range t.nodes {

		if i.value == val {
			return errors.New("Node Already Exists")
		}
	}
	t.nodes = append(t.nodes, &Node{value: val, in: nil, out: nil})
	return nil
}

//Prints DAG Insert tree
func DAGInsertprint(t *Tree) string {
	var str = ""
	for _, i := range t.nodes {
		str = str + strconv.Itoa(i.value) + " "
	}
	return str
}

func find(t *Tree, v1 int, v2 int) (int, error) {
	route1 = nil
	route2 = nil
	return 0, nil
}

//returns the common ancestor of the two nodes
func ancestor(root *Node, v1 int, v2 int) (int, error) {

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

	return false
}

func (t *Tree) insert(val int) *Tree {
	// if t.root == nil {
	// 	t.root = &Node{value: val, left: nil, right: nil}
	// } else {
	// 	t.root.insert(val)
	// }
	return t
}

func (n *Node) insert(val int) {

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
	// if node == nil {
	// 	return " "
	// }
	// str = str + strconv.Itoa(node.value) + " "
	// print(node.left)
	// print(node.right)
	return " "
}

func main() {
}
