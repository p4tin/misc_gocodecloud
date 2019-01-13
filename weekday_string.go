// Code generated by "stringer -type Weekday enums_play.go"; DO NOT EDIT.

package main

import "strconv"

const _Weekday_name = "SundayMondayTuesdayWednesdayThursdayFridaySaturday"

var _Weekday_index = [...]uint8{0, 6, 12, 19, 28, 36, 42, 50}

func (i Weekday) String() string {
	if i < 0 || i >= Weekday(len(_Weekday_index)-1) {
		return "Weekday(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Weekday_name[_Weekday_index[i]:_Weekday_index[i+1]]
}
