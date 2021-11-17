package main

import (
	"flag"
	"fmt"
	"log"
	"manifold/internal/build"
	"os"
)

func main() {
	// TODO: nice usage is required :)
	flag.Usage = func() {
		usage := `Usage: %s build [path]

build [path]      - build project or workspace located in current directory or on path if defined
`

		_, _ = fmt.Fprintln(flag.CommandLine.Output(), fmt.Sprintf(usage, os.Args[0]))
	}

	os.Exit(realMain())
}

func realMain() int {
	switch {
	case len(os.Args) > 2 && os.Args[1] == "build":
		var path string
		if len(os.Args) == 2 {
			path, _ = os.Getwd()
		} else {
			path = os.Args[2]
		}

		err := build.Build(path)

		if err != nil {
			log.Println(fmt.Sprintf("build failed: %v", err))
			return -1
		}

		log.Println("done")
		return 0

	default:
		flag.Usage()
		return 0
	}
}
