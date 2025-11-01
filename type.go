// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gatomic

import (
	"sync/atomic"
)

// A Bool is an atomic boolean value.
// The zero value is false.
//
// Bool must not be copied after first use.
type Bool[T ~bool] atomic.Bool

// Load atomically loads and returns the value stored in x.
func (x *Bool[T]) Load() T { return T((*atomic.Bool)(x).Load()) }

// Store atomically stores val into x.
func (x *Bool[T]) Store(val T) { (*atomic.Bool)(x).Store(bool(val)) }

// Swap atomically stores new into x and returns the previous value.
func (x *Bool[T]) Swap(new T) (old T) { return T((*atomic.Bool)(x).Swap(bool(new))) }

// CompareAndSwap executes the compare-and-swap operation for the boolean value x.
func (x *Bool[T]) CompareAndSwap(old, new T) (swapped bool) {
	return (*atomic.Bool)(x).CompareAndSwap(bool(old), bool(new))
}

// An Int32 is an atomic int32. The zero value is zero.
//
// Int32 must not be copied after first use.
type Int32[T ~int32] atomic.Int32

// Load atomically loads and returns the value stored in x.
func (x *Int32[T]) Load() T { return T((*atomic.Int32)(x).Load()) }

// Store atomically stores val into x.
func (x *Int32[T]) Store(val T) { (*atomic.Int32)(x).Store(int32(val)) }

// Swap atomically stores new into x and returns the previous value.
func (x *Int32[T]) Swap(new T) (old T) { return T((*atomic.Int32)(x).Swap(int32(new))) }

// CompareAndSwap executes the compare-and-swap operation for x.
func (x *Int32[T]) CompareAndSwap(old, new T) (swapped bool) {
	return (*atomic.Int32)(x).CompareAndSwap(int32(old), int32(new))
}

// Add atomically adds delta to x and returns the new value.
func (x *Int32[T]) Add(delta T) (new T) { return T((*atomic.Int32)(x).Add(int32(delta))) }

// And atomically performs a bitwise AND operation on x using the bitmask
// provided as mask and returns the old value.
func (x *Int32[T]) And(mask T) (old T) { return T((*atomic.Int32)(x).And(int32(mask))) }

// Or atomically performs a bitwise OR operation on x using the bitmask
// provided as mask and returns the old value.
func (x *Int32[T]) Or(mask T) (old T) { return T((*atomic.Int32)(x).Or(int32(mask))) }

// An Int64 is an atomic int64. The zero value is zero.
//
// Int64 must not be copied after first use.
type Int64[T ~int64] atomic.Int64

// Load atomically loads and returns the value stored in x.
func (x *Int64[T]) Load() T { return T((*atomic.Int64)(x).Load()) }

// Store atomically stores val into x.
func (x *Int64[T]) Store(val T) { (*atomic.Int64)(x).Store(int64(val)) }

// Swap atomically stores new into x and returns the previous value.
func (x *Int64[T]) Swap(new T) (old T) { return T((*atomic.Int64)(x).Swap(int64(new))) }

// CompareAndSwap executes the compare-and-swap operation for x.
func (x *Int64[T]) CompareAndSwap(old, new T) (swapped bool) {
	return (*atomic.Int64)(x).CompareAndSwap(int64(old), int64(new))
}

// Add atomically adds delta to x and returns the new value.
func (x *Int64[T]) Add(delta T) (new T) { return T((*atomic.Int64)(x).Add(int64(delta))) }

// And atomically performs a bitwise AND operation on x using the bitmask
// provided as mask and returns the old value.
func (x *Int64[T]) And(mask T) (old T) { return T((*atomic.Int64)(x).And(int64(mask))) }

// Or atomically performs a bitwise OR operation on x using the bitmask
// provided as mask and returns the old value.
func (x *Int64[T]) Or(mask T) (old T) { return T((*atomic.Int64)(x).Or(int64(mask))) }

// A Uint32 is an atomic uint32. The zero value is zero.
//
// Uint32 must not be copied after first use.
type Uint32[T ~uint32] atomic.Uint32

// Load atomically loads and returns the value stored in x.
func (x *Uint32[T]) Load() T { return T((*atomic.Uint32)(x).Load()) }

// Store atomically stores val into x.
func (x *Uint32[T]) Store(val T) { (*atomic.Uint32)(x).Store(uint32(val)) }

// Swap atomically stores new into x and returns the previous value.
func (x *Uint32[T]) Swap(new T) (old T) { return T((*atomic.Uint32)(x).Swap(uint32(new))) }

// CompareAndSwap executes the compare-and-swap operation for x.
func (x *Uint32[T]) CompareAndSwap(old, new T) (swapped bool) {
	return (*atomic.Uint32)(x).CompareAndSwap(uint32(old), uint32(new))
}

// Add atomically adds delta to x and returns the new value.
func (x *Uint32[T]) Add(delta T) (new T) { return T((*atomic.Uint32)(x).Add(uint32(delta))) }

// And atomically performs a bitwise AND operation on x using the bitmask
// provided as mask and returns the old value.
func (x *Uint32[T]) And(mask T) (old T) { return T((*atomic.Uint32)(x).And(uint32(mask))) }

// Or atomically performs a bitwise OR operation on x using the bitmask
// provided as mask and returns the old value.
func (x *Uint32[T]) Or(mask T) (old T) { return T((*atomic.Uint32)(x).Or(uint32(mask))) }

// A Uint64 is an atomic uint64. The zero value is zero.
//
// Uint64 must not be copied after first use.
type Uint64[T ~uint64] atomic.Uint64

// Load atomically loads and returns the value stored in x.
func (x *Uint64[T]) Load() T { return T((*atomic.Uint64)(x).Load()) }

// Store atomically stores val into x.
func (x *Uint64[T]) Store(val T) { (*atomic.Uint64)(x).Store(uint64(val)) }

// Swap atomically stores new into x and returns the previous value.
func (x *Uint64[T]) Swap(new T) (old T) { return T((*atomic.Uint64)(x).Swap(uint64(new))) }

// CompareAndSwap executes the compare-and-swap operation for x.
func (x *Uint64[T]) CompareAndSwap(old, new T) (swapped bool) {
	return (*atomic.Uint64)(x).CompareAndSwap(uint64(old), uint64(new))
}

// Add atomically adds delta to x and returns the new value.
func (x *Uint64[T]) Add(delta T) (new T) { return T((*atomic.Uint64)(x).Add(uint64(delta))) }

// And atomically performs a bitwise AND operation on x using the bitmask
// provided as mask and returns the old value.
func (x *Uint64[T]) And(mask T) (old T) { return T((*atomic.Uint64)(x).And(uint64(mask))) }

// Or atomically performs a bitwise OR operation on x using the bitmask
// provided as mask and returns the old value.
func (x *Uint64[T]) Or(mask T) (old T) { return T((*atomic.Uint64)(x).Or(uint64(mask))) }
