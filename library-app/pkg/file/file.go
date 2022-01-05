package file

import "os"

func Exists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func Remove(filename string) error {
	if err := os.Remove(filename); err != nil {
		return err
	}
	return nil
}
