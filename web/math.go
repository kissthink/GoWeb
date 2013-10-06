// Copyright 2012 Pavel Vershinin. All rights reserved.
// master-dev@inbox.ru
// Use of this source code is governed by a BSD-style
package web

import (
	"math/rand"
	"time"
)
var initRand bool

/* ============================================================================================ */
func Rand(min int, max int) int {
	if min > max {
		return 0
	}
	if initRand == false {
		rand.Seed(time.Now().UnixNano())
		initRand = true
	}
	return rand.Intn(max-min+1) + min
}

