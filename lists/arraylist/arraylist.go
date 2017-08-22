package arraylist

import (
	"sync"

	"github.com/emirpasic/gods/lists"
	arraylistnts "github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/utils"
)

func assertListImplementation() {
	var _ lists.List = (*List)(nil)
}

// List holds the elements in a slice
type List struct {
	nts *arraylistnts.List
	mu  sync.RWMutex
}

// New instantiates a new empty list
func New() *List {
	return &List{nts: arraylistnts.New()}
}

// Add appends a value at the end of the list
func (list *List) Add(values ...interface{}) {
	list.mu.Lock()
	list.nts.Add(values...)
	list.mu.Unlock()
}

// Get returns the element at index.
// Second return parameter is true if index is within bounds of the array and array is not empty, otherwise false.
func (list *List) Get(index int) (interface{}, bool) {
	list.mu.RLock()
	defer list.mu.RUnlock()
	return list.nts.Get(index)
}

// Remove removes one or more elements from the list with the supplied indices.
func (list *List) Remove(index int) {
	list.mu.Lock()
	list.nts.Remove(index)
	list.mu.Unlock()
}

// Contains checks if elements (one or more) are present in the set.
// All elements have to be present in the set for the method to return true.
// Performance time complexity of n^2.
// Returns true if no arguments are passed at all, i.e. set is always super-set of empty set.
func (list *List) Contains(values ...interface{}) bool {
	list.mu.RLock()
	defer list.mu.RUnlock()
	return list.nts.Contains(values...)
}

// Values returns all elements in the list.
func (list *List) Values() []interface{} {
	list.mu.RLock()
	defer list.mu.RUnlock()
	return list.nts.Values()
}

// Empty returns true if list does not contain any elements.
func (list *List) Empty() bool {
	list.mu.RLock()
	defer list.mu.RUnlock()
	return list.nts.Empty()
}

// Size returns number of elements within the list.
func (list *List) Size() int {
	list.mu.RLock()
	defer list.mu.RUnlock()
	return list.nts.Size()
}

// Clear removes all elements from the list.
func (list *List) Clear() {
	list.mu.Lock()
	list.nts.Clear()
	list.mu.Unlock()
}

// Sort sorts values (in-place) using.
func (list *List) Sort(comparator utils.Comparator) {
	list.mu.Lock()
	list.nts.Sort(comparator)
	list.mu.Unlock()
}

// Swap swaps the two values at the specified positions.
func (list *List) Swap(i, j int) {
	list.mu.Lock()
	list.nts.Swap(i, j)
	list.mu.Unlock()
}

// Insert inserts values at specified index position shifting the value at that position (if any) and any subsequent elements to the right.
// Does not do anything if position is negative or bigger than list's size
// Note: position equal to list's size is valid, i.e. append.
func (list *List) Insert(index int, values ...interface{}) {
	list.mu.Lock()
	list.nts.Insert(index, values...)
	list.mu.Unlock()
}

// String returns a string representation of container
func (list *List) String() string {
	list.mu.RLock()
	defer list.mu.RUnlock()
	return list.nts.String()
}
