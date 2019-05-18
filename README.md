# Multiple mutual exclusion lock

[![GoDoc](https://godoc.org/resenje.org/multex?status.svg)](https://godoc.org/resenje.org/multex)
[![Build Status](https://travis-ci.org/janos/multex.svg?branch=master)](https://travis-ci.org/janos/multex)

Package mutex provides multiple mutual exclusion lock. The name is
constructed by combining a common name for mutual exclusion locks, Mutex, and
word multiple, which is a property of this specific implementation. Multex
locking and unlocking for a single key is a few times slower then locking
with sync.Mutex, but provides the ability to lock the same block of code with
specific keys, allowing concurrent execution of the same code only for
different keys.

Performance comaprison of Multex with a single key and sync.Mutex can be done
by running benchmarks in this package.

```
BenchmarkMultex-8        20000000           69.6 ns/op         0 B/op          0 allocs/op
BenchmarkMutex-8        100000000           15.8 ns/op         0 B/op          0 allocs/op
```

## Installation

Run `go get resenje.org/multex` from command line.
