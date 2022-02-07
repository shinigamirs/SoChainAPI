package domain

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"reflect"
	"strconv"
	"time"
)

//Time is struct for Custom time
type Time struct {
	time.Time
}

// UnmarshalJSON to convert timestamp to time.Time format
func (t *Time) UnmarshalJSON(b []byte) (err error) {
	// converting the bytes to int64
	timestamp, err := strconv.ParseInt(string(b), 10, 64)
	if timestamp == 0 {
		log.Panic("Timestamp is nil")
		return
	}
	t.Time = time.Unix(timestamp, 0)
	return
}

// MarshalJSON to convert timestamp to time.Time format
func (t Time) MarshalJSON() (b []byte, err error) {
	if !reflect.DeepEqual(t, time.Time{}) {
		// formatting the time to dd-mm-yyyy hh:mm
		formattedTime := fmt.Sprintf("\"%d-%d-%d %d:%d\"", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
		return []byte(formattedTime), nil
	} else {
		return []byte(""), nil
	}
}
