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

func hash(key string) int {
	sum := 0

	for _, v := range key {
		sum += int(v)
	}

	return sum % ARRAY_SIZE
}

// insert will take in a key and add it to the hash table array
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)

}

// search will take in a key and return true if exist
func (h *HashTable) Search(key string) {

}

// insert will take in a key and delete from the hash table
func (h *HashTable) Delete(key string) {

}

// bucket structure
type bucket struct {
	// store address of bucketNode
	head *bucketNode
}

func (b *bucket) insert(k string) {
	newNode := &bucketNode{key: k}

	// become the current head
	newNode.next = b.head

	// assign a new head
	b.head = newNode
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

  testBucket := &bucket{}
  testBucket.insert("RANDY")
}
