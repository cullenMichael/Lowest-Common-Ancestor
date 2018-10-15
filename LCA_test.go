package main

import (
	"fmt"
	"strings"
	"testing"
)

//test to print out tree with 4 values
func TestPrintInsert(t *testing.T) {
	tree := &Tree{}
	addNode(tree, 10)
	addNode(tree, 1)
	addNode(tree, 15)
	addNode(tree, 25)
	str := DAGInsertprint(tree)
	if str != "10 1 15 25 " {
		t.Errorf("Wrong answer of Binary Tree, got:%s, wanted:10,1,15,25", str)
	} else {
		fmt.Printf("Correct! DAG Insert Works with 4 values!\n")
	}
}

//test to print tree with duplicate values
func TestPrintDoubleInsert(t *testing.T) {
	tree := &Tree{}
	addNode(tree, 10)
	addNode(tree, 1)
	addNode(tree, 10)
	addNode(tree, 25)
	addNode(tree, 25)
	addNode(tree, 15)
	str := DAGInsertprint(tree)
	if str != "10 1 25 15 " {
		t.Errorf("Wrong answer of Binary Tree, got:%s, wanted:10,1,25,15", str)
	} else {
		fmt.Printf("Correct! DAG Insert Works with Duplicate values!\n")
	}
}

//test to add edge 1 way between 2 nodes
func TestPrintAddEdge1way(t *testing.T) {
	tree := &Tree{}
	addNode(tree, 10)
	addNode(tree, 1)
	e := addEdge(tree, 10, 1)
	if e == nil {
		str := PrintEdges(tree)
		if str == "(10 in() out(1 ))    (1 in(10 ) out())    " {
			fmt.Printf("Correct! DAG Insert Edge Works with 1 pair + 1 edge!\n")
		} else {
			t.Errorf("Wrong answer of DAG insert Edge, got: %s, wanted: (10 in() out(1 ))    (1 in(10 ) out()) ", str)
		}
	} else {
		t.Errorf("Error Caused! %s", e.Error())
	}
}

//Tests 2 edges in the same direction
func TestDoubleEdge(t *testing.T) {
	tree := &Tree{}
	addNode(tree, 10)
	addNode(tree, 1)
	e := addEdge(tree, 10, 1)
	if e != nil {
		t.Errorf("Allowed Connection! Got: %s", e.Error())
	}
	q := addEdge(tree, 10, 1)
	if q == nil {
		fmt.Printf("Should Crash as Duplicate Edge! Got: %s\n", q.Error())
	} else {
		fmt.Printf("Correct! Detected Duplicate Edges! Got: %s\n", q.Error())
	}
}

//Test to add edge with a node not in the tree
func TestEdgeNoNode(t *testing.T) {
	tree := &Tree{}
	addNode(tree, 10)
	e := addEdge(tree, 10, 1)
	if e == nil {
		t.Errorf("Cannot add edge to Non Existant Node!")
	} else {
		fmt.Printf("Correct! Detected Node not in Tree! Got: %s\n", e.Error())
	}
}

//Test to add edge to itself
func TestEdgeToItself(t *testing.T) {
	tree := &Tree{}
	addNode(tree, 10)
	e := addEdge(tree, 10, 10)
	if e == nil {
		t.Errorf("Cannot add edge to itself!")
	} else {
		fmt.Printf("Correct! Detected Edge to iteslf! Got: %s\n", e.Error())
	}
}

//test to createa cycle between 2 nodes
func TestCycle2Node(t *testing.T) {
	tree := &Tree{}
	addNode(tree, 10)
	addNode(tree, 1)
	e := addEdge(tree, 10, 1)
	if e == nil {
		str := PrintEdges(tree)
		if str != "(10 in() out(1 ))    (1 in(10 ) out())    " {
			t.Errorf("Wrong answer of DAG insert Edge, got: %s, wanted: (10 in() out(1 ))    (1 in(10 ) out())  ", str)
		}
		q := addEdge(tree, 1, 10)
		if q == nil {
			t.Errorf("Insert Should Create Cycle!")
		} else {
			fmt.Printf("Correct! Detected Single Cycle! Got: %s\n", q.Error())
		}
	}
}

