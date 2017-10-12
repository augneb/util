package utils

import (
	"fmt"
	"time"
	"strings"
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

func Println(v ...interface{}) {
	l := v[len(v)-1]

	pref := ""
	switch val := l.(type) {
	case string:
		switch val {
		case "blue":
			pref = "\033[44;37;1m"
		case "green":
			pref = "\033[42;37;1m"
		case "red":
			pref = "\033[41;37;1m"
		case "yellow":
			pref = "\033[43;37;1m"
		}
	}

	str := []string{Date("[m-d H:i:s]")}

	n := len(v)
	if pref != "" {
		n--
		v = v[:n]

		str = append(str, pref)
	}

	for i := 0; i < n; i++ {
		str = append(str, "%v")
	}

	if pref != "" {
		str = append(str, "\033[0m")
	}

	str = append(str, "\n")

	fmt.Printf(strings.Join(str, " "), v...)
}

