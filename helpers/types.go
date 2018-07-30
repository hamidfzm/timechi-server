package helpers

import (
	"time"
	"fmt"
)

type JSONTime struct {
	time.Time
}

func (t JSONTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", t.UTC().Format(DateTimeUTCFormat))), nil
}
