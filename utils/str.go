package utils

import (
	"fmt"
	"strings"
)

func FillHeader(v string, headers map[string]string) {
	index := strings.Index(v, ":")
	if index < 0 {
		return
	}

	vIndex := index + 1
	if len(v) >= vIndex {
		value := strings.TrimPrefix(v[vIndex:], " ")

		if _, ok := headers[v[:index]]; ok {
			headers[v[:index]] = fmt.Sprintf("%s; %s", headers[v[:index]], value)
		} else {
			headers[v[:index]] = value
		}
	}
}
