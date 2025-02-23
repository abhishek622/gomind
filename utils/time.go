package utils

import (
	"time"

	"github.com/mergestat/timediff"
)

func FormatTime(timestamp string) string {
	t, _ := time.Parse(time.RFC3339, timestamp)
	return timediff.TimeDiff(t)
}
