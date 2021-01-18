package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/Nurboldy/todo/path"
	"github.com/spf13/cobra"
)

var undoCmd = &cobra.Command{
	Use:   "undo",
	Short: "Unmark task a complete",
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
			line := strings.TrimSpace(string(b))
			if match && strings.HasPrefix(line, "-") {
				_, err = fmt.Fprintf(w, "%s\n", line[1:])
				if err != nil {
					return err
				}
				fmt.Printf("Task undone: %s\n", line[1:])
			} else {
				_, err = fmt.Fprintf(w, "%s\n", line)
				if err != nil {
					return err
				}
			}
		}
		w.Close()
		f.Close()
		err = os.Remove(path.GetPath())
		if err != nil {
			return err
		}
		return os.Rename(path.GetPath()+"_", path.GetPath())

	},
}

func init() {
	RootCmd.AddCommand(undoCmd)
}
