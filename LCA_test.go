package main

import (
	"fmt"
	"testing"
)

func TestPrint(t *testing.T) {

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
	str := print(tree.root)

	if str != "10 5 -50 -75 -10 80 60 30 15 55 75 85 " {
		t.Errorf("Wrong answer of Binary Tree, got: %s, wnated: 10 5 -50 -75 -10 80 60 30 15 55 75 85 ", str)
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

	if s1 != "10,5,-50,-75,-10,80,60,30" {
		t.Errorf("Wrong answer %s should be 10,5,-50,-75,-10,80,60,30", s1)
	} else {
		fmt.Printf("Arraylist Works Correctly for route1!\n")
	}

	if s2 != "10,5,-50,-75,-10,80,60,30,15,55" {
		t.Errorf("Wrong answer %s sould be 10,5,-50,-75,-10,80,60,30,15,55", s2)
	} else {
		fmt.Printf("Arraylist Works Correctly for route2!\n")
	}
}
