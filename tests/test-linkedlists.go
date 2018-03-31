package tests

import (
	"fmt"
	"linkedlist/linkedlist"
)

// utility for testing linked list
func assertEq(a, b []int) bool {

	if a == nil && b == nil {
		return true;
	}

	if a == nil || b == nil {
		return false;
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestLinkedList() {
	list := linkedlist.LinkedList{}

	// test appending
	for i := 1; i <= 5; i++  {
		list.Append(i)
	}
	if assertEq(list.ToSlice(), []int {1, 2, 3, 4, 5}) {
		fmt.Println("Appended items successful")
	} else {
		fmt.Println("Appending items failed")
		fmt.Printf("Expected %v to equal %v\n", list.ToSlice(), []int {1, 2, 3, 4, 5})
	}

	// test deletion
	// remove second item
	var secondNode  = list.Head.Next
	list.Remove(secondNode)
	if assertEq(list.ToSlice(), []int {1, 3, 4, 5}) {
		fmt.Println("Deleting second item successful")
	} else {
		fmt.Println("Deleting second item failed")
		fmt.Printf("Expected %v to equal %v\n", list.ToSlice(), []int {1, 3, 4, 5})
	}

	// remove first item
	list.Remove(list.Head)
	if assertEq(list.ToSlice(), []int {3, 4, 5}) {
		fmt.Println("Deleting first item successful")
	} else {
		fmt.Println("Deleting first item failed")
		fmt.Printf("Expected %v to equal %v\n", list.ToSlice(), []int {3, 4, 5})
	}

	// test find
	itemToFind := 4
	if list.Find(itemToFind) != nil {
		fmt.Printf("Finding item %d successful\n", itemToFind)
	} else {
		fmt.Printf("Finding item %d failed\n", itemToFind)
	}

	itemToFind = 1
	if list.Find(itemToFind) == nil {
		fmt.Println("Finding inexistent item successful")
	} else {
		fmt.Println("Finding inexistent item failed")
	}
}
