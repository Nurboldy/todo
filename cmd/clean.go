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

var cleanCmd = &cobra.Command{
	Use:   "clean",
	Short: "remove all done tasks",
	RunE: func(cmd *cobra.Command, args []string) error {
		w, err := os.Create(path.GetPath() + "_")
		if err != nil {
			return err
		}
		defer w.Close()
		f, err := os.Open(path.GetPath())
		if err != nil {
			return nil
		}
		defer f.Close()
		br := bufio.NewReader(f)
		for {
			b, _, err := br.ReadLine()
			if err != nil {
				if err != io.EOF {
					return err
				}
				break
			}
			line := string(b)
			if !strings.HasPrefix(line, "-") {
				_, err = fmt.Fprintf(w, "%s\n", line)
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
	RootCmd.AddCommand(cleanCmd)
}
