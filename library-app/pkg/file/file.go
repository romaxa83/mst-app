package file

import "os"

func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
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
