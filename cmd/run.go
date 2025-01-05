package cmd

import (
	"github.com/spf13/cobra"
	"github.com/taylormonacelli/howbob/run"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		path, _ := cmd.Flags().GetString("path")
		brewfile, _ := cmd.Flags().GetString("brewfile")
		checker, _ := cmd.Flags().GetString("checker")
		taps, _ := cmd.Flags().GetString("taps")
		run.Brewfile(path, brewfile, checker, taps)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.Flags().StringP("path", "p", "manifest.k", "Path to manifest.k")
	runCmd.Flags().StringP("brewfile", "b", "Brewfile", "Path to Brewfile")
	runCmd.Flags().StringP("checker", "c", "version_checker.sh", "Path to version_checker.sh")
	runCmd.Flags().StringP("taps", "t", "taps.sh", "Path to taps file")
}
