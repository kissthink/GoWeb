// Copyright 2012 Pavel Vershinin. All rights reserved.
// master-dev@inbox.ru
// Use of this source code is governed by a BSD-style
package php

import (
	"os"
	"strings"
	"io/ioutil"
	"net/http"
	"syscall"
)
/* ============================================================================================ */
func BaseName(path string, suffix string) string {
	stat, err := os.Stat(path)
	setErr(err)
	if err != nil {
		return ""
	}
	path = stat.Name()
	if suffix != "" {
		index := strings.LastIndex(path, suffix)
		if index+len(suffix) == len(path) {
			path = SubStr(path, 0, index)
		}
	}
	return path
}
/* ============================================================================================ */
func ChGrp(filename string, group int) bool {
	err := os.Chown(filename, -1, group)
	setErr(err)
	if err == nil {
		return true
	}
	return false
}
/* ============================================================================================ */
func ChMod(filename string, mode int) bool {
	err := syscall.Chmod(filename, uint32(mode))
	if err == nil {
		return true
	}
	return false
}
/* ============================================================================================ */
func ChOwn(filename string, user int) bool {
	err := os.Chown(filename, user, -1)
	if err == nil {
		return true
	}
	return false
}
/* ============================================================================================ */
func Copy(source string, dest string) bool {
	file := FileGetContents(source)
	return FilePutContents(dest, file)
}
/* ============================================================================================ */
func DirName(path string) string {
	//@TODO Должна таки быть в Go своя функция!??
	pathArr := strings.Split(path, "/")
	if len(pathArr) == 1 {
		return pathArr[0]
	}
	if pathArr[len(pathArr)-1] == "" {
		return "/"+pathArr[len(pathArr)-3]
	} 
	return "/"+pathArr[len(pathArr)-2]
}
/* ============================================================================================ */
func FileExists(filename string) bool {
	stat, err := os.Stat(filename)
	setErr(err)
	if stat == nil && err != nil {
		return false
	}
	return true
}
/* ============================================================================================ */
func File(filename string) []string {
	file := FileGetContents(filename)
	return strings.Split(file, "\n")
}
/* ============================================================================================ */
func FileMTime(filename string) int64 {
	stat, err := os.Stat(filename)
	setErr(err)
	if err != nil {
		return 0
	}
	time := stat.ModTime()
	return time.Unix()
}
/* ============================================================================================ */
func FileSize(filename string) int64 {
	stat, err := os.Stat(filename)
	setErr(err)
	if err != nil {
		return 0
	}
	return stat.Size()
}
/* ============================================================================================ */
func FileType(filename string) string {
	stat, err := os.Stat(filename)
	setErr(err)
	if err != nil {
		return ""
	}
	if stat.IsDir() {
		return "dir"
	}
	return "file"
}
/* ============================================================================================ */
func IsDir(filename string) bool {
	if FileType(filename) == "dir" {
		return true
	}
	return false
}
/* ============================================================================================ */
func IsFile(filename string) bool {
	if FileType(filename) == "file" {
		return true
	}
	return false
}
/* ============================================================================================ */
func Rename(oldName string, newName string) bool {
	err := os.Rename(oldName, newName)
	if err == nil {
		return true
	}
	return false
}
/* ============================================================================================ */
func UnLink(filename string) bool {
	err := os.Remove(filename)
	if err == nil {
		return true
	}
	return false
}
/* ============================================================================================ */
func MkDir(dirname string, mode int, recursive bool) bool {
	if mode == -1 {
		mode = 0777
	}
	if recursive == true {
		err = os.MkdirAll(dirname, 0777)
	} else {
		err = os.Mkdir(dirname, 0777)
	}
	setErr(err)
	if err == nil {
		return true
	}
	return false
}
/* ============================================================================================ */
func RmDir(dirname string, all bool) bool {
	if all == true {
		err = os.RemoveAll(dirname)
	} else {
		err = os.Remove(dirname)
	}
	setErr(err)
	if err == nil {
		return true
	}
	return false
}
/* ============================================================================================ */
func FilePutContents(filename string, data string) bool {
	err := ioutil.WriteFile(filename, []byte(data), 0775)
	setErr(err)
	if err == nil {
		return true
	}
	return false
}
/* ============================================================================================ */
func FileGetContents(filename string) string {
	var file []byte
	if strings.Index(filename, "http://") == -1 {
		file, err = ioutil.ReadFile(filename)
		setErr(err)
	} else {
		var res *http.Response
		res, err = http.Get(filename)
		setErr(err)
		defer res.Body.Close()
		file, err = ioutil.ReadAll(res.Body)
		setErr(err)
	}
	return string(file)
}
/* ============================================================================================ */




