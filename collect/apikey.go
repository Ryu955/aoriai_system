package collect

import (
	"io/ioutil"
	"strings"
)

func GetApiKey() string {
	data, err := ioutil.ReadFile("./.env")
	if err != nil {
		panic(err)
	}
	return "token " +strings.TrimRight(string(data), "\n")
}
