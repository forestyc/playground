package time

import (
	"encoding/json"
	"strings"
	"time"
)

type Time time.Time

func (t Time) MarshalJSON() ([]byte, error) {
	formattedTime := time.Time(t).Format("2006-01-02 15:04:05")
	formattedTime, _ = strings.CutSuffix(formattedTime, " 00:00:00")
	return json.Marshal(formattedTime)
}
