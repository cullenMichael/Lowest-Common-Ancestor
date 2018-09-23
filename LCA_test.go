package main

import (
	"fmt"
	"testing"
)

func TestPrintEmptyTree(t *testing.T) {

	tree := &Tree{}
	str := printTree(tree.root)
	if str != "" {
		t.Errorf("Wrong answer of Binary Tree, got: %s, wanted:  ", str)
	} else {
		fmt.Printf("Binary Tree Works Correctly With 1 Element!\n")
	}
}

func TestPrint1Tree(t *testing.T) {

	tree1 := &Tree{}
	tree1.insert(2)
	str1 := printTree(tree1.root)
	if str1 != "2 " {
		t.Errorf("Wrong answer of Binary Tree, got: %s, wanted: 2 ", str1)
	} else {
		fmt.Printf("Binary Tree Works Correctly With No Elements!\n")
	}

}

func TestPrintFullTree(t *testing.T) {

	tree := &Tree{}
	tree.insert(10)
	str := printTree(tree.root)
	tree.insert(5).
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
	str = printTree(tree.root)

	if str != "10 5 -50 -75 -10 80 60 30 15 55 75 85 " {
		t.Errorf("Wrong answer of Binary Tree, got: %s, wanted: 10 5 -50 -75 -10 80 60 30 15 55 75 85 ", str)
	} else {
		fmt.Printf("Binary Tree Works Correctly!\n")
	}
}

func TestArraylist(t *testing.T) {

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
	find(tree, -10, 55)
	r1 := returnRoute1()
	r2 := returnRoute2()
	s1 := arrayToString(r1, ",")
	s2 := arrayToString(r2, ",")

	if s1 != "10,5,-50,-10" {
		t.Errorf("Wrong answer %s should be 10,5,-50,-10", s1)
	} else {
		fmt.Printf("Arraylist Works Correctly for route1!\n")
	}
	if s2 != "10,80,60,30,55" {
		t.Errorf("Wrong answer %s should be 10,80,60,30,55", s2)
	} else {
		fmt.Printf("Arraylist Works Correctly for route2!\n")
	}
}

func TestFind2Elements(t *testing.T) {

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
	f := find(tree, -10, 55)
	if f != 10 {
		t.Errorf("Wrong answer %d should be 10", f)
	} else {
		fmt.Printf("Finding ancestor of -10 & 55 is 10!\n")
	}
}

func TestFind1ElementInTree(t *testing.T) {

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
	f := find(tree, -10, 99)
	if f != -1 {
		t.Errorf("Wrong answer %d should be -1", f)
	} else {
		fmt.Printf("Finding ancestor of -10 & 99 is -1!\n")
	}
}
