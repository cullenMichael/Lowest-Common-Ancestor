package main

import "testing"

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
    str := print( tree.root)

    if str != "10 5 -50 -75 -10 80 60 30 15 55 75 85 " {
       t.Errorf("Wrong answer of %s", str)
    }
}