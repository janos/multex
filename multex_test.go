// Copyright (c) 2019, Janoš Guljaš <janos@resenje.org>
// All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package multex_test

import (
	"sync"
	"testing"
	"time"

	"resenje.org/multex"
)

func TestSingleKeyLocking(t *testing.T) {
	mu := multex.New[string]()

	delay := 100 * time.Millisecond

	start := time.Now()
	mu.Lock("key")
	unlocked := make(chan struct{})
	go func() {
		time.Sleep(delay)
		mu.Unlock("key")
		close(unlocked)
	}()

	<-unlocked
	if got := time.Since(start); got < delay {
		t.Errorf("unlocked in %s, before expected delay %s", got, delay)
	}
}

func TestMultipleKeysLocking(t *testing.T) {
	mu := multex.New[int]()

	mu.Lock(1)
	mu.Lock(2)

	unlocked1 := make(chan struct{})
	go func() {
		mu.Lock(1)
		defer mu.Unlock(1)

		close(unlocked1)
	}()

	unlocked2 := make(chan struct{})
	go func() {
		mu.Lock(2)
		defer mu.Unlock(2)

		close(unlocked2)
	}()

	mu.Unlock(1)
	<-unlocked1

	select {
	case <-unlocked2:
		t.Error("key 2 unlocked")
	default:
	}

	mu.Unlock(2)
	<-unlocked2
}

func BenchmarkMultex(b *testing.B) {
	mu := multex.New[string]()
	var r string
	for i := 0; i < b.N; i++ {
		mu.Lock("")
		r = ""
		mu.Unlock("")
	}
	_ = r
}

func BenchmarkMutex(b *testing.B) {
	var mu sync.Mutex
	var r string
	for i := 0; i < b.N; i++ {
		mu.Lock()
		r = ""
		mu.Unlock()
	}
	_ = r
}
