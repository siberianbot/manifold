package main

import (
	"fmt"
	"log"
	"manifold/internal/build"
	"os"
)

func main() {
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
		println(help())
		return 0
	}
}
