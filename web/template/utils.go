package template

import (
	"fmt"
	"strings"
	"time"
)

func formatContactNumber(phone string) string {
	phone = strings.TrimPrefix(phone, "+55")

	if len(phone) != 11 {
		return phone
	}

	return fmt.Sprintf("(%s) %s-%s", phone[:2], phone[2:7], phone[7:])
}

func dateDiffInYears(date time.Time) int64 {
	return int64(time.Now().Sub(date).Hours() / 24 / 365)
}

func dateDiffInMonths(date time.Time) int64 {
	return int64(time.Now().Sub(date).Hours() / 24 / 30)
}
