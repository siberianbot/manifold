package main

import (
	"github.com/google/shlex"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
	"os/exec"
)

type projectDefinition struct {
	Name string `yaml:"name"`
}

type stepDefinition struct {
	// TODO: move this
	Cmd string `yaml:"cmd"`
}

type projectDocument struct {
	Project projectDefinition `yaml:"project"`
	Steps   []stepDefinition  `yaml:"steps"`
}

func main() {
	if len(os.Args) < 2 {
		log.Println("no args")
		os.Exit(0)

		return
	}

	file, err := os.ReadFile(os.Args[1])

	if err != nil {
		log.Println(err)
		os.Exit(-1)

		return
	}

	projectFile := projectDocument{}
	err = yaml.Unmarshal(file, &projectFile)

	if err != nil {
		log.Println(err)
		os.Exit(-2)

		return
	}

	log.Println("building", projectFile.Project.Name)

	for _, step := range projectFile.Steps {

		cmdSplit, _ := shlex.Split(step.Cmd)

		cmd := exec.Command("cmd", "/c")

		cmd.Args = append(cmd.Args, cmdSplit...)

		mw := io.MultiWriter(log.Writer())
		cmd.Stdout = mw
		cmd.Stderr = mw

		cmdErr := cmd.Run()

		if cmdErr != nil {
			log.Println("finished step with error:", cmdErr)
			os.Exit(-3)

			return
		} else {
			log.Println("finished step without errors")
		}
	}
}
