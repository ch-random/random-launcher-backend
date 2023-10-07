package utils

import (
	"os"
	"reflect"
	"runtime"
	"strings"
)

// https://stackoverflow.com/a/70535822
func GetFuncName(i interface{}) string {
	strs := strings.Split((runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()), ".")
	return strs[len(strs)-1]
}

func GetEnvOrDefault(key string, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}
