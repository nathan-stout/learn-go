#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}🚀 Starting Albums API Development Environment${NC}"

# Check if .env file exists
if [ ! -f .env ]; then
    echo -e "${YELLOW}⚠️  No .env file found. Creating from template...${NC}"
    if [ -f env.template ]; then
        cp env.template .env
        echo -e "${GREEN}✅ Created .env file from template${NC}"
        echo -e "${YELLOW}📝 Please edit .env with your configuration if needed${NC}"
    else
        echo -e "${RED}❌ No env.template found. Please create .env file manually.${NC}"
        exit 1
    fi
fi

# Check if Docker is running
if ! docker info > /dev/null 2>&1; then
    echo -e "${RED}❌ Docker is not running. Please start Docker first.${NC}"
    exit 1
fi

# Start the development environment
echo -e "${YELLOW}📦 Building and starting containers...${NC}"
docker-compose up --build -d

# Wait for services to be ready
echo -e "${YELLOW}⏳ Waiting for services to be ready...${NC}"
sleep 5

# Check if services are running
if docker-compose ps | grep -q "Up"; then
    echo -e "${GREEN}✅ Development environment is ready!${NC}"
    echo -e "${GREEN}📝 API Server: http://localhost:8080${NC}"
    echo -e "${GREEN}🗄️  Database: localhost:5432${NC}"
    echo -e "${GREEN}📋 View logs: docker-compose logs -f${NC}"
    echo -e "${GREEN}🛑 Stop: docker-compose down${NC}"
else
    echo -e "${RED}❌ Something went wrong. Check logs: docker-compose logs${NC}"
    exit 1
fi 