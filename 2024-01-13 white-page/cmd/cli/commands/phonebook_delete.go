package commands

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
	"white-page/internal/di"
	"white-page/internal/entries"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use: commandDelete,
	Run: func(cmd *cobra.Command, args []string) {
		service, err := di.GetService[entries.EntryRepository]()
		if err != nil {
			panic(fmt.Errorf("error : %v", err))
		}

		idx, err := strconv.Atoi(args[0])
		if err != nil {
			panic(fmt.Errorf("convert string to int error : %v", err))
		}

		err = service.Delete(idx)
		if err != nil {
			panic(fmt.Errorf("db error : %v", err))
		}

		fmt.Println("entry deleted")
	},
}