// test if detects cycle of 1 node
func TestCycle1Node(t *testing.T) {
	tree := &Tree{}
	addNode(tree, 10)
	e := addEdge(tree, 10, 10)
	if e == nil {
		t.Errorf("Cannot add edge to itself!")
	} else {
		fmt.Printf("Correct! Detected Cycle to Itself! Got: %s\n", e.Error())
	}
}

//Tests to see detection of a triangle cycle
func TestCycleTriangle(t *testing.T) {
	tree := &Tree{}
	addNode(tree, 10)
	addNode(tree, 5)
	addNode(tree, 1)
	e := addEdge(tree, 10, 5)
	if e != nil {
		t.Errorf("Allowed Connection! Got: %s", e.Error())
	}
	q := addEdge(tree, 5, 1)
	if q != nil {
		t.Errorf("Allowed Connection! Got: %s", q.Error())
	}
	r := addEdge(tree, 1, 10)
	if r == nil {
		t.Errorf("Wrong! Creates Cycle! Got: %s", r.Error())
	} else {
		fmt.Printf("Correct! Detected Cycle as Triangle! Got: %s\n", r.Error())
	}
}

//Tests Graph
func TestGraph(t *testing.T) {
	tree := &Tree{}
	addNode(tree, 0)
	addNode(tree, 1)
	addNode(tree, 2)
	addNode(tree, 3)
	addNode(tree, 4)
	addNode(tree, 5)
	addEdge(tree, 1, 0)
	addEdge(tree, 1, 2)
	addEdge(tree, 1, 4)
	addEdge(tree, 2, 3)
	addEdge(tree, 4, 5)
	addEdge(tree, 5, 3)
	addEdge(tree, 2, 0)
	var str = ""
	var str1 = ""
	gra := CreateGraph(tree)
	for _, d := range gra {
		str += strings.Replace(fmt.Sprint(d), " ", str1, -1)
	}
	if str == "[12][][1][25][1][4]" {
		fmt.Printf("Correct! Graph Created! Got: %s\n", str)
	} else {
		t.Errorf("Wrong! Graph should be outputted!")
	}
}

//Tests Routing of 2 nodes to every other node
func TestRoutes(t *testing.T) {
	tree := &Tree{}
	addNode(tree, 0)
	addNode(tree, 1)
	addNode(tree, 2)
	addNode(tree, 3)
	addNode(tree, 4)
	addNode(tree, 5)
	addEdge(tree, 1, 0)
	addEdge(tree, 1, 2)
	addEdge(tree, 1, 4)
	addEdge(tree, 2, 3)
	addEdge(tree, 4, 5)
	addEdge(tree, 5, 3)
	addEdge(tree, 2, 0)

	find(tree, 0, 3)
	rt1 := returnRoute1()
	rt2 := returnRoute2()
	s1 := arrayToString(rt1, ",")
	s2 := arrayToString(rt2, ",")

	if s1 != "0,1,1,9223372036854775807,9223372036854775807,9223372036854775807" {
		t.Errorf("Wrong! Incorrect route1! Got: %s", s1)
	}
	if s2 != "9223372036854775807,2,1,0,2,1" {
		t.Errorf("Wrong! Incorrect route2! Got: %s", s2)
	} else {
		fmt.Printf("Correct! Route1 & Route2 are right!\n")
	}
}

//Tests random route nodes
func TestRoutesRandom(t *testing.T) {
	tree := &Tree{}
	addNode(tree, 10) //0
	addNode(tree, 21) //1
	addNode(tree, 42) //2
	addNode(tree, 13) //3
	addNode(tree, 64) //4
	addNode(tree, 5)  //5
	addEdge(tree, 21, 10)
	addEdge(tree, 21, 42)
	addEdge(tree, 21, 64)
	addEdge(tree, 42, 13)
	addEdge(tree, 64, 5)
	addEdge(tree, 5, 13)
	addEdge(tree, 42, 10)
	find(tree, 10, 13)
	rt1 := returnRoute1()
	rt2 := returnRoute2()
	s1 := arrayToString(rt1, ",")
	s2 := arrayToString(rt2, ",")
	if s1 != "0,1,1,9223372036854775807,9223372036854775807,9223372036854775807" {
		t.Errorf("Wrong! Incorrect route1! Got: %s", s1)
	}
	if s2 != "9223372036854775807,2,1,0,2,1" {
		t.Errorf("Wrong! Incorrect route2! Got: %s", s2)
	}
	fmt.Printf("Correct! Route1 & Route2 are right for Random Nodes!\n")
}

