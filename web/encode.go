// Copyright 2012 Pavel Vershinin. All rights reserved.
// master-dev@inbox.ru
// Use of this source code is governed by a BSD-style
package web

import (
    "fmt"
    "bytes"
    "regexp"
    "strings"
    "crypto/md5"
    "crypto/sha1"
    "encoding/hex"
    "encoding/base64"
)
/* ============================================================================================ */
func Base64Encode(data string) string {
    var buf bytes.Buffer
    encoder := base64.NewEncoder(base64.StdEncoding, &buf)
    encoder.Write([]byte(data))
    encoder.Close()
    return buf.String()
}
/* ============================================================================================ */
func Base64Decode(data string) string {
   data = strings.Replace(data, "\r", "", -1)
   data = strings.Replace(data, "\n", "", -1)
   var buf = bytes.NewBufferString(data)
   var res bytes.Buffer
   decoder := base64.NewDecoder(base64.StdEncoding, buf)
   res.ReadFrom(decoder)
   return res.String()
}
/* ============================================================================================ */
func Md5(data string) string {
    var h = md5.New()
    h.Write([]byte(data))
    return fmt.Sprintf("%x", h.Sum(nil))
}
/* ============================================================================================ */
func Sha1(data string) string {
    var h = sha1.New()
    h.Write([]byte(data))
    return fmt.Sprintf("%x", h.Sum(nil))
}
/* ============================================================================================ */
func QuotedPrintableDecode(data string) string {
    data = strings.Replace(data, "\r", "", -1)
    data = strings.Replace(data, "=\n", "", -1)
    var reg, _  = regexp.Compile(`[=][0-9A-F]{2}`)
    for _, line := range reg.FindAllString(data, -1) {
        dst, err := hex.DecodeString(strings.Trim(line, "="))
        if err == nil {
            data = strings.Replace(data, line, string(dst), 1)
        }
    }
    return data
}
