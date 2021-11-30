package main

import (
	"fmt"
	"github.com/mitchellh/cli"
	"sort"
	"strings"
)

func helpFunc(cmdFactories map[string]cli.CommandFactory) string {
	content := `%s

Commands:
%s

%s`

	commands := make([]string, 0)
	keys := make([]string, 0)

	for key := range cmdFactories {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		cmd, cmdErr := cmdFactories[key]()

		if cmdErr != nil {
			panic(cmdErr)
		}

		commands = append(commands, fmt.Sprintf("  %-13s %s", key, cmd.Synopsis()))
	}

	return fmt.Sprintf(content, usage(UsagePlaceholder, UsagePlaceholderArgs), strings.Join(commands, "\n"), version())
}

func version() string {
	return fmt.Sprintf(VersionBody, Version, GitHubLink)
}

func usage(name string, args string) string {
	return fmt.Sprintf(Usage, Name, name, args)
}
