// Copyright 2012 Pavel Vershinin. All rights reserved.
// master-dev@inbox.ru
// Use of this source code is governed by a BSD-style
package web

/* ============================================================================================ */
func InArrayInt(arr []int, needle int) bool {
    for _, item := range arr {
        if item == needle {
            return true
        }
    }
    return false
}

/* ============================================================================================ */
func InArrayString(arr []string, needle string) bool {
    for _, item := range arr {
        if item == needle {
            return true
        }
    }
    return false
}
