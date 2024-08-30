package hash_tables

import (
	"errors"
	"fmt"
)

// This implementation assumes key value pairs to be of type: int && string
// It uses hashing by division
// It uses chaining with linked lists

var (
	defaultTableSize    = 10
	loadFactorThreshold = 0.75 // Resize table when the load factor exceeds 75%
)

type Entry struct {
	key   int
	value string
	next  *Entry
}

type Hash struct {
	entries []*Entry
}

type HashTable struct {
	table      *Hash
	numEntries *int
}

// HashFunction Calculate index for new entry
func hashFunction(key int, size int) int {
	return key % size
}

// NewHashTable Create a new hash table
func NewHashTable() HashTable {
	num := 0
	hash := Hash{make([]*Entry, defaultTableSize)}
	return HashTable{table: &hash, numEntries: &num}
}

// Display Displays HashTable as human-readable
func (ht *HashTable) Display() {
	for i, entry := range ht.table.entries {
		fmt.Printf("%d :", i)
		for entry != nil {
			fmt.Printf("[%d, %s]->", entry.key, entry.value)
			entry = entry.next
		}
		fmt.Println("nil")
	}
	fmt.Printf("----- \n")
}

// Put add new element into HashTable
func (ht *HashTable) Put(key int, value string) {

	if float64(*ht.numEntries)/float64(len(ht.table.entries)) > loadFactorThreshold {
		ht.resize()
	}

	index := hashFunction(key, len(ht.table.entries))
	entry := ht.table.entries[index]
	newEntry := &Entry{key: key, value: value}

	// If there are no entries at this index, insert the new entry directly
	if entry == nil {
		ht.table.entries[index] = newEntry
	} else {
		// Traverse the linked list to the end
		for entry.next != nil {
			entry = entry.next
		}
		// Append the new entry to the end of the list
		entry.next = newEntry
	}

	// Increment the number of entries in the hash table
	*ht.numEntries++

}

func (ht *HashTable) Find(key int) string {
	index := hashFunction(key, len(ht.table.entries))
	entry := ht.table.entries[index]

	for entry != nil {
		if entry.key == key {
			return entry.value
		}
		entry = entry.next
	}

	return "Entry not found"
}

func (ht *HashTable) Delete(key int) error {
	index := hashFunction(key, len(ht.table.entries))
	entry := ht.table.entries[index]

	// If there's no entry at this index, return an error
	if entry == nil {
		return errors.New("no value found for given key")
	}

	// If there's only one entry and it matches the key
	if entry.key == key && entry.next == nil {
		ht.table.entries[index] = nil
		*ht.numEntries--
		return nil
	}

	// Traverse to find the last entry with the matching key
	var prev *Entry
	for entry != nil {
		if entry.key == key {
			if entry.next == nil {
				if prev != nil {
					// Remove the last entry by updating the previous entry's next pointer
					prev.next = nil
				} else {
					ht.table.entries[index] = nil
				}
				*ht.numEntries--
				return nil
			}
			prev = entry
		}
		entry = entry.next
	}

	return nil
}

// Resize increases the size of the hash table and rehashes all entries
func (ht *HashTable) resize() {
	fmt.Printf("resizing")
	oldEntries := ht.table.entries
	oldSize := len(oldEntries)
	newSize := oldSize * 2
	newEntries := make([]*Entry, newSize)

	ht.table.entries = newEntries

	for _, entry := range oldEntries {
		for entry != nil {
			next := entry.next

			index := hashFunction(entry.key, newSize)
			entry.next = newEntries[index]
			newEntries[index] = entry

			entry = next
		}
	}

	// Update the table size and reset numEntries
	*ht.numEntries = 0
}
