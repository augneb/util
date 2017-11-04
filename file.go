package util

import (
	"os"
	"io"
)

func FileExists(path string) (bool, error) {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

func FileCopy(dest, source string) error {
	df, err := os.Create(dest)
	if err != nil {
		return err
	}

	f, err := os.Open(source)
	if err != nil {
		return err
	}

	_, err = io.Copy(df, f)
	f.Close()

	return err
}

func IsDir(path string) (bool, error) {
	i, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, err
	}

	return i.IsDir(), nil
}