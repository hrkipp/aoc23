package util

import (
	"strconv"
	"strings"
)

func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}

func Atoi(str string) int {
	return Must(strconv.Atoi(str))
}

func SplitInts(str string, sep string) []int {
	strs := strings.Split(str, sep)
	var out []int
	for _, str := range strs {
		if str == "" {
			continue
		}
		out = append(out, Atoi(strings.TrimSpace(str)))
	}
	return out

}

func Lines(str string) []string {
	return strings.Split(strings.TrimSpace(str), "\n")
}
