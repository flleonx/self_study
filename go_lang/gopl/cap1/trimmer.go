package main

import (
	"regexp"
	"strings"
)

func trimmer(str string) string {
	regex := regexp.MustCompile(`\s+`)
	trimmedStr := regex.ReplaceAllString(str, " ")
	res := strings.Replace(trimmedStr, "\n", " ", -1)
	return res
}
