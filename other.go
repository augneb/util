package util

func If(cond bool, v1, v2 interface{}) interface{} {
	if cond {
		return v1
	}

	return v2
}

