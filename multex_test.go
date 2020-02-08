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
	mu := multex.New()

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
	mu := multex.New()

	mu.Lock("key1")
	mu.Lock("key2")

	unlocked1 := make(chan struct{})
	go func() {
		mu.Lock("key1")
		defer mu.Unlock("key1")

		close(unlocked1)
	}()

	unlocked2 := make(chan struct{})
	go func() {
		mu.Lock("key2")
		defer mu.Unlock("key2")

		close(unlocked2)
	}()

	mu.Unlock("key1")
	<-unlocked1

	select {
	case <-unlocked2:
		t.Error("key2 unlocked")
	default:
	}

	mu.Unlock("key2")
	<-unlocked2
}

func BenchmarkMultex(b *testing.B) {
	mu := multex.New()
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
