package main

import (
	"fmt"
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

//
// //Tests print of an empty binary Tree
// func TestPrintEmptyTree(t *testing.T) {
//
// 	tree := &Tree{}
// 	str := printTree(tree.root)
// 	if str != "" {
// 		t.Errorf("Wrong answer of Binary Tree, got: %s, wanted:  ", str)
// 	} else {
// 		fmt.Printf("Correct! Binary Tree Works With 1 Element!\n")
// 	}
// }
//
// //Tests a binary tree with 1 element
// func TestPrint1Tree(t *testing.T) {
//
// 	tree1 := &Tree{}
// 	tree1.insert(2)
// 	str1 := printTree(tree1.root)
// 	if str1 != "2 " {
// 		t.Errorf("Wrong answer of Binary Tree, got: %s, wanted: 2 ", str1)
// 	} else {
// 		fmt.Printf("Correct! Binary Tree Works With No Elements!\n")
// 	}
//
// }
//
// //Tests that a full binary tree is in order
// func TestPrintFullTree(t *testing.T) {
//
// 	tree := &Tree{}
// 	tree.insert(10)
// 	str := printTree(tree.root)
// 	tree.insert(5).
// 		insert(-50).
// 		insert(-75).
// 		insert(80).
// 		insert(60).
// 		insert(30).
// 		insert(55).
// 		insert(85).
// 		insert(15).
// 		insert(75).
// 		insert(-10)
// 	str = printTree(tree.root)
//
// 	if str != "10 5 -50 -75 -10 80 60 30 15 55 75 85 " {
// 		t.Errorf("Wrong answer of Binary Tree, got: %s, wanted: 10 5 -50 -75 -10 80 60 30 15 55 75 85 ", str)
// 	} else {
// 		fmt.Printf("Correct! Binary Tree Works with a full Tree\n")
// 	}
// }
//
// //Tests that the route arraylist works
// func TestArraylist(t *testing.T) {
//
// 	tree := &Tree{}
// 	tree.insert(10).
// 		insert(5).
// 		insert(-50).
// 		insert(-75).
// 		insert(80).
// 		insert(60).
// 		insert(30).
// 		insert(55).
// 		insert(85).
// 		insert(15).
// 		insert(75).
// 		insert(-10)
// 	find(tree, -10, 55)
// 	r1 := returnRoute1()
// 	r2 := returnRoute2()
// 	s1 := arrayToString(r1, ",")
// 	s2 := arrayToString(r2, ",")
//
// 	if s1 != "10,5,-50,-10" {
// 		t.Errorf("Wrong answer %s should be 10,5,-50,-10", s1)
// 	} else {
// 		fmt.Printf("Correct! Arraylist Works for route1!\n")
// 	}
// 	if s2 != "10,80,60,30,55" {
// 		t.Errorf("Wrong answer %s should be 10,80,60,30,55", s2)
// 	} else {
// 		fmt.Printf("Correct! Arraylist Works for route2!\n")
// 	}
// }
//
// //finds route beween to nodes of an empty tree
// func TestFindEmpty(t *testing.T) {
// 	tree := &Tree{}
// 	f, e := find(tree, 5, 10)
// 	if e.Error() == "" {
// 		t.Errorf("Wrong answer %d should be No Paths Exist!", f)
// 	} else {
// 		fmt.Printf("Correct! Search in empty tree returns: %s\n", e.Error())
// 	}
//
// }
//
// //finds ancestor of 2 nodes in the tree
// func TestFind2Elements(t *testing.T) {
// 	tree := &Tree{}
// 	tree.insert(10).
// 		insert(5).
// 		insert(-50).
// 		insert(-75).
// 		insert(80).
// 		insert(60).
// 		insert(30).
// 		insert(55).
// 		insert(85).
// 		insert(15).
// 		insert(75).
// 		insert(-10)
// 	f, e := find(tree, -10, 55)
// 	if e != nil {
// 		if f != 10 {
// 			t.Errorf("Wrong answer %d should be 10", f)
// 		} else {
// 			fmt.Printf("Correct! Finding ancestor of -10 & 55 is 10!\n")
// 		}
// 	}
// }
//
// //Tests one node in the tree and one not in the tree: returns Error
// func TestFind1Elements(t *testing.T) {
// 	tree := &Tree{}
// 	tree.insert(10)
// 	f, e := find(tree, 10, 55)
// 	if e.Error() == "" {
// 		t.Errorf("Wrong answer %d should be Error %s", f, e.Error())
// 	} else {
// 		fmt.Printf("Correct! Error of one node in tree returns: %s!\n", e.Error())
// 	}
// }
//
// //Tests an element in the tree and one that is not. returns error
// func TestFind1ElementInTree(t *testing.T) {
//
// 	tree := &Tree{}
// 	tree.insert(10).
// 		insert(5).
// 		insert(-50).
// 		insert(-75).
// 		insert(80).
// 		insert(60).
// 		insert(30).
// 		insert(55).
// 		insert(85).
// 		insert(15).
// 		insert(75).
// 		insert(-10)
// 	f, e := find(tree, -10, 99)
// 	if e.Error() == "" {
// 		t.Errorf("Wrong answer %d should be No Paths Exist!", f)
// 	} else {
// 		fmt.Printf("Correct! One nodes not in the tree returns: %s\n", e.Error())
// 	}
// }
//
// //Searches for no 2 nodes in the tree: returns error
// func TestNoElementInTree(t *testing.T) {
//
// 	tree := &Tree{}
// 	tree.insert(10).
// 		insert(5).
// 		insert(-50).
// 		insert(-75).
// 		insert(80).
// 		insert(60).
// 		insert(30).
// 		insert(55).
// 		insert(85).
// 		insert(15).
// 		insert(75).
// 		insert(-10)
// 	f, e := find(tree, 98, 99)
// 	if e.Error() == "" {
// 		t.Errorf("Wrong answer %d should be No Paths Exist!", f)
// 	} else {
// 		fmt.Printf("Correct! Two nodes not in the tree returns: %s\n", e.Error())
// 	}
// }
//
// //Tests an ancestor with itself
// func TestFindItself(t *testing.T) {
//
// 	tree := &Tree{}
// 	tree.insert(10)
// 	f, e := find(tree, 10, 10)
// 	if e != nil {
// 		if f != 10 {
// 			t.Errorf("Wrong answer %d should be 10", f)
// 		} else {
// 			fmt.Printf("Correct! Finding ancestor of 10 & 10 is 10!\n")
// 		}
// 	}
// }
//
// //Tests root as the ancestor
// func TestRootAsParent(t *testing.T) {
//
// 	tree := &Tree{}
// 	tree.insert(10).insert(5).insert(20)
// 	f, e := find(tree, 10, 5)
// 	if e != nil {
// 		if f != 10 {
// 			t.Errorf("Wrong answer %d should be 10", f)
// 		} else {
// 			fmt.Printf("Correct! Finding ancestor of 10 & 5 is 10!\n")
// 		}
// 	}
// }
//
// func TestChildAsparent(t *testing.T) {
// 	tree := &Tree{}
// 	tree.insert(10).
// 		insert(10).
// 		insert(10)
// 	f, e := find(tree, 10, 10)
// 	if e != nil {
// 		if f != 10 {
// 			t.Errorf("Wrong answer %d should be 10", f)
// 		} else {
// 			fmt.Printf("Correct! Finding ancestor of 10 & 10 is 10!\n")
// 		}
// 	}
// }
