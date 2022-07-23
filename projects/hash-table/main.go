package main

import "fmt"

// hashTable structure
// - insert
// - search
// - delete

// bucket structure
// - insert
// - search
// - delete

const ARRAY_SIZE = 7

// HashTable will hold an array
type HashTable struct {
	array [ARRAY_SIZE]*bucket
}

// bucket structure
type bucket struct {
	// store address of bucketNode
	head *bucketNode
}

// bucketNode structure
type bucketNode struct {
	key  string
	next *bucketNode
}

func InitHashTable() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}

	return result
}

func main() {
	testHashTable := InitHashTable()
	fmt.Println(testHashTable)
}
