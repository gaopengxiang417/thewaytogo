package greeting

import (
	"time"
)

func IsAm() bool {
	now := time.Now()
	hour := now.Hour()
	if hour > 8 && hour < 12 {
		return true
	}
	return false
}

func IsPm() bool {
	hour := time.Now().Hour()
	if hour > 12 && hour < 20 {
		return true
	}
	return false
}

func IsEvening() bool {
	hour := time.Now().Hour()
	if (hour > 20 && hour < 24) || (hour >= 0 && hour <= 8) {
		return true
	}
	return false
}
