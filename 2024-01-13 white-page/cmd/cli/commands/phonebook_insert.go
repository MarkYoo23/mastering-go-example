package commands

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"regexp"
	"strings"
	"white-page/internal/di"
	"white-page/internal/entries"
)

func init() {
	rootCmd.AddCommand(insertCmd)
}

var insertCmd = &cobra.Command{
	Use: commandInsert,
	Run: func(cmd *cobra.Command, args []string) {
		service, err := di.GetService[entries.EntryRepository]()
		if err != nil {
			panic(fmt.Errorf("error : %v", err))
		}

		var name = args[0]
		var surname = args[1]
		var tel = args[2]

		match, err := regexp.Match(`^\d{3}-\d{3,4}-\d{4}$`, []byte(tel))
		if err != nil {
			panic(err)
		}

		if !match {
			panic(errors.New("tel format error"))
		}

		tel = strings.ReplaceAll(tel, "-", "")

		result, err := service.Add(name, surname, tel)
		if err != nil {
			panic(fmt.Errorf("db error : %v", err))
		}

		fmt.Println(fmt.Sprintf("entry : %v", result))
	},
}
