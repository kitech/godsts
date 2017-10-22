package hashmap

import (
	"sync"

	"github.com/emirpasic/gods/maps"
	hashmapnts "github.com/emirpasic/gods/maps/hashmap"
)

func assertMapImplementation() {
	var _ maps.Map = (*Map)(nil)
}

// Map holds the elements in go's native map
type Map struct {
	nts *hashmapnts.Map
	mu  sync.RWMutex
}

// New instantiates a hash map.
func New() *Map {
	return &Map{nts: hashmapnts.New()}
}

// Put inserts element into the map.
func (m *Map) Put(key interface{}, value interface{}) {
	m.mu.Lock()
	m.nts.Put(key, value)
	m.mu.Unlock()
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map) Get(key interface{}) (value interface{}, found bool) {
	m.mu.RLock()
	value, found = m.nts.Get(key)
	m.mu.RUnlock()
	return
}

// Remove removes the element from the map by key.
func (m *Map) Remove(key interface{}) {
	m.mu.Lock()
	m.nts.Remove(key)
	m.mu.Unlock()
}

// Empty returns true if map does not contain any elements
func (m *Map) Empty() bool {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.nts.Empty()
}

// Size returns number of elements in the map.
func (m *Map) Size() int {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.nts.Size()
}

// Keys returns all keys (random order).
func (m *Map) Keys() []interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.nts.Keys()
}

// Values returns all values (random order).
func (m *Map) Values() []interface{} {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.nts.Values()
}

// Clear removes all elements from the map.
func (m *Map) Clear() {
	m.mu.Lock()
	m.nts.Clear()
	m.mu.Unlock()
}

// String returns a string representation of container
func (m *Map) String() string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.nts.String()
}

func (m *Map) Has(key interface{}) bool {
	_, has := m.Get(key)
	return has
}
