package repository

import (
	"encoding/base64"
	"time"

	"github.com/rs/zerolog/log"
)

// t := time.Now()
// log.Println("current time:", t.Format(timeLayout))
const (
	// 2006-01-02T15:04:05.999999999Z07:00
	timeLayout = time.RFC3339Nano
)

// DecodeCursor will decode cursor from user for mysql
func DecodeCursor(encodedTime string) (t time.Time, err error) {
	if encodedTime == "" {
		t = time.Now()
		log.Print("t: ", t)
		return t, nil
	}

	timeByte, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, err
	}

	timeString := string(timeByte)
	t, err = time.Parse(timeLayout, timeString)
	return
}

// EncodeCursor will encode cursor from mysql to user
func EncodeCursor(t time.Time) string {
	timeString := t.Format(timeLayout)
	return base64.StdEncoding.EncodeToString([]byte(timeString))
}
