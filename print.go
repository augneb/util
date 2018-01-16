package util

import (
	"fmt"
	"github.com/liudng/godump"
)

func PrintWithColor(v ...interface{}) {
	l := v[len(v)-1]

	pref := ""
	switch val := l.(type) {
	case string:
		switch val {
		case "blue":
			pref = "\033[49;34;1m"
		case "green":
			pref = "\033[49;32;1m"
		case "red":
			pref = "\033[49;31;1m"
		case "yellow":
			pref = "\033[49;33;1m"
		}
	}

	str := ""

	switch v[0].(type) {
	case string:
		if v[0].(string) == "\n" {
			str += "\n"
			v = v[1:]
		}
	}

	n := len(v)
	if pref != "" {
		n--
		v = v[:n]
		str += pref
	}

	for i := 0; i < n; i++ {
		str += "%v "
	}

	str = str[:len(str)-1]

	if pref != "" {
		str += "\033[0m"
	}

	str += "\n"

	fmt.Printf(str, v...)
}

func Dump(v ...interface{}) string {
	l := len(v)
	s := false
	if l > 1 {
		if v, ok := v[l-1].(bool); ok && v {
			s = true
			l--
		}
	}

	r := ""
	for i := 0; i < l; i++ {
		if s {
			r += godump.Sdump(v[i])
		} else {
			godump.Dump(v[i])
		}
	}

	return r
}
