package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/joshua-temple/goconf/pkg/goconf"
)

var rootCmd = &cobra.Command{
	Use:   "goconf",
	Short: "CLI for config generation and constant updates",
}

var configGenCmd = &cobra.Command{
	Use:   "generate [-t <target-path>] [-o <out-path>] [-p <package>]",
	Short: "Generate configuration constants",
	Run: func(cmd *cobra.Command, args []string) {
		if err := goconf.GenerateConstants(targetPath, outPath, pkg, backup); err != nil {
			fmt.Println("Error running generate:", err)
			os.Exit(1)
		}
	},
}

var (
	backup, dryRun                             bool
	directories                                []string
	oldFile, newFile, targetPath, pkg, outPath string
)

var updateConstCmd = &cobra.Command{
	Use:   "update -n <new> -o <old> [-d | -b] [--dirs <directories>]",
	Short: "Update constant usages based on new vs old .go files",
	Args:  cobra.MinimumNArgs(3),
	Run: func(cmd *cobra.Command, args []string) {
		if err := goconf.UpdateConstants(oldFile, newFile, dryRun, backup, directories); err != nil {
			fmt.Println("Error running update:", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&backup, "backup", "b", false, "Backup files before updating")

	updateConstCmd.PersistentFlags().BoolVarP(&dryRun, "dry-run", "d", false, "Dry run mode: show modifications without writing changes")
	updateConstCmd.Flags().StringVarP(&newFile, "new", "n", "", "New constant .go file to evaluate")
	updateConstCmd.Flags().StringVarP(&oldFile, "old", "o", "", "Old constant .go file to evaluate")
	updateConstCmd.Flags().StringSliceVar(&directories, "dirs", []string{}, "Directories to search for .go files to update")
	if err := updateConstCmd.MarkFlagRequired("new"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if err := updateConstCmd.MarkFlagRequired("old"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rootCmd.AddCommand(updateConstCmd)

	configGenCmd.Flags().StringVarP(&targetPath, "target-path", "t", ".", "Target directory (or file) for config files to generate the constants")
	configGenCmd.Flags().StringVarP(&outPath, "out-path", "o", ".", "Output directory (or file) for the generated constants")
	configGenCmd.Flags().StringVarP(&pkg, "package", "p", "config", "Package name for the generated constants")
	rootCmd.AddCommand(configGenCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
