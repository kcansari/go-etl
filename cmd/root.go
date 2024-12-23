package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "go-etl",
		Short: "A CLI tool for ETL operations and knowledge graph generation",
		Long: `A CLI application that helps generate and manage knowledge graphs,
perform ETL operations, and generate CSV files for nodes and edges.`,
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
