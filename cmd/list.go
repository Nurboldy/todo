package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/Nurboldy/todo/path"
	"github.com/spf13/cobra"
)

const (
	doneMark   = "\u2610"
	undoneMark = "\u2611"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all of your tasks",
	RunE: func(cmd *cobra.Command, args []string) error {

		f, err := os.Open(path.GetPath())
		if err != nil {
			return err
		}
		defer f.Close()
		br := bufio.NewReader(f)
		n := 1
		for {
			b, _, err := br.ReadLine()
			if err != nil {
				if err != io.EOF {
					return err
				}
				break
			}
			line := string(b)
			if strings.HasPrefix(line, "-") {
				fmt.Printf("%s %d: %s\n", undoneMark, n, strings.TrimSpace(line[1:]))
			} else {
				fmt.Printf("%s %d: %s\n", doneMark, n, strings.TrimSpace(line))
			}
			n++
		}

		return err
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
