package files

import "io/ioutil"

func ReadFile(name string) string {
	b, _ := ioutil.ReadFile(name)
	return string(b)
}
