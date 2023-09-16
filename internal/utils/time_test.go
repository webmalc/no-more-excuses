package utils

import (
	"testing"
	"time"
	"webmalc/no-more-excuses/common/test"

	"github.com/stretchr/testify/assert"
)

func TestGetWeekdayFromStr(t *testing.T) {
	weekday, err := GetWeekdayFromStr("mon")
	assert.Nil(t, err)
	assert.Equal(t, time.Monday, weekday)

	weekday, _ = GetWeekdayFromStr("sun")
	assert.Equal(t, time.Sunday, weekday)

	_, err = GetWeekdayFromStr("invalid")
	assert.Error(t, err)
}

func TestGetDurationsFromStrRange(t *testing.T) {
	start, end, err := GetDurationsFromStrRange("07:00-10:00")
	assert.Nil(t, err)
	assert.Equal(t, time.Hour*7, start)
	assert.Equal(t, time.Hour*10, end)

	start, end, err = GetDurationsFromStrRange("07:00-23:59")
	assert.Nil(t, err)
	assert.Equal(t, time.Hour*7, start)
	assert.Equal(t, time.Hour*23+time.Minute*59, end)

	_, _, err = GetDurationsFromStrRange("invalid")
	assert.ErrorContains(t, err, "HH:MM-HH:MM")

	_, _, err = GetDurationsFromStrRange("99:00-00:00")
	assert.ErrorContains(t, err, "hour > 23")

	_, _, err = GetDurationsFromStrRange("23:00-00:00")
	assert.ErrorContains(t, err, "begin time >")
}

func TestGetDurationsFromStr(t *testing.T) {
	duration, err := GetDurationsFromStr("07:00")
	assert.Nil(t, err)
	assert.Equal(t, time.Hour*7, duration)

	duration, err = GetDurationsFromStr("21:59")
	assert.Nil(t, err)
	assert.Equal(t, time.Hour*21+time.Minute*59, duration)

	_, err = GetDurationsFromStr("99:00")
	assert.ErrorContains(t, err, "hour > 23")

	_, err = GetDurationsFromStr("12:99")
	assert.ErrorContains(t, err, "minutes > 59")

	_, err = GetDurationsFromStr("invalid")
	assert.ErrorContains(t, err, "HH:MM")
}

func TestMain(m *testing.M) {
	test.Run(m)
}
