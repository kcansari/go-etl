package cmd

import (
	"context"
	"fmt"
	"go-etl/core"
	"go-etl/core/services"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var (
	outputDir   string
	displayName string
)

var generateCmd = &cobra.Command{
	Use:   "csv-generate",
	Short: "Generate CSV files for nodes and edges",
	Long: `Generate CSV files for nodes and edges based on the configured resolvers.
The files will be organized in separate directories for nodes and edges.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Validate output directory
		if outputDir == "" {
			return fmt.Errorf("output directory is required")
		}

		// Create absolute path
		absOutputDir, err := filepath.Abs(outputDir)
		if err != nil {
			return fmt.Errorf("failed to resolve absolute path: %w", err)
		}

		// Ensure output directory exists
		if err := os.MkdirAll(absOutputDir, 0755); err != nil {
			return fmt.Errorf("failed to create output directory: %w", err)
		}

		// Initialize your service with resolvers
		service, err := core.NewKnowledgeGraphService(displayName, absOutputDir)
		if err != nil {
			return fmt.Errorf("failed to create knowledge graph service: %w", err)
		}

		// Create CSV generator
		generator := services.NewCsvGenerator(service)

		// Generate files
		if err := generator.GenerateFiles(context.Background()); err != nil {
			return fmt.Errorf("failed to generate CSV files: %w", err)
		}

		fmt.Printf("Successfully generated CSV files in: %s\n", absOutputDir)
		return nil
	},
}

func init() {
	// Add generate command to root command
	rootCmd.AddCommand(generateCmd)

	// Add flags for generate command
	generateCmd.Flags().StringVarP(&outputDir, "output", "o", "", "Output directory for generated files (required)")
	generateCmd.Flags().StringVarP(&displayName, "name", "n", "Knowledge Graph", "Display name for the knowledge graph")

}
