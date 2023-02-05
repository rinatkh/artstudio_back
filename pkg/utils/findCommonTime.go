package utils

import (
	consts "github.com/rinatkh/artstudio_back/internal/constants"
	"github.com/rinatkh/artstudio_back/pkg/constants"
)

func bigger(first, second int64) int64 {
	if first >= second {
		return first
	}
	return second
}
func smaller(first, second int64) int64 {
	if first <= second {
		return first
	}
	return second
}

func FindCommonTime(firstStart, firstEnd, secondStart, secondEnd int64) (bool, int64, int64) {
	if firstStart >= secondEnd || firstEnd <= secondStart {
		return false, 0, 0
	}
	first := bigger(firstStart, secondStart)
	second := smaller(firstEnd, secondEnd)
	return true, first, second
}

func RemoveDuplicate[T string | int64](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}

func IsTime15MinDuration(args ...int64) error {
	for _, num := range args {
		if num%consts.DurationForInputTime != 0 {
			return constants.ErrTimeNot15
		}
	}
	return nil
}
