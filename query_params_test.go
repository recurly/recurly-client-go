package recurly

import (
	"testing"
	"time"
)

func TestFormatBool(test *testing.T) {
	t := &T{test}

	trueToString := formatBool(true)
	falseToString := formatBool(false)
	t.Assert(trueToString, "true", "formatBool()")
	t.Assert(falseToString, "false", "formatBool()")
}

func TestFormatTime(test *testing.T) {
	t := &T{test}
	time := time.Now()
	formattedTime := formatTime(time)
	t.Assert(formattedTime, formatTime(time), "formatTime()")
}
