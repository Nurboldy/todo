package cmd

import (
	"github.com/spf13/cobra"
)

// RootCmd root command
var RootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Task is a CLI task manager",
}
