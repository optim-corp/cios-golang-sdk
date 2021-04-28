package convert

import (
	"strings"
	"unicode/utf8"
)

func Nano2Milli(nanoTimeStamp interface{}) string {
	nanoTimeStampStr := MustStr(nanoTimeStamp)
	length := utf8.RuneCountInString(nanoTimeStampStr)
	if length < 6 {
		return nanoTimeStampStr
	}
	return nanoTimeStampStr[:length-6]
}
func Nano2MilliN(nanoTimeStamp interface{}) int64 {
	return MustInt64(nanoTimeStamp) / 1000000
}

func Milli2Nano(milliTimestamp interface{}) string {
	return MustStr(milliTimestamp) + "000000"
}
func Milli2NanoN(milliTimestamp interface{}) int64 {
	return MustInt64(milliTimestamp) * 1000000
}
func range2StartEnd(timestampRange string, sep string) (startTimestamp, endTimestamp string) {
	if sp := strings.Split(timestampRange, sep); len(sp) >= 2 {
		startTimestamp, endTimestamp = sp[0], sp[1]
	} else if len(sp) == 1 {
		startTimestamp = sp[0]
	}
	return
}
func ToTimestampRange(start, end interface{}, sep string) string {
	return MustStr(start) + sep + MustStr(end)
}
