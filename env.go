package main

import (
	"fmt"
	"os"
	"strings"
)

func ParseEnv(text string) {
	Arr := strings.Split(strings.TrimSpace(text), "\n")
	fmt.Println(Arr)
	for _, s := range Arr {
		I := strings.Index(s, "=")
		os.Setenv(s[:I], s[I+1:])
	}
}
