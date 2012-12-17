// Copyright 2012 Pavel Vershinin. All rights reserved.
// master-dev@inbox.ru
// Use of this source code is governed by a BSD-style
package web

import (
    "bytes"
    "encoding/base64"
)

func Base64Encode(data string) string {
    var buf bytes.Buffer
    encoder := base64.NewEncoder(base64.StdEncoding, &buf)
    encoder.Write([]byte(data))
    encoder.Close()
    return buf.String()
}

