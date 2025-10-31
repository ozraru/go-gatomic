// Copyright 2022 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package gatomic_test

import (
	"fmt"

	"github.com/ozraru/go-gatomic"
)

type Fruit int32

const (
	Apple Fruit = iota
	Banana
	Cherry
)

var atomicFruit gatomic.Int32[Fruit]

func Example() {
	go func() {
		atomicFruit.Store(Apple)
	}()
	go func() {
		atomicFruit.Store(Banana)
	}()
	go func() {
		atomicFruit.Store(Cherry)
	}()
	switch atomicFruit.Load() {
	case Apple:
		fmt.Println("Apple")
	case Banana:
		fmt.Println("Banana")
	case Cherry:
		fmt.Println("Cherry")
	}
}
