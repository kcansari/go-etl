package core

import "context"

// NodeResolver interface defines the contract for node resolvers
type NodeResolver interface {
	GetNodeLabel() string
	PrepareData(ctx context.Context) ([]map[string]interface{}, error)
	GetEdgeResolvers() []EdgeResolver
}

// EdgeResolver interface defines the contract for edge resolvers
type EdgeResolver interface {
	GetEdgeLabel() string
	PrepareData(ctx context.Context) ([]map[string]interface{}, error)
}

// KnowledgeGraphService interface defines the main service contract
type KnowledgeGraphService interface {
	GetDisplayName() string
	GetOutputDir() string
	GetNodeResolvers() []NodeResolver
	GetEdgeResolvers() []EdgeResolver
}
