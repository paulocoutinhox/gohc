package warm

import (
	"time"
)

var (
	StartedAt int64
	WarmTime  int64
)

func InWarmTime() bool {
	currentTime := time.Now().UTC().UnixNano() / int64(time.Millisecond)

	if (currentTime - StartedAt) > WarmTime {
		return false
	}

	return true
}
