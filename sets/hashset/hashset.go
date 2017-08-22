// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package hashset implements a set backed by a hash table.
//
// Structure is not thread safe.
//
// References: http://en.wikipedia.org/wiki/Set_%28abstract_data_type%29
package hashset

import (
	"sync"

	"github.com/emirpasic/gods/sets"
	hashsetnts "github.com/emirpasic/gods/sets/hashset"
)

func assertSetImplementation() {
	var _ sets.Set = (*Set)(nil)
}

// Set holds elements in go's native map
type Set struct {
	nts *hashsetnts.Set
	mu  sync.RWMutex
}

var itemExists = struct{}{}

// New instantiates a new empty set
func New() *Set {
	return &Set{nts: hashsetnts.New()}
}

// Add adds the items (one or more) to the set.
func (set *Set) Add(items ...interface{}) {
	set.mu.Lock()
	set.nts.Add(items...)
	set.mu.Unlock()
}

// Remove removes the items (one or more) from the set.
func (set *Set) Remove(items ...interface{}) {
	set.mu.Lock()
	set.nts.Remove(items...)
	set.mu.Unlock()
}

// Contains check if items (one or more) are present in the set.
// All items have to be present in the set for the method to return true.
// Returns true if no arguments are passed at all, i.e. set is always superset of empty set.
func (set *Set) Contains(items ...interface{}) bool {
	set.mu.RLock()
	defer set.mu.RUnlock()
	return set.nts.Contains(items...)
}

// Empty returns true if set does not contain any elements.
func (set *Set) Empty() bool {
	set.mu.RLock()
	defer set.mu.RUnlock()
	return set.nts.Empty()
}

// Size returns number of elements within the set.
func (set *Set) Size() int {
	set.mu.RLock()
	defer set.mu.RUnlock()
	return set.nts.Size()
}

// Clear clears all values in the set.
func (set *Set) Clear() {
	set.mu.Lock()
	set.nts.Clear()
	set.mu.Unlock()
}

// Values returns all items in the set.
func (set *Set) Values() []interface{} {
	set.mu.RLock()
	defer set.mu.RUnlock()
	return set.nts.Values()
}

// String returns a string representation of container
func (set *Set) String() string {
	set.mu.RLock()
	defer set.mu.RUnlock()
	return set.nts.String()
}
