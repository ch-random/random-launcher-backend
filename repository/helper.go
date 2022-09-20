package repository

import (
	"encoding/base64"
	"log"
	"time"
)

// t := time.Now()
// log.Println("current time:", t.Format(TIME_LAYOUT))
const (
	// 2006-01-02T15:04:05.999999999Z07:00
	TIME_LAYOUT = time.RFC3339Nano
)

// DecodeCursor will decode cursor from user for mysql
func DecodeCursor(encodedTime string) (t time.Time, err error) {
	if encodedTime == "" {
		t = time.Now()
		log.Println("t:", t)
		return t, nil
	}

	timeByte, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, err
	}

	timeString := string(timeByte)
	t, err = time.Parse(TIME_LAYOUT, timeString)
	return
}

// EncodeCursor will encode cursor from mysql to user
func EncodeCursor(t time.Time) string {
	timeString := t.Format(TIME_LAYOUT)
	return base64.StdEncoding.EncodeToString([]byte(timeString))
}
