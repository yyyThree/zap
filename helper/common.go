package helper

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"reflect"
	"time"
)

func IsStruct(s interface{}) bool {
	v := reflect.ValueOf(s)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	// uninitialized zero value of a struct
	if v.Kind() == reflect.Invalid {
		return false
	}

	return v.Kind() == reflect.Struct
}

func StructToMd5(s interface{}) string {
	if !IsStruct(s) {
		return ""
	}

	js, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return Md5(string(js))
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func FormatDateNow() string {
	return time.Now().Format("2006-01-02")
}

func FormatDateNowBySlash() string {
	return time.Now().Format("2006/01/02")
}
