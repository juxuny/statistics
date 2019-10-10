package statistics

import "os"

// create the directory if it is not exists
func TouchDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, 0775)
	}
	return nil
}
