package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Displays a list of the databases on the DBHub.io server.
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Returns a list of available databases",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Databases on %s\n\n", cloud)
		fmt.Println("Name\tSize\t\tLast Modified")
		fmt.Println("****\t****\t\t*************")

		list := getDBList()
		for _, j := range list {
			fmt.Printf("%s\t%d bytes\t%s\n", j.Database, j.Size, j.LastModified)
		}
		fmt.Println()
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
