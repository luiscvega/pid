package pid

import (
	"io/ioutil"
	"os"
	"strconv"
)

func Write(path string) error {
	return ioutil.WriteFile(path, []byte(strconv.Itoa(os.Getpid())), 0644)
}
