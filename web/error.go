// Copyright 2012 Pavel Vershinin. All rights reserved.
// master-dev@inbox.ru
// Use of this source code is governed by a BSD-style
package web

import (
	"strings"
	"fmt"
)

var errList []string
var err     error

/* ============================================================================================ */
func setErr(err error) {
	if err != nil {
		errList = append(errList, err.Error())
	}
}
/* ============================================================================================ */
func Error(separator string) string {
	if separator == "" {
		separator = "\n"
	}
	return strings.Join(errList, separator)
}
/* ============================================================================================ */
func ShowError() {
	if len(errList) > 0 {
		fmt.Println(Error("\n"))
	}
}
/* ============================================================================================ */
func ClearError() {
	errList = make([]string, 0)
}