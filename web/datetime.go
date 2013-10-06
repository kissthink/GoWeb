// Copyright 2012 Pavel Vershinin. All rights reserved.
// master-dev@inbox.ru
// Use of this source code is governed by a BSD-style
package web

import (
	"time"
	"strconv"
	"strings"
)
/* ============================================================================================ */
func prepare(format string) string {
    var formatArr = strings.Split(format, "")
    for i, _ := range formatArr {
        switch formatArr[i] {
            case "d":
                formatArr[i] = "02"
            case "j":
                formatArr[i] = "_2"
            case "m":
                formatArr[i] = "01"
            case "n":
                formatArr[i] = "_1"
            case "t":
                formatArr[i] = "31"
            case "Y":
                formatArr[i] = "2006"
            case "y":
                formatArr[i] = "06"
            case "a":
                formatArr[i] = "pm"
            case "A":
                formatArr[i] = "PM"
            case "g":
                formatArr[i] = "3"
            case "G":
                formatArr[i] = "15"// Что-то нет, 24 часового формата без ведущего нуля (((
            case "H":
                formatArr[i] = "15"
            case "h":
                formatArr[i] = "03"
            case "i":
                formatArr[i] = "04"
            case "s":
                formatArr[i] = "05"
            case "O":
                formatArr[i] = "-0700"
            case "P":
                formatArr[i] = "-07:00"
            case "T":
                formatArr[i] = "MST"
            case "Z":
                formatArr[i] = "Z0700"
            case "D":
                formatArr[i] = "Mon"
            case "l":
                formatArr[i] = "Monday"
            case "F":
                formatArr[i] = "January"
            case "M":
                formatArr[i] = "Jan"
        }
    }
	return strings.Join(formatArr, "")
}
/* ============================================================================================ */
func Date(format string, timestamp int64) string {
	if timestamp == -1 {
		timestamp = time.Now().Unix()
	}
	t := time.Unix(timestamp, 0)
	return t.Format(prepare(format))
}
/* ============================================================================================ */
func Time() int64 {
	return time.Now().Unix()
}
/* ============================================================================================ */
func MicroTime(asFloat bool) string {
	var strTime string
	if asFloat == true {
		strTime = strconv.FormatInt(time.Now().Unix(), 10)+"."+strconv.Itoa(time.Now().Nanosecond())
	} else {
		strTime = strconv.Itoa(time.Now().Nanosecond())+" "+strconv.FormatInt(time.Now().Unix(), 10)
	}
	return strTime
}
/* ============================================================================================ */
func StrToTime(strTime string, format string) int64 {
	if format == "" {
		format = "2006-01-02 15:04:05"
	} else {
		format = prepare(format)
	}
	t, err := time.Parse(format, Trim(strTime))
	setErr(err)
	if err != nil {
		return 0
	}
	return t.Unix()
}
