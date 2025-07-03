.PHONY: dev dev-down dev-logs build test clean help setup setup-env install-dev

# Default target
help:
	@echo "Available commands:"
	@echo "  setup      - First time setup (creates .env, installs dependencies)"
	@echo "  dev        - Start development environment (database + app with hot reload)"
	@echo "  dev-down   - Stop development environment"
	@echo "  dev-logs   - Follow development logs"
	@echo "  build      - Build the application"
	@echo "  test       - Run tests"
	@echo "  clean      - Clean up containers and volumes"
	@echo "  db-only    - Start only the database"
	@echo "  app-local  - Run app locally (requires database to be running)"
	@echo "  setup-env  - Create .env file from template"
	@echo "  install-dev - Install development dependencies"

# Start full development environment
dev:
	@echo "🚀 Starting development environment..."
	docker-compose up --build

# Start development environment in background
dev-bg:
	@echo "🚀 Starting development environment in background..."
	docker-compose up --build -d
	@echo "✅ Development environment started!"
	@echo "📝 App: http://localhost:8080"
	@echo "🗄️  Database: localhost:5432"
	@echo "📋 Run 'make dev-logs' to see logs"

# Stop development environment
dev-down:
	@echo "🛑 Stopping development environment..."
	docker-compose down

# Follow logs
dev-logs:
	docker-compose logs -f

# Start only database (original workflow)
db-only:
	@echo "🗄️  Starting database only..."
	docker-compose up postgres -d

# Run app locally (requires database)
app-local:
	@echo "🏃 Running app locally..."
	go run main.go

# Build the application
build:
	@echo "🔨 Building application..."
	go build -o bin/albums-api .

# Run tests
test:
	@echo "🧪 Running tests..."
	go test ./...

# Clean up everything
clean:
	@echo "🧹 Cleaning up..."
	docker-compose down -v
	docker system prune -f
	rm -rf tmp/
	rm -rf bin/

# Install development dependencies
install-dev:
	@echo "📦 Installing development dependencies..."
	go install github.com/air-verse/air@latest

# Setup environment file
setup-env:
	@echo "⚙️  Setting up environment file..."
	@if [ ! -f .env ]; then \
		cp env.template .env; \
		echo "✅ Created .env file from template"; \
		echo "📝 Please edit .env with your configuration"; \
	else \
		echo "⚠️  .env file already exists"; \
	fi

# First time setup
setup:
	@echo "🚀 Setting up development environment..."
	@make setup-env
	@make install-dev
	@echo "✅ Setup complete! Run 'make dev' to start developing" 