// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashbidimap implements a bidirectional map backed by two hashmaps.
//
// A bidirectional map, or hash bag, is an associative data structure in which the (key,value) pairs form a one-to-one correspondence.
// Thus the binary relation is functional in each direction: value can also act as a key to key.
// A pair (a,b) thus provides a unique coupling between 'a' and 'b' so that 'b' can be found when 'a' is used as a key and 'a' can be found when 'b' is used as a key.
//
// Elements are unordered in the map.
//
// Structure is not thread safe.
//
// Reference: https://en.wikipedia.org/wiki/Bidirectional_map
package hashbidimap

import (
	"sync"

	"github.com/emirpasic/gods/maps"
	hashbidimapnts "github.com/emirpasic/gods/maps/hashbidimap"
)

func assertMapImplementation() {
	var _ maps.BidiMap = (*Map)(nil)
}

// Map holds the elements in two hashmaps.
type Map struct {
	nts *hashbidimapnts.Map
	mu  sync.RWMutex
}

// New instantiates a bidirectional map.
func New() *Map {
	return &Map{nts: hashbidimapnts.New()}
}

// Put inserts element into the map.
func (m *Map) Put(key interface{}, value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.nts.Put(key, value)
}

// Get searches the element in the map by key and returns its value or nil if key is not found in map.
// Second return parameter is true if key was found, otherwise false.
func (m *Map) Get(key interface{}) (value interface{}, found bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.nts.Get(key)
}

// GetKey searches the element in the map by value and returns its key or nil if value is not found in map.
// Second return parameter is true if value was found, otherwise false.
func (m *Map) GetKey(value interface{}) (key interface{}, found bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.nts.GetKey(value)
}

// Remove removes the element from the map by key.
func (m *Map) Remove(key interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.nts.Remove(key)
}

func (m *Map) RemoveValue(value interface{}) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if key, found := m.nts.GetKey(value); found {
		m.nts.Remove(key)
	}
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
	defer m.mu.Unlock()
	m.nts.Clear()
}

// String returns a string representation of container
func (m *Map) String() string {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.nts.String()
}
