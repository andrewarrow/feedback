package files

import (
	"io/ioutil"
	"os"
)

func SaveFile(name, data string) {
	os.Remove(name)
	ioutil.WriteFile(name, []byte(data), 0644)
}
