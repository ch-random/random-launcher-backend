package repository

import (
	"encoding/base64"
	"time"

	"github.com/rs/zerolog/log"
	validator "gopkg.in/go-playground/validator.v9"
)

const (
	// 2006-01-02T15:04:05.999999999Z07:00
	timeLayout = time.RFC3339Nano
)

// string (timeLayout) -> time.Time
func DecodeCursor(encodedTime string) (t time.Time, err error) {
	now := time.Now()
	log.Printf("current time: %v", now.Format(timeLayout))

	if encodedTime == "" {
		t = time.Now()
		log.Printf("t: %v", t)
		return
	}

	timeByte, err := base64.StdEncoding.DecodeString(encodedTime)
	if err != nil {
		return time.Time{}, err
	}

	timeString := string(timeByte)
	t, err = time.Parse(timeLayout, timeString)
	return
}

// time.Time -> string (timeLayout)
func EncodeCursor(t time.Time) string {
	timeString := t.Format(timeLayout)
	return base64.StdEncoding.EncodeToString([]byte(timeString))
}

func NewValidator() *validator.Validate {
	v := validator.New()
	return v
}
