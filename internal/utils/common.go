package utils

import (
	"fmt"
	"lark-hotel-server/internal/consts"
	"time"
)

const (
	LateShift  = "晚班"
	EarlyShift = "早班"
	MidShift   = "中班"

	DateFormat      = "2006-01-02"
	DateTimeFormat  = "2006-01-02 15:04:05"
	DateTimeFormat2 = "2006/01/02 15:04"
)

func GetShift() string {
	currentTime := time.Now()
	hour := currentTime.Hour()

	if hour >= 0 && hour < 8 {
		return LateShift
	} else if hour >= 8 && hour < 16 {
		return EarlyShift
	} else {
		return MidShift
	}
}

func WithTimeAtStartOfDate(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
}

func WithTimeAtEndOfDate(date time.Time) time.Time {
	return time.Date(date.Year(), date.Month(), date.Day(), 23, 59, 59, int(time.Nanosecond)-1, date.Location())
}

func FormatTime(t time.Time) string {
	return t.Format(DateTimeFormat2)
}

func JudgeRoomType(roomNumber string) consts.RoomType {
	for _, room := range consts.SingleBedArray {
		if room == roomNumber {
			return consts.SingleBed
		}
	}
	for _, room := range consts.DoubleBedArray {
		if room == roomNumber {
			return consts.DoubleBed
		}
	}
	for _, room := range consts.JuniorSuiteArray {
		if room == roomNumber {
			return consts.JuniorSuite
		}
	}
	for _, room := range consts.SmallMahjongSuiteArray {
		if room == roomNumber {
			return consts.SmallMahjongSuite
		}
	}
	for _, room := range consts.LargeSuiteArray {
		if room == roomNumber {
			return consts.LargeSuite
		}
	}
	return consts.None
}

func UniqueBillKey(name string, phone string) string {
	return fmt.Sprintf("%s-%s", name, phone)
}
