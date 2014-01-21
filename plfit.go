// Copyright 2014, Hǎiliàng Wáng. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package plfit

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"time"
)

func Fit(a []float64, xmin_ int) (alpha, xmin, l, d float64, err error) {
	params := []string{"-s", fmt.Sprint(time.Now().UnixNano()), "-b", "/dev/stdin"}
	if xmin_ > 0 {
		params = append(params, []string{"-m", fmt.Sprint(xmin_)}...)
	}
	cmd := exec.Command("plfit", params...)
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return
	}
	err = cmd.Start()
	if err != nil {
		return
	}
	for _, v := range a {
		_, err = fmt.Fprintln(stdin, v)
		if err != nil {
			return
		}
	}
	stdin.Close()
	buf, err := ioutil.ReadAll(stdout)
	if err != nil {
		return
	}
	var file, t string
	_, err = fmt.Sscanf(string(buf), "%s%s%f%f%f%f", &file, &t, &alpha, &xmin, &l, &d)
	if err != nil {
		return
	}
	buf, err = ioutil.ReadAll(stderr)
	if err != nil {
		return
	}
	p(string(buf))
	err = cmd.Wait()
	return
}
