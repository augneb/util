package utils

import (
	"regexp"
)

const (
	mobilePattern = `^1[34578][0-9]{9}$`
	mailPattern   = `^[a-z0-9A-Z]+([\-_\.][a-z0-9A-Z]+)*@([a-z0-9A-Z]+(-[a-z0-9A-Z]+)*?\.)+[a-zA-Z]{2,4}$`
)

var (
	mobileRegexp = regexp.MustCompile(mobilePattern)
	mailRegexp   = regexp.MustCompile(mailPattern)
)

func IsMobile(s interface{}) bool {
	switch s.(type) {
	case []byte:
		return mobileRegexp.Match(s.([]byte))
	default:
		return mobileRegexp.MatchString(s.(string))
	}
}

func IsMail(s interface{}) bool {
	switch s.(type) {
	case []byte:
		return mailRegexp.Match(s.([]byte))
	default:
		return mailRegexp.MatchString(s.(string))
	}
}
