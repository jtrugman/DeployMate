package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var projectRoot string
var isMonorepo bool

var rootCmd = &cobra.Command{
	Use:   "deploymate",
	Short: "Deploymate is a tool for managing deployment configurations",
	Long:  `Deploymate helps you set up and manage deployment pipelines and Dockerfiles for your projects.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&projectRoot, "project-root", "", "Specify the project root directory")
	rootCmd.PersistentFlags().BoolVar(&isMonorepo, "monorepo", false, "Specify if the project is a monorepo")

	// Add this line to include the add command
	rootCmd.AddCommand(addCmd)
}
