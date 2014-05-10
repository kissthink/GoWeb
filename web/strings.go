// Copyright 2012 Pavel Vershinin. All rights reserved.
// master-dev@inbox.ru
// Use of this source code is governed by a BSD-style
package web

import (
    "strconv"
	"strings"
	"html"
	"html/template"
	"regexp"
	"net/url"
	"encoding/hex"
)

/* ============================================================================================ */
func Trim(str string) string {
	return LTrim(RTrim(str))
}
/* ============================================================================================ */
func LTrim(str string) string {
	return strings.TrimLeft(str, " \n\r\t")
}
/* ============================================================================================ */
func RTrim(str string) string {
	return strings.TrimRight(str, " \n\r\t")
}
/* ============================================================================================ */
func HtmlSpecialChars(str string) string {
	str = html.EscapeString(str)
	str = strings.Replace(str, `&#34;`, `&quot;`, -1)
	return str
}
/* ============================================================================================ */
func HtmlSpecialCharsDecode(str string) string {
	return html.UnescapeString(str)
}
/* ============================================================================================ */
func UrlEncode(str string) string {
	return template.URLQueryEscaper(str)
}
/* ============================================================================================ */
func UrlDecode(str string) string {
    res, err := url.QueryUnescape(str)
    setErr(err)
	if err == nil {
		return res
	}
	return str
}
/* ============================================================================================ */
func RawUrlDecode(str string) string {
    re, _ := regexp.Compile(`(?Ui)%[0-9A-F]{2}`)
    str = re.ReplaceAllStringFunc(str, func(s string) string {
        b, err := hex.DecodeString(s[1:])
        if err == nil {
            return string(b)
        }
        return s
    })
	return str
}
/* ============================================================================================ */
func SubStr(str string, start int, length int) string {
    if start < 0 {
        start = 0
    }
	if length == 0 {
		length = len(str)
	}
	strArr := strings.Split(str, "")
	i      := 0
	str     = ""
	for i=start; i<start+length; i++ {
		if i >= len(strArr) {
			break
		}
		str += strArr[i]
	}
	return str
}
/* ============================================================================================ */
func StrPos(haystack string, needle string, offset int) int {
	if offset > 0 {
		haystack = SubStr(haystack, offset, len(haystack))
	}
	return strings.Index(haystack, needle) + offset
}
/* ============================================================================================ */
func StrRPos(haystack string, needle string) int {
	return strings.LastIndex(haystack, needle)
}
/* ============================================================================================ */
func StrStr(haystack string, needle string) string {
	start := StrPos(haystack, needle, 0)
	if start == -1 {
		return ""
	}
	return SubStr(haystack, start, len(haystack))
}
/* ============================================================================================ */
func StrChr(haystack string, needle string) string {
	return StrStr(haystack, needle)
}
/* ============================================================================================ */
func StrIStr(haystack string, needle string) string {
	start := StrPos(strings.ToLower(haystack), strings.ToLower(needle), 0)
	return SubStr(haystack, start, len(haystack))
}
/* ============================================================================================ */
func StrRChr(haystack string, needle string) string {
	start := StrRPos(haystack, needle)
	if start == -1 {
		return ""
	}
	return SubStr(haystack, start, len(haystack))
}
/* ============================================================================================ */
func SubStrCount(haystack string, needle string) int {
	return strings.Count(haystack, needle)
}
/* ============================================================================================ */
func StrSpn(str1 string, str2 string) int {
	strArr1 := strings.Split(str1, "")
	strArr2 := strings.Split(str2, "")
	var i  int
	var j  int
	var ok bool
	for i=0; i<len(strArr1); i++ {
		ok = false
		for j=0; j<len(strArr2); j++ {
			if strArr1[i] == strArr2[j] {
				ok = true
				break
			}
		}
		if ok == false {
			break
		}
	}
	return i
}
/* ============================================================================================ */
func StrCSpn(str1 string, str2 string) int {
	strArr1 := strings.Split(str1, "")
	strArr2 := strings.Split(str2, "")
	var i  int
	var j  int
	var ok bool
	for i=0; i<len(strArr1); i++ {
		ok = false
		for j=0; j<len(strArr2); j++ {
			if strArr1[i] == strArr2[j] {
				ok = true
				break
			}
		}
		if ok == true {
			break
		}
	}
	return i
}
/* ============================================================================================ */
func StrLen(str string) int {
	return len(str)
}
/* ============================================================================================ */
func WordWrap(str string, width int, breakStr string, cut bool) string {
	strArr := strings.Split(str, "")
	str = ""
	var i, j int
	for i=0; i<len(strArr); i++ {
		j++
		if j >= width && (strArr[i] == " " || cut == true) {
			str += breakStr + strArr[i]
			j = 0
		} else {
			str += strArr[i]		
		}
	}
	return str
}
/* ============================================================================================ */
func StrReplace(from string, to string, str string) string {
	return strings.Replace(str, from, to, -1)
}
/* ============================================================================================ */
func SubStrReplace(str string, replacement string, start int, length int) string {
	if length < 0 {
		length = len(str)
	}
	strArr := strings.Split(str, "")
	str = ""
	var i int
	for i=0; i<len(strArr); i++ {
		if i == start {
			str += replacement
		} else if i > start && i < start + length {
			continue
		} else {
			str += strArr[i]
		}
	}
	return str
}
/* ============================================================================================ */
func StrTr(str string, from string, to string) string {
	if len(from) > len(to) {
		from = SubStr(from, 0, len(to))
	}
	strArr  := strings.Split(str, "")
	fromArr := strings.Split(from, "")
	toArr   := strings.Split(to, "")
	var i int
	var j int
	for i=0; i<len(strArr); i++ {
		for j=0; j<len(fromArr); j++ {
			if fromArr[j] == strArr[i] {
				strArr[i] = toArr[j]
				break
			}
		}
	}
	return strings.Join(strArr, "")
}
/* ============================================================================================ */
func StripSlashes(str string) string {
	twoSlashes := "###TWO_SLASHES###"
	for StrPos(str, twoSlashes, 0) > -1 {
		twoSlashes += "#"
	}
	str = strings.Replace(str, `\\`, twoSlashes, -1)
	str = strings.Replace(str, `\`, "", -1)
	str = strings.Replace(str, twoSlashes, `\\`, -1)
	return str
}
/* ============================================================================================ */
func AddSlashes(str string) string {
	str = strings.Replace(str, `\`, `\\`, -1)
	str = strings.Replace(str, `"`, `\"`, -1)
	str = strings.Replace(str, `'`, `\'`, -1)
	str = strings.Replace(str, "`", "\\`", -1)
	return str
}
/* ============================================================================================ */
func QuoteMeta(str string) string {
	return regexp.QuoteMeta(str)
}
/* ============================================================================================ */
func StrRev(str string) string {
	strArr := strings.Split(str, "")
	str = ""
	var i int
	for i=len(strArr)-1; i>=0; i-- {
		str += strArr[i]
	}
	return str
}
/* ============================================================================================ */
func StrRepeat(str string, num int) string {
	return strings.Repeat(str, num)
}
/* ============================================================================================ */
func StrPad(str string, padLength int, padStr string, padType int) string {
	if len(str) >= padLength {
		return str
	}
	if padType == 1 { // right
		for padLength > len(str) {
			if len(str) + len(padStr) > padLength {
				padStr = SubStr(padStr, 0, padLength-len(str))
			}
			str += padStr
		}
	} else if padType == 2 { //left
		for padLength > len(str) {
			if len(str) + len(padStr) > padLength {
				padStr = SubStr(padStr, 0, padLength-len(str))
			}
			str = padStr + str
		}
	} else if padType == 3 { // both
		for padLength > len(str) {
			if len(str) + len(padStr) > padLength {
				padStr = SubStr(padStr, 0, padLength-len(str))
			}
			str += padStr
			if len(str) + len(padStr) > padLength {
				padStr = SubStr(padStr, 0, padLength-len(str))
			}
			str = padStr + str
		}
	}
	return str
}
/* ============================================================================================ */
func ChunkSplit(str string, chunkLen int, end string) string {
    var rStr = []rune(str+" ")
    var rEnd = []rune(end)
    var result []rune

    if chunkLen <= 0 {
        chunkLen = 76
    }
    if len(rEnd) == 0 {
        rEnd = []rune("\r\n")
    }

    for i:=0; i<len(rStr); i+=chunkLen {
        from := i
        to   := i+chunkLen
        if to >= len(rStr) {
            to = len(rStr)-1
        }
        result = append(result, rStr[from:to]...)
        result = append(result, rEnd...)
    }

	return strings.TrimSpace(string(result))
}
/* ============================================================================================ */
func Explode(arg string, str string, maxlimit int) []string {
	if maxlimit <= 0 {
		return strings.Split(str, arg)
	}
	return strings.SplitN(str, arg, maxlimit)
}
/* ============================================================================================ */
func Implode(arg string, array []string) string {
	return strings.Join(array, arg)
}
/* ============================================================================================ */
func Join(arg string, array []string) string {
	return strings.Join(array, arg)
}
/* ============================================================================================ */
func IntJoin(a []int, sep string) string {
    var buf []string
    for _, val := range a {
        buf = append(buf, strconv.Itoa(val))
    }
    return strings.Join(buf, sep)
}
/* ============================================================================================ */
func IntSplit(s, sep string) []int {
    var result []int
    for _, val := range strings.Split(strings.Trim(s, "{}"), sep) {
        iVal, err := strconv.Atoi(strings.TrimSpace(val))
        if err == nil {
            result = append(result, iVal)
        }
    }
    return result
}
/* ============================================================================================ */
func Int64Join(a []int64, sep string) string {
    var buf []string
    for _, val := range a {
        buf = append(buf, strconv.FormatInt(val, 10))
    }
    return strings.Join(buf, sep)
}
/* ============================================================================================ */
func Int64Split(s, sep string) []int64 {
    var result []int64
    for _, val := range strings.Split(strings.Trim(s, "{}"), sep) {
        iVal, err := strconv.ParseInt(val, 10, 64)
        if err == nil {
            result = append(result, iVal)
        }
    }
    return result
}
/* ============================================================================================ */
func ParceURL(strUrl string) *url.URL {
	result, err := url.Parse(strUrl)
	setErr(err)
	return result
}
/* ============================================================================================ */
func StrToLower(str string) string {
	return strings.ToLower(str)
}
/* ============================================================================================ */
func StrToUpper(str string) string {
	return strings.ToUpper(str)
}
/* ============================================================================================ */
func UcFirst(str string) string {
	first := SubStr(str, 0, 1)
	last  := SubStr(str, 1, 0)
	return strings.ToUpper(first)+strings.ToLower(last)
}
/* ============================================================================================ */
func UcWords(str string) string {
	strArr := strings.Split(str, " ")
	for i:=0; i<len(strArr); i++ {
		strArr[i] = UcFirst(strArr[i])
	}
	return strings.Join(strArr, " ")
}
/* ============================================================================================ */
func StripTags(s string) string {
    reg, _ := regexp.Compile(`(?Uis)<script[^>]*>.*</script>`)
    s = string(reg.ReplaceAll([]byte(s), []byte("")))
    reg, _ = regexp.Compile(`(?Uis)<.*>`)
    return string(reg.ReplaceAll([]byte(s), []byte("")))
}









