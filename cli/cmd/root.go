package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

const URL = "http://localhost:8000"

var rootCmd = &cobra.Command{
	Use:   "task",
	Short: "Use to create a task tracker",
	Long:  `This is the root cmd that will help to execute the diffrent functionality of the task tracker`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Println("error while executing the rootcmd: ", err)
		os.Exit(1)
	}
}

func init() {
}
