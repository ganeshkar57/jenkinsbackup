// compile

// Copyright 2021 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

func fn() func(interface{}) {
	return func(o interface{}) {
		switch v := o.(type) {
		case *int:
			*v = 1
		}
	}
}

func main() {
	fn()
}
