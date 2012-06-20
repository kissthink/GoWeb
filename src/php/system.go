// Copyright 2012 Pavel Vershinin. All rights reserved.
// master-dev@inbox.ru
// Use of this source code is governed by a BSD-style
package php

import (
	"time"
)
/* ============================================================================================ */
func Sleep(seconds time.Duration) {
	time.Sleep(seconds * time.Second)
}