//Tests LCA
func TestLCA(t *testing.T) {
	tree := &Tree{}
	addNode(tree, 0)
	addNode(tree, 1)
	addNode(tree, 2)
	addNode(tree, 3)
	addNode(tree, 4)
	addNode(tree, 5)
	addEdge(tree, 1, 0)
	addEdge(tree, 1, 2)
	addEdge(tree, 1, 4)
	addEdge(tree, 2, 3)
	addEdge(tree, 4, 5)
	addEdge(tree, 5, 3)
	addEdge(tree, 2, 0)
	lca, e := find(tree, 0, 3)
	if e == nil {
		if lca == 2 {
			fmt.Printf("Correct! LCA got 2!\n")
		} else {
			t.Errorf("Wrong! Incorrect Ancestor! Got: %d", lca)
		}
	} else {
		t.Errorf("Wrong! Error Occoured! Got: %s", e.Error())
	}
}

//Tests LCA Root as Ancestor
func TestLCARootAncestor(t *testing.T) {
	tree := &Tree{}
	addNode(tree, 0)
	addNode(tree, 1)
	addEdge(tree, 0, 1)
	lca, e := find(tree, 0, 1)
	if e == nil {
		if lca == 0 {
			fmt.Printf("Correct! LCA got 0!\n")
		} else {
			t.Errorf("Wrong! Incorrect Ancestor! Got: %d", lca)
		}
	} else {
		t.Errorf("Wrong! Error Occoured! Got: %s", e.Error())
	}
}

//Tests LCA With node not in tree
func TestLCANotInTree(t *testing.T) {
	tree := &Tree{}
	addNode(tree, 0)
	addNode(tree, 1)
	addNode(tree, 2)
	addNode(tree, 3)
	addNode(tree, 4)
	addNode(tree, 5)
	addEdge(tree, 1, 0)
	addEdge(tree, 1, 2)
	addEdge(tree, 1, 4)
	addEdge(tree, 2, 3)
	addEdge(tree, 4, 5)
	addEdge(tree, 5, 3)
	addEdge(tree, 2, 0)
	lca, e := find(tree, 0, 9)
	if e == nil {
		t.Errorf("Wrong! Node 9 not in Tree! Got: %d", lca)
	} else {
		fmt.Printf("Correct! Detected find route with node not in Tree! Got: %s\n", e.Error())
	}
}

//Tests LCA Withno edges
func TestLCANoEdges(t *testing.T) {
	tree := &Tree{}
	addNode(tree, 0)
	addNode(tree, 1)
	addNode(tree, 2)
	addNode(tree, 3)
	addNode(tree, 4)
	addNode(tree, 5)
	lca, e := find(tree, 0, 5)
	if e == nil {
		t.Errorf("Wrong! No Edges in Tree! Got: %d", lca)
	} else {
		fmt.Printf("Correct! Detected find route with no Edges! Got: %s\n", e.Error())
	}
}

//Tests LCA With node not in tree
func TestLCAFindIteslf(t *testing.T) {
	tree := &Tree{}
	addNode(tree, 0)
	addNode(tree, 1)
	addNode(tree, 2)
	addNode(tree, 3)
	addNode(tree, 4)
	addNode(tree, 5)
	addEdge(tree, 1, 0)
	addEdge(tree, 1, 2)
	addEdge(tree, 1, 4)
	addEdge(tree, 2, 3)
	addEdge(tree, 4, 5)
	addEdge(tree, 5, 3)
	addEdge(tree, 2, 0)
	lca, e := find(tree, 1, 1)
	if e == nil {
		t.Errorf("Wrong! Can't call find on the same Node! Got: %d", lca)
	} else {
		fmt.Printf("Correct! Detected call find on the same Node! Got: %s\n", e.Error())
	}
}
