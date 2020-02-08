# Multiple mutual exclusion lock

[![GoDoc](https://godoc.org/resenje.org/multex?status.svg)](https://godoc.org/resenje.org/multex)
[![Go](https://github.com/janos/multex/workflows/Go/badge.svg)](https://github.com/janos/multex/actions?query=workflow%3AGo)

Package mutex provides multiple mutual exclusion lock. The name is
constructed by combining a common name for mutual exclusion locks, Mutex, and
word multiple, which is a property of this specific implementation. Multex
locking and unlocking for a single key is a few times slower then locking
with sync.Mutex, but provides the ability to lock the same block of code with
specific keys, allowing concurrent execution of the same code only for
different keys.

Performance comparison of Multex with a single key and sync.Mutex can be done
by running benchmarks in this package.

```
BenchmarkMultex-8       17200414                66.0 ns/op             0 B/op          0 allocs/op
BenchmarkMutex-8        95875299                11.8 ns/op             0 B/op          0 allocs/op
```

## Installation

Run `go get resenje.org/multex` from command line.
