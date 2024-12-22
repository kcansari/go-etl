include .env

run:
	go run main.go

build:
	go build -o go-etl main.go

# Generate schema with improved error handling and options
generate-schema:
	@echo "Cleaning existing models..."
	@ls ./models/*.go | grep -v "types.go" | xargs rm -f
	@echo "Generating new models..."
	xo schema $(PG_URI)?sslmode=disable -o ./models 
		--go-custom-type="NullYear:sql.NullInt16"
		--go-custom-type="Tsvector:string"
		--go-custom-type="Trigger:string" 
		--go-custom-type-pkg="database/sql"
	@echo "Schema generation completed successfully"

# Help command
help:
	@echo "Available commands:"
	@echo "  make generate-schema  - Generate models from database schema"
	@echo "  make run             - Start the project"
	@echo "  make build           - Build the project"