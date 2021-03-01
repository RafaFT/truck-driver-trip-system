package entity

import (
	"fmt"
	"strconv"
	"time"
)

type BirthDate struct {
	time.Time
}

func NewBirthDate(birthDate time.Time) BirthDate {
	return BirthDate{
		birthDate,
	}
}

func (bd BirthDate) CalculateAge() int {
	now := time.Now()

	years := now.Year() - bd.Year()
	if years < 0 {
		return 0
	}

	birthMonthNDay, _ := strconv.Atoi(fmt.Sprintf("%d%d", int(bd.Month()), bd.Day()))
	baseDateMonthNDay, _ := strconv.Atoi(fmt.Sprintf("%d%d", int(now.Month()), now.Day()))

	if birthMonthNDay > baseDateMonthNDay {
		years--
	}

	return years
}
