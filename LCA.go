package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

var str = "" // global variable for print
var route1 []int
var route2 []int
var index = 0
var nCount1 = 0
var nCount2 = 0

// a node consists of its value and its children (Left & Right)
type Node struct {
	out   []*Node
	in    []*Node
	value int
}

//Creates a tree structure of nodes
type Tree struct {
	nodes     []*Node
	nodeCount int
}

//adds a node to the graph
func addNode(t *Tree, val int) error {

	for _, i := range t.nodes {

		if i.value == val {
			return errors.New("Node Already Exists")
		}
	}
	t.nodes = append(t.nodes, &Node{value: val, in: nil, out: nil})
	t.nodeCount++
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

// adds an edge between 2 nodes in the tree
func addEdge(t *Tree, ind1 int, ind2 int) error {

	var j2 = -1
	var j1 = -1
	if ind1 == ind2 {
		return errors.New("Cant add edge to itself!")
	}
	for i, in := range t.nodes {
		if ind1 == in.value {
			j1 = i
		}
		if ind2 == in.value {
			j2 = i
		}
	}

	if j1 == -1 {
		return errors.New("Node 1 does not Exist!")
	}
	if j2 == -1 {
		return errors.New("Node 2 does not Exist!")
	}

	for _, j := range t.nodes[j1].out {
		if j.value == ind2 {
			return errors.New("Cant have Duplicate Edges!")
		}

	}
	t.nodes[j1].out = append(t.nodes[j1].out, t.nodes[j2])
	t.nodes[j2].in = append(t.nodes[j2].in, t.nodes[j1])
	if Cycles(t) {
		t.nodes[j1].out = t.nodes[j1].out[:len(t.nodes[j1].out)-1]
		t.nodes[j2].in = t.nodes[j1].in[:len(t.nodes[j1].in)-1]
		return errors.New("Creates Cycle!")
	}
	return nil
}

// prints the edges in and out of a node
func PrintEdges(t *Tree) string {

	var str = ""
	for _, i := range t.nodes {
		str += "("
		str += strconv.Itoa(i.value) + " in("
		for _, j := range i.in {
			str += strconv.Itoa(j.value) + " "
		}
		str += ")" + " out("
		for _, j := range i.out {
			str += strconv.Itoa(j.value) + " "
		}
		str += "))    "
	}
	return str
}

// checks for cycles of a tree
func Cycles(t *Tree) bool {

	visited := make([]bool, t.nodeCount)
	recStack := make([]bool, t.nodeCount)
	var hit = false

	for i := 0; i < t.nodeCount-1; i++ {
		if isCyclicUtil(i, visited, recStack, t) {
			hit = true
			break
		}
	}
	return hit
}

//helper function of Cycles
func isCyclicUtil(i int, v []bool, r []bool, t *Tree) bool {

	if r[i] {
		return true
	}
	if v[i] {
		return false
	}
	v[i] = true
	r[i] = true
	out := t.nodes[i].out

	for _, j := range out {
		for q, z := range t.nodes {
			if z.value == j.value {
				index = q
				break
			}
		}
		if isCyclicUtil(index, v, r, t) {
			return true
		}
	}
	r[i] = false
	return false
}

func find(t *Tree, v1 int, v2 int) (int, error) {
	route1 = make([]int, t.nodeCount)
	route2 = make([]int, t.nodeCount)
	nCount1 = 0
	nCount2 = 0
	CreateGraph(t)
	lca, e := ancestor(t, v1, v2, CreateGraph(t))
	if e == nil {
		return lca, nil
	} else {
		return -1, e
	}
}

func CreateGraph(t *Tree) [][]int {

	graph := make([][]int, t.nodeCount)
	for i := 0; i < t.nodeCount; i++ {
		for _, v := range t.nodes[i].in {
			graph[i] = append(graph[i], v.value)
		}
	}
	return graph
}

//returns the common ancestor of the two nodes
func ancestor(t *Tree, v1i int, v2i int, g [][]int) (int, error) {

	var v1 = -1
	var v2 = -1
	for i, j := range t.nodes {
		route1[i] = math.MaxInt64
		route2[i] = math.MaxInt64
		if j.value == v1i {
			v1 = i
		}
		if j.value == v2i {
			v2 = i
		}
	}
	if (v1 == -1) || (v2 == -1) {
		return -1, errors.New("Node does not Exist!")
	}
	for i := 0; i < t.nodeCount; i++ {
		route1[i] = math.MaxInt64
		route2[i] = math.MaxInt64
	}
	route1[v1] = 0
	route2[v2] = 0
	route1 := getRoute(t, v1i, g, 0, route1)
	route2 := getRoute(t, v2i, g, 0, route2)
	var count = math.MaxInt64
	var c2 = 0
	for j, c1 := range route1 {
		c2 = route2[j]
		if (c2+c1 < count) && (c2 != 0) && (c1 != 0) && (c1 != math.MaxInt64) && (c2 != math.MaxInt64) {
			count = c2 + c1
		}
	}
	if count == math.MaxInt64 {
		return -1, errors.New("No Path Exists!")
	}
	return count, nil
}

//DFS of all the connected nodes
func getRoute(t *Tree, v3 int, g [][]int, count int, r []int) []int {
	var v1 = 0
	for j, i := range t.nodes {
		if i.value == v3 {
			v1 = j
		}
	}
	arr := g[v1]
	if r[v1] > count {
		r[v1] = count
	}
	for _, i := range arr {
		count++
		getRoute(t, i, g, count, r)
		count--
	}
	return r
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

func main() {
}
