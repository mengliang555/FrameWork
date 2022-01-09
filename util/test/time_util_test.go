package test

import (
	"FrameWork/util/time_util"
	"github.com/stretchr/testify/assert"
	"regexp"
	"testing"
)

const DatetimeLayout_RegularExpression = "[0-9]{4}-[0-9]{2}-[0-9]{2} [0-9]{2}:[0-9]{2}:[0-9]{2}"
const DatetimeLayoutRfc3339_RegularExpression = "[0-9]{4}-[0-9]{2}-[0-9]{2}[A-Z]{1}[0-9]{2}:[0-9]{2}:[0-9]{2}.([0-9]{2}:[0-9]{2})*"

func TestGetCurrentTime(t *testing.T) {
	match, err := regexp.Match(DatetimeLayout_RegularExpression, []byte(time_util.GetCurrentTimeToSecond()))
	assert.Nil(t,  err)
	assert.Equal(t, true, match)
}

func TestGetCurrentTimeToRec3339(t *testing.T) {
	println(time_util.GetCurrentTimeToRec3339())
	match, err := regexp.Match(DatetimeLayoutRfc3339_RegularExpression, []byte(time_util.GetCurrentTimeToRec3339()))
	assert.Nil(t,  err)
	assert.Equal(t, true, match)
}
