// Copyright 2012 Pavel Vershinin. All rights reserved.
// master-dev@inbox.ru
// Use of this source code is governed by a BSD-style
package php

import (
	"time"
	"strconv"
)
/* ============================================================================================ */
func prepare(format string) string {
	format = StrReplace("d", "02", format)
	format = StrReplace("D", "Mon", format)
	format = StrReplace("j", "_2", format)
	format = StrReplace("l", "Monday", format)
	format = StrReplace("F", "January", format)
	format = StrReplace("m", "01", format)
	format = StrReplace("M", "Jan", format)
	format = StrReplace("n", "_1", format)
	format = StrReplace("t", "31", format)
	format = StrReplace("Y", "2006", format)
	format = StrReplace("y", "06", format)
	format = StrReplace("a", "pm", format)
	format = StrReplace("A", "PM", format)
	format = StrReplace("g", "3", format)
	format = StrReplace("G", "15", format)// Что-то нет, 24 часового формата без ведущего нуля (((
	format = StrReplace("H", "15", format)
	format = StrReplace("h", "03", format)
	format = StrReplace("i", "04", format)
	format = StrReplace("s", "05", format)
	format = StrReplace("O", "-0700", format)
	format = StrReplace("P", "-07:00", format)
	format = StrReplace("T", "MST", format)
	format = StrReplace("Z", "Z0700", format)
	return format
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
	t, err := time.Parse(format, strTime)
	setErr(err)
	if err != nil {
		return 0
	}
	return t.Unix()
}
