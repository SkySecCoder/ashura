package scanner

import (
	"os/exec"
	"strings"
	"strconv"
)

func Ulimit() int64 {
	out, err := exec.Command("ulimit", "-n").Output()
	if err != nil {
		panic(err)
	}
	stringOutput := strings.TrimSpace(string(out))
	ulimit, err := strconv.ParseInt(stringOutput, 10, 64)
	if err != nil {
		panic(err)
	}
	return ulimit
}
