// Copyright ©2015 The gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package c128

// DotuUnitary is
//  for i, v := range x {
//  	sum += y[i] * v
//  }
//  return sum
func DotuUnitary(x, y []complex128) (sum complex128) {
	for i, v := range x {
		sum += y[i] * v
	}
	return sum
}

// DotuInc is
//  for i := 0; i < int(n); i++ {
//  	sum += y[iy] * x[ix]
//  	ix += incX
//  	iy += incY
//  }
//  return sum
func DotuInc(x, y []complex128, n, incX, incY, ix, iy uintptr) (sum complex128) {
	for i := 0; i < int(n); i++ {
		sum += y[iy] * x[ix]
		ix += incX
		iy += incY
	}
	return sum
}
