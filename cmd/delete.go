package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/Nurboldy/todo/path"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete the todo",
	RunE: func(cmd *cobra.Command, args []string) error {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Println("Failed to parse the argument", arg)
			} else {
				ids = append(ids, id)
			}
		}
		w, err := os.Create(path.GetPath() + "_")
		if err != nil {
			return err
		}
		defer w.Close()

		f, err := os.Open(path.GetPath())
		if err != nil {
			return err
		}

		defer f.Close()
		br := bufio.NewReader(f)
		for n := 1; ; n++ {
			b, _, err := br.ReadLine()
			if err != nil {
				if err != io.EOF {
					return err
				}
				break
			}
			match := false
			for _, id := range ids {
				if id == n {
					match = true
				}
			}
			if match {
				fmt.Printf("Todo deleted: %s\n", string(b))
			} else {
				_, err = fmt.Fprintf(w, "%s\n", string(b))
				if err != nil {
					return err
				}
			}
		}
		f.Close()
		w.Close()

		err = os.Remove(path.GetPath())
		if err != nil {
			return err
		}
		return os.Rename(path.GetPath()+"_", path.GetPath())
	},
}

func init() {
	RootCmd.AddCommand(deleteCmd)
}
