package tests

import (
	//"fmt"
	"linkedlist/linkedlist"
	//"container/list"
	"testing"
	"github.com/stretchr/testify/assert"
)

func get12345LinkedList() linkedlist.LinkedList {
	list := linkedlist.LinkedList{}

	for i := 1; i <= 5; i++  {
		list.Append(i)
	}

	return list
}

func TestAppendingItemsToList(t *testing.T) {
	list := linkedlist.LinkedList{}

	// test appending
	for i := 1; i <= 5; i++  {
		list.Append(i)
	}

	assert.Equal(t, []int {1, 2, 3, 4, 5}, list.ToSlice(), "List Appending failed.")
}

func TestDeletingFirstItem(t *testing.T) {
	list := get12345LinkedList()
	firstNode := list.Head

	list.Remove(firstNode)

	assert.Equal(t, []int {2, 3, 4, 5}, list.ToSlice())
}

func TestDeletingSecondItem(t *testing.T) {
	list := get12345LinkedList()
	secondNode := list.Head.Next

	list.Remove(secondNode)

	assert.Equal(t, []int {1, 3, 4, 5}, list.ToSlice())
}

func TestFindingExistingItem(t *testing.T) {
	list := get12345LinkedList()
	itemToFind := 4

	assert.NotNil(t, list.Find(itemToFind), "Failed to find existing item in list")
}

func TestGettingNonExistingItem(t *testing.T) {
	list := get12345LinkedList()
	itemToFind := 6

	assert.Nil(t, list.Find(itemToFind), "Found a non existing item in list")
}

//func TestLinkedList() {
//
//
//	// test find
//	itemToFind := 4
//	if list.Find(itemToFind) != nil {
//		fmt.Printf("Finding item %d successful\n", itemToFind)
//	} else {
//		fmt.Printf("Finding item %d failed\n", itemToFind)
//	}
//
//	itemToFind = 1
//	if list.Find(itemToFind) == nil {
//		fmt.Println("Finding inexistent item successful")
//	} else {
//		fmt.Println("Finding inexistent item failed")
//	}
//}
