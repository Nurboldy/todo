package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/Nurboldy/todo/path"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a task to yout task list",
	RunE:  add,
}

func add(cmd *cobra.Command, args []string) error {
	w, err := os.OpenFile(path.GetPath(), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}
	defer w.Close()
	todo := strings.Join(args, " ")
	fmt.Printf("Added \"%s\" to your todo list.\n", todo)
	_, err = fmt.Fprintln(w, todo)
	return err
}

func init() {
	RootCmd.AddCommand(addCmd)
}
