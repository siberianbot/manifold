package main

import (
	"fmt"
	"log"
	"manifold/internal/backend"
	"os"
)

func main() {
	os.Exit(realMain())
}

func realMain() int {
	b := backend.NewBackend()

	switch {
	case len(os.Args) > 2 && os.Args[1] == "build":
		buildOptions := backend.BuildOptions{}

		if len(os.Args) == 2 {
			buildOptions.Path, _ = os.Getwd()
		} else {
			buildOptions.Path = os.Args[2]
		}

		err := b.Build(buildOptions)

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
