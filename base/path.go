package base

import "os"

// MakePath -
func MakePath(dir string) error {
	_, err := os.Stat(dir)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		err1 := os.Mkdir(dir, os.ModePerm)
		if err1 != nil {
			return err1
		}

		return nil
	}
	return err
}
