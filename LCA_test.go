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
		t.Errorf("Wrong answer of Binary Tree", str)
	} else {
		fmt.Printf("Binary Tree Works Correctly!\n")
	}
}

func TestArraylist(t *testing.T) {
	s := find(1, 2)
	if s[0].value != 10 {
		t.Errorf("Wrong answer of Arraylist!")
	} else {
		fmt.Printf("Arraylist Works Correctly!\n")
	}
}
