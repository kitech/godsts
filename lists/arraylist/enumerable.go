// Copyright (c) 2015, Emir Pasic. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package arraylist

import "github.com/emirpasic/gods/containers"

func assertEnumerableImplementation() {
	var _ containers.EnumerableWithIndex = (*List)(nil)
}

// Each calls the given function once for each element, passing that element's index and value.
func (list *List) Each(f func(index int, value interface{})) {
	list.mu.Lock()
	list.nts.Each(f)
	list.mu.Unlock()
}

// Map invokes the given function once for each element and returns a
// container containing the values returned by the given function.
func (list *List) Map(f func(index int, value interface{}) interface{}) *List {
	list.mu.Lock()
	defer list.mu.Unlock()
	return &List{nts: list.nts.Map(f)}
}

// Select returns a new container containing all elements for which the given function returns a true value.
func (list *List) Select(f func(index int, value interface{}) bool) *List {
	list.mu.RLock()
	defer list.mu.RUnlock()
	return &List{nts: list.nts.Select(f)}
}

// Any passes each element of the collection to the given function and
// returns true if the function ever returns true for any element.
func (list *List) Any(f func(index int, value interface{}) bool) bool {
	list.mu.RLock()
	defer list.mu.RUnlock()
	return list.nts.Any(f)
}

// All passes each element of the collection to the given function and
// returns true if the function returns true for all elements.
func (list *List) All(f func(index int, value interface{}) bool) bool {
	list.mu.RLock()
	defer list.mu.RUnlock()
	return list.nts.All(f)
}

// Find passes each element of the container to the given function and returns
// the first (index,value) for which the function is true or -1,nil otherwise
// if no element matches the criteria.
func (list *List) Find(f func(index int, value interface{}) bool) (int, interface{}) {
	list.mu.RLock()
	defer list.mu.RUnlock()
	return list.nts.Find(f)
}
