package util

import (
	"fmt"
	"time"
	"github.com/liudng/godump"
)

func Debug(msg string, time ...time.Time) {
	var d string
	if len(time) > 0 {
		d = Date("m-d H:i:s", time[0])
	} else {
		d = Date("m-d H:i:s")
	}

	fmt.Printf("[%s] %s\n", d, msg)
}

func Print(v ...interface{}) {
	printWithColor(false, false, v)
}

func PrintLog(v ...interface{}) {
	printWithColor(true, false, v)
}

func Println(v ...interface{}) {
	printWithColor(false, true, v)
}

func PrintlnLog(v ...interface{}) {
	printWithColor(true, true, v)
}

func printWithColor(date bool, eof bool, v []interface{}) {
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

	if date {
		str += Date("[m-d H:i:s] ")
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

	if eof {
		str += "\n"
	}

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
