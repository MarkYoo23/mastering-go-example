package main

import (
	"fmt"
	"os"
	"strings"
	"white-page/cmd/cli/commands"
	"white-page/internal/constants"
	"white-page/internal/db"
	"white-page/internal/di"
)

// https://manytools.org/hacker-tools/ascii-banner/

const banner = `

  █████████  ██████████   ██████   █████
 ███░░░░░███░░███░░░░███ ░░██████ ░░███ 
░███    ░░░  ░███   ░░███ ░███░███ ░███ 
░░█████████  ░███    ░███ ░███░░███░███ 
 ░░░░░░░░███ ░███    ░███ ░███ ░░██████ 
 ███    ░███ ░███    ███  ░███  ░░█████ 
░░█████████  ██████████   █████  ░░█████
 ░░░░░░░░░  ░░░░░░░░░░   ░░░░░    ░░░░░ 
-------------------------------------------------------------------------------
Write page application v{ApplicationVersion}, created by 'Saturday Day Night'
-------------------------------------------------------------------------------
{args}
`

func init() {
}

func fillTemplate() string {
	template := banner
	template = strings.ReplaceAll(banner, "{ApplicationVersion}", constants.Version)
	template = strings.ReplaceAll(banner, "{args}", fmt.Sprintf("args : %v", os.Args[1:]))

	return template
}

func main() {
	fmt.Println(fillTemplate())

	err := di.Initialize()
	if err != nil {
		panic(err)
	}

	dbContext, err := di.GetService[db.DbContext]()
	if err != nil {
		panic(err)
	}

	err = dbContext.Open()
	if err != nil {
		panic(err)
	}

	err = commands.Execute()
	if err != nil {
		panic(err)
	}

	err = dbContext.Close()
	if err != nil {
		panic(err)
	}
}
