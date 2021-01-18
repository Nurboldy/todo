package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to yout task list",
	Run: func(cmd *cobra.Command, args []string) {
		// w, err := os.OpenFile("filename", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
		// if err != nil {
		// 	return err
		// }
		// defer w.Close()
		todo := strings.Join(args, " ")
		fmt.Println(todo)
		//fmt.Printf("Added \"%s\" to your todo list.\n", todo)
		//_, err = fmt.Fprintln(w, todo)
		//return err
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
