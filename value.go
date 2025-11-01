// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gatomic

import (
	"sync/atomic"
)

// A Value provides an atomic load and store of a consistently typed value.
// The zero value for a Value returns nil from [Value.Load].
// Once [Value.Store] has been called, a Value must not be copied.
//
// A Value must not be copied after first use.
type Value[T any] atomic.Value

// Load returns the value set by the most recent Store.
// It returns zero value if there has been no call to Store for this Value.
func (v *Value[T]) Load() (val T) {
	val, _ = (*atomic.Value)(v).Load().(T)
	return
}

// Store sets the value of the [Value] v to val.
// All calls to Store for a given Value must use values of the same concrete type.
// Store of an inconsistent type panics, as does Store(nil).
func (v *Value[T]) Store(val T) {
	(*atomic.Value)(v).Store(val)
}

// Swap stores new into Value and returns the previous value. It returns zero value if
// the Value is empty.
//
// All calls to Swap for a given Value must use values of the same concrete
// type. Swap of an inconsistent type panics, as does Swap(nil).
func (v *Value[T]) Swap(new T) (old T) {
	old, _ = (*atomic.Value)(v).Swap(new).(T)
	return
}

// CompareAndSwap executes the compare-and-swap operation for the [Value].
//
// All calls to CompareAndSwap for a given Value must use values of the same
// concrete type. CompareAndSwap of an inconsistent type panics, as does
// CompareAndSwap(old, nil).
func (v *Value[T]) CompareAndSwap(old, new T) (swapped bool) {
	return (*atomic.Value)(v).CompareAndSwap(old, new)
}
