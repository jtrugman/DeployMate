package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [flags] [components...]",
	Short: "Add deployment components to your project",
	Long:  `Add deployment components such as GitHub Actions workflows and Dockerfiles to your project.`,
	Run:   runAdd,
}

func init() {
	rootCmd.AddCommand(addCmd)
}

func runAdd(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		fmt.Println("Please specify at least one component to add (e.g., sandbox, production, dockerfile)")
		return
	}

	root := determineProjectRoot()

	for _, component := range args {
		switch component {
		case "sandbox":
			addWorkflow(root, "sandbox.yml")
		case "production":
			addWorkflow(root, "production.yml")
		case "dockerfile":
			addDockerfile(root)
		default:
			fmt.Printf("Unknown component: %s\n", component)
		}
	}
}

func determineProjectRoot() string {
	if projectRoot != "" {
		return projectRoot
	}

	// Try to find .git directory
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		os.Exit(1)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, ".git")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}
		dir = parent
	}

	fmt.Println("Could not determine project root. Please use --project-root flag.")
	os.Exit(1)
	return ""
}

func addWorkflow(root, filename string) {
	workflowsDir := filepath.Join(root, ".github", "workflows")
	err := os.MkdirAll(workflowsDir, 0755)
	if err != nil {
		fmt.Printf("Error creating workflows directory: %v\n", err)
		return
	}

	templatePath := filepath.Join("templates", "github_actions", filename)
	destPath := filepath.Join(workflowsDir, filename)

	err = copyFile(templatePath, destPath)
	if err != nil {
		fmt.Printf("Error adding workflow file: %v\n", err)
	} else {
		fmt.Printf("Added %s workflow to %s\n", filename, destPath)
	}
}

func addDockerfile(root string) {
	// For simplicity, we'll always use the Python Dockerfile template
	// You can extend this to choose between different Dockerfile templates
	templatePath := filepath.Join("templates", "dockerfiles", "python.dockerfile")
	destPath := filepath.Join(root, "Dockerfile")

	err := copyFile(templatePath, destPath)
	if err != nil {
		fmt.Printf("Error adding Dockerfile: %v\n", err)
	} else {
		fmt.Printf("Added Dockerfile to %s\n", destPath)
	}
}

func copyFile(src, dest string) error {
	input, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	err = os.WriteFile(dest, input, 0644)
	if err != nil {
		return err
	}

	return nil
}
