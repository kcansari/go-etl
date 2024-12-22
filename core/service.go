package core

import (
	"fmt"
)

// BaseKnowledgeGraphService provides a base implementation
type BaseKnowledgeGraphService struct {
	nodeResolvers []NodeResolver
	edgeResolvers []EdgeResolver
	displayName   string
	outputDir     string
}

// NewKnowledgeGraphService creates a new instance of the service
func NewKnowledgeGraphService(displayName, outputDir string, resolvers ...interface{}) (*BaseKnowledgeGraphService, error) {
	service := &BaseKnowledgeGraphService{
		displayName: displayName,
		outputDir:   outputDir,
	}

	// Validate and categorize resolvers
	for _, resolver := range resolvers {
		switch r := resolver.(type) {
		case NodeResolver:
			service.nodeResolvers = append(service.nodeResolvers, r)
			// Collect edge resolvers from node resolvers
			service.edgeResolvers = append(service.edgeResolvers, r.GetEdgeResolvers()...)
		default:
			return nil, fmt.Errorf("invalid resolver type: %T", resolver)
		}
	}

	return service, nil
}

func (s *BaseKnowledgeGraphService) GetDisplayName() string {
	return s.displayName
}

func (s *BaseKnowledgeGraphService) GetOutputDir() string {
	return s.outputDir
}

func (s *BaseKnowledgeGraphService) GetNodeResolvers() []NodeResolver {
	return s.nodeResolvers
}

func (s *BaseKnowledgeGraphService) GetEdgeResolvers() []EdgeResolver {
	return s.edgeResolvers
}
