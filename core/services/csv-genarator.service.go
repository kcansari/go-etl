package services

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go-etl/core"
)

type CsvGenerator struct {
	service core.KnowledgeGraphService
}

func NewCsvGenerator(service core.KnowledgeGraphService) *CsvGenerator {
	return &CsvGenerator{service: service}
}

func (g *CsvGenerator) GenerateFiles(ctx context.Context) error {
	outputDir := g.service.GetOutputDir()

	// Create directory structure
	nodeCsvsDir := filepath.Join(outputDir, "csvs", "nodes")
	edgeCsvsDir := filepath.Join(outputDir, "csvs", "edges")

	if err := os.MkdirAll(nodeCsvsDir, 0755); err != nil {
		return fmt.Errorf("failed to create nodes directory: %w", err)
	}
	if err := os.MkdirAll(edgeCsvsDir, 0755); err != nil {
		return fmt.Errorf("failed to create edges directory: %w", err)
	}

	// Generate node CSVs
	for _, resolver := range g.service.GetNodeResolvers() {
		if err := g.generateNodeCsv(ctx, resolver, nodeCsvsDir); err != nil {
			return err
		}
	}

	// Generate edge CSVs
	for _, resolver := range g.service.GetEdgeResolvers() {
		if err := g.generateEdgeCsv(ctx, resolver, edgeCsvsDir); err != nil {
			return err
		}
	}

	return nil
}

func (g *CsvGenerator) generateNodeCsv(ctx context.Context, resolver core.NodeResolver, dir string) error {
	start := time.Now()
	fmt.Printf("Generating CSV file for node: %s\n", resolver.GetNodeLabel())

	data, err := resolver.PrepareData(ctx)
	if err != nil {
		return fmt.Errorf("failed to prepare data for node %s: %w", resolver.GetNodeLabel(), err)
	}

	filename := filepath.Join(dir, fmt.Sprintf("%s.csv", resolver.GetNodeLabel()))
	if err := writeCsvFile(filename, data); err != nil {
		return err
	}

	fmt.Printf("Completed generating CSV for node %s in %v\n", resolver.GetNodeLabel(), time.Since(start))
	return nil
}

func (g *CsvGenerator) generateEdgeCsv(ctx context.Context, resolver core.EdgeResolver, dir string) error {
	// Similar to generateNodeCsv
	// Implementation omitted for brevity
	return nil
}

func writeCsvFile(filename string, data []map[string]interface{}) error {
	// Implementation of CSV writing logic
	// This would handle converting the map data to CSV format
	return nil
}
