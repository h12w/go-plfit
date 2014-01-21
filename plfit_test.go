// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plfit

import (
	"fmt"
	"io"
	"os"
	"testing"
)

func Test_Fit(t *testing.T) {
	f, err := os.Open("discrete_data.txt")
	c(err)
	defer f.Close()

	a := []float64{}
	for err != io.EOF {
		var v float64
		_, err = fmt.Fscanln(f, &v)
		a = append(a, v)
	}
	alpha, xmin, l, d, err := Fit(a, 1)
	c(err)
	p(alpha, xmin, l, d)
}

