// Copyright (c) 2019, Janoš Guljaš <janos@resenje.org>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package multex provides multiple mutual exclusion lock. The name is
// constructed by combining a common name for mutual exclusion locks, Mutex, and
// word multiple, which is a property of this specific implementation. Multex
// locking and unlocking for a single key is a few times slower then locking
// with sync.Mutex, but provides the ability to lock the same block of code with
// specific keys, allowing concurrent execution of the same code only for
// different keys.
package multex

import (
	"sync"
)

// Multex is a mutual exclusion lock with support for multiple keys.
type Multex[K comparable] struct {
	c *sync.Cond
	s map[K]struct{}
}

// New constructs a new Multex instance.
func New[K comparable]() (m *Multex[K]) {
	return &Multex[K]{
		c: sync.NewCond(new(sync.Mutex)),
		s: make(map[K]struct{}),
	}
}

// Lock a specific key in Multex. This method is blocking until Unlock is called
// with the same key.
func (m *Multex[K]) Lock(key K) {
	m.c.L.Lock()

	for _, ok := m.s[key]; ok; _, ok = m.s[key] {
		m.c.Wait()
	}
	m.s[key] = struct{}{}

	m.c.L.Unlock()
}

// Unlock a specific key in Multex.
func (m *Multex[K]) Unlock(key K) {
	m.c.L.Lock()

	delete(m.s, key)
	m.c.Broadcast()

	m.c.L.Unlock()
}
