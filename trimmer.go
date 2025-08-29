package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
Trim the path on your shell prompt
*/

func main() {
	usageError := `Usage: "trimmer path/to/trim [n]"`
	home := os.Getenv("HOME")
	if len(os.Args) < 2 {
		panic(usageError)
	}
	path := os.Args[1]

	keepDirs := 1
	if len(os.Args) == 3 {
		n, err := strconv.Atoi(os.Args[2])
		if err != nil {
			panic(usageError)
		}
		keepDirs = n
	}

	path = strings.Replace(path, home, "~", 1)

	folders := strings.Split(path, "/")

	if len(folders) > 2 {
		path = ""
		for i, f := range folders {
			if f == "~" {
				path += f + "/"
				continue
			}

			if i >= len(folders)-keepDirs {
				path += f
				if i != len(folders)-1 {
					path += "/"
				}
			} else {
				if len(f) > 3 {
					path += f[:3] + "…/"
				} else {
					path += f + "/"
				}
			}
		}
	}

	fmt.Printf("%s\n", path)
}
