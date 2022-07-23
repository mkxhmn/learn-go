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

const ARRAY_SIZE = 2

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
func (h *HashTable) Search(key string) bool {
	index := hash(key)

	return h.array[index].search(key)
}

// insert will take in a key and delete from the hash table
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// bucket structure
type bucket struct {
	// store address of bucketNode
	head *bucketNode
}

func (b *bucket) insert(k string) {
	if !b.search(k) {
		newNode := &bucketNode{key: k}

		// become the current head
		newNode.next = b.head

		// assign a new head
		b.head = newNode
	} else {
		fmt.Printf("%v already exist\n", k)
	}
}

// delete will take in a key and delete the node if exist
// linked list will be skipping the currentNode
// and make the currentNode point to the next node
func (b *bucket) delete(k string) {
	previousNode := b.head

	// handle matching key is the head
	if b.head.key == k {
		b.head = b.head.next

		return
	}

	for previousNode != nil {
		if previousNode.next.key == k {
			previousNode.next = previousNode.next.next
		}

		previousNode = previousNode.next
	}
}

// search will take in a key and return true if exist
func (b *bucket) search(k string) bool {
	currentNode := b.head

	// loop until the current node is empty
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}

		currentNode = currentNode.next

	}

	return false
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
	hashTable := InitHashTable()
	randy := "RANDY"

	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		randy,
		"BUTTERS",
		"TOKEN",
	}

	for _, v := range list {
		hashTable.Insert(v)
	}

  // expect true
	fmt.Println(hashTable.Search(randy))

	hashTable.Delete(randy)
  // expect false
	fmt.Println(hashTable.Search(randy))
}
