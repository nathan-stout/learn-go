# Go Albums API

A RESTful API for managing music albums built with Go, Gin web framework, and PostgreSQL. This project demonstrates clean architecture principles with proper separation of concerns using handlers, services, repositories, and database layers.

## 🏗️ Architecture

This project follows a layered architecture pattern with the Repository pattern:

```
┌─────────────────┐
│   HTTP Layer    │  ← Gin routes, middleware
├─────────────────┤
│  Handler Layer  │  ← HTTP concerns (parsing, validation, responses)
├─────────────────┤
│  Service Layer  │  ← Business logic, validation, orchestration
├─────────────────┤
│ Repository Layer│  ← Data access abstraction
├─────────────────┤
│ Database Layer  │  ← PostgreSQL connection and migrations
└─────────────────┘
```

### Project Structure

```
├── main.go              # Application entry point
├── go.mod              # Go module definition
├── docker-compose.yml  # PostgreSQL setup for development
├── config/             # Configuration management
│   └── config.go       # Environment-based configuration
├── database/           # Database connection and migrations
│   └── connection.go   # PostgreSQL connection, migrations, seeding
├── repositories/       # Data access layer
│   └── album_repository.go # Album database operations
├── services/           # Business logic layer
│   └── albums.go       # Album business logic and validation
├── handlers/           # HTTP handlers (HTTP concerns only)
│   └── albums.go       # Album-related HTTP handlers
└── routes/             # Route configuration
    └── routes.go       # Route setup and dependency injection
```

## 🚀 Features

- **RESTful API** for album management
- **Clean Architecture** with separation of concerns
- **Repository Pattern** for data access abstraction
- **PostgreSQL Integration** with connection pooling
- **Database Migrations** and seeding
- **Environment-based Configuration**
- **Business Logic Validation**:
  - Title and artist cannot be empty
  - Price must be greater than 0
  - No duplicate albums (same title + artist)
- **Proper Error Handling** with specific business errors
- **Dependency Injection** for testability
- **Route Grouping** and organization
- **API Versioning** (`/api/v1/`)
- **Docker Compose** for easy PostgreSQL setup

## 📋 API Endpoints

| Method | Endpoint | Description | Request Body |
|--------|----------|-------------|--------------|
| `GET` | `/api/v1/albums` | Get all albums | None |
| `GET` | `/api/v1/albums/:id` | Get album by ID | None |
| `POST` | `/api/v1/albums` | Create new album | `AlbumRequest` |
| `DELETE` | `/api/v1/albums/:id` | Delete album | None |

### Data Models

#### Album
```json
{
  "id": "string",
  "title": "string",
  "artist": "string", 
  "price": number
}
```

#### AlbumRequest (for POST)
```json
{
  "title": "string",
  "artist": "string",
  "price": number
}
```

## 🛠️ Setup and Installation

### Prerequisites

- Go 1.24.3 or higher
- Docker and Docker Compose
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd learn-go
   ```

2. **Install dependencies and setup environment**
   ```bash
   go mod download
   make setup  # Creates .env file and installs dev dependencies
   ```

3. **Configure your environment (optional)**
   
   Edit the `.env` file created by setup if you need custom configuration:
   ```bash
   # .env file (automatically created by make setup)
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=postgres
   DB_PASSWORD=your_secure_password_here
   DB_NAME=albums_db
   DB_SSLMODE=disable
   PORT=8080
   ```

4. **Start the development environment**

   **Option A: One command with Docker Compose (Recommended)**
   ```bash
   make dev
   # or
   docker-compose up --build
   ```

   **Option B: Using the development script**
   ```bash
   ./dev.sh
   ```

   **Option C: Background with status**
   ```bash
   make dev-bg
   ```

   **Option D: Traditional approach (database + manual app start)**
   ```bash
   # Start database only
   make db-only
   # or
   docker-compose up postgres -d
   
   # Create .env file (if not using Docker Compose)
   # Then run app locally
   make app-local
   # or 
   go run main.go
   ```

The server will start on `http://localhost:8080` and automatically:
- Connect to PostgreSQL
- Run database migrations
- Seed initial data
- Enable hot reloading (when using Docker Compose)

### Available Development Commands

```bash
make help          # Show all available commands
make setup         # First time setup (creates .env, installs dependencies)
make dev           # Start full development environment
make dev-bg        # Start in background
make dev-down      # Stop development environment
make dev-logs      # View development logs
make db-only       # Start only the database
make app-local     # Run app locally (requires database)
make build         # Build the application
make test          # Run tests
make clean         # Clean up containers and volumes
make setup-env     # Create .env file from template
make install-dev   # Install development dependencies
```

## 📖 Usage Examples

### Get All Albums
```bash
curl http://localhost:8080/api/v1/albums
```

**Response:**
```json
[
  {
    "id": "1",
    "title": "Blue Train",
    "artist": "John Coltrane",
    "price": 56.99
  },
  {
    "id": "2", 
    "title": "Jeru",
    "artist": "Gerry Mulligan",
    "price": 17.99
  }
]
```

### Get Album by ID
```bash
curl http://localhost:8080/api/v1/albums/1
```

### Create New Album
```bash
curl -X POST http://localhost:8080/api/v1/albums \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Kind of Blue",
    "artist": "Miles Davis",
    "price": 49.99
  }'
```

### Delete Album
```bash
curl -X DELETE http://localhost:8080/api/v1/albums/1
```

## ⚠️ Error Handling

The API returns appropriate HTTP status codes and error messages:

### Business Validation Errors (400 Bad Request)
```bash
# Empty title
curl -X POST http://localhost:8080/api/v1/albums \
  -H "Content-Type: application/json" \
  -d '{"title":"","artist":"Test","price":10.99}'
```
**Response:**
```json
{
  "message": "title cannot be empty"
}
```

### Duplicate Album (409 Conflict)
```bash
curl -X POST http://localhost:8080/api/v1/albums \
  -H "Content-Type: application/json" \
  -d '{"title":"Blue Train","artist":"John Coltrane","price":29.99}'
```
**Response:**
```json
{
  "message": "album with this title and artist already exists"
}
```

### Not Found (404)
```bash
curl http://localhost:8080/api/v1/albums/999
```
**Response:**
```json
{
  "message": "album not found"
}
```

## 🧪 Testing

### Manual Testing with curl

Test all the validation rules:

```bash
# Test empty title validation
curl -X POST http://localhost:8080/api/v1/albums \
  -H "Content-Type: application/json" \
  -d '{"title":"","artist":"Test Artist","price":10.99}'

# Test empty artist validation  
curl -X POST http://localhost:8080/api/v1/albums \
  -H "Content-Type: application/json" \
  -d '{"title":"Test Album","artist":"","price":10.99}'

# Test invalid price validation
curl -X POST http://localhost:8080/api/v1/albums \
  -H "Content-Type: application/json" \
  -d '{"title":"Test Album","artist":"Test Artist","price":-5.99}'

# Test duplicate album validation
curl -X POST http://localhost:8080/api/v1/albums \
  -H "Content-Type: application/json" \
  -d '{"title":"Blue Train","artist":"John Coltrane","price":29.99}'
```

## 🗄️ Database

### PostgreSQL Schema

The application uses PostgreSQL with the following schema:

```sql
CREATE TABLE albums (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    artist VARCHAR(255) NOT NULL,
    price DECIMAL(10,2) NOT NULL CHECK (price > 0),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(title, artist)
);
```

### Database Management

- **Migrations**: Automatically run on startup
- **Seeding**: Initial data inserted if table is empty
- **Connection Pooling**: Configured for production use
- **Indexes**: Added on title and artist for performance

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_HOST` | PostgreSQL host | `localhost` |
| `DB_PORT` | PostgreSQL port | `5432` |
| `DB_USER` | Database user | `postgres` |
| `DB_PASSWORD` | Database password | *(required)* |
| `DB_NAME` | Database name | `albums_db` |
| `DB_SSLMODE` | SSL mode | `disable` |
| `PORT` | Server port | `8080` |

## Security

This project follows security best practices:

### Environment Variables
- **Never commit `.env` files** - they contain sensitive information
- Use `env.template` as a reference for required environment variables
- The `.gitignore` file prevents accidental commits of sensitive files

### Database Security
- All database queries use parameterized statements to prevent SQL injection
- Connection strings are built from environment variables
- Database credentials are never hardcoded in source code

### Development vs Production
- **Development**: Uses default passwords for convenience
- **Production**: Requires strong passwords and proper SSL configuration
- **Google Cloud**: Uses Cloud SQL with managed security

### Quick Security Setup
```bash
# Create your environment file
make setup-env

# Edit .env with secure values
nano .env  # or your preferred editor

# Never commit .env
git status  # should not show .env file
```

## 🏛️ Architecture Decisions

### Why This Structure?

1. **Separation of Concerns**: Each layer has a single responsibility
   - **Handlers**: HTTP-specific logic only
   - **Services**: Business logic and validation
   - **Repositories**: Data access abstraction
   - **Database**: Connection management and migrations

2. **Repository Pattern Benefits**:
   - **Testability**: Easy to mock data layer
   - **Flexibility**: Can switch databases without changing business logic
   - **Maintainability**: Data access logic centralized

3. **Dependency Injection**: Clean dependency flow from database → repository → service → handler

4. **Configuration Management**: Environment-based config for different deployment environments

### Design Patterns Used

- **Repository Pattern**: Abstracts data access
- **Dependency Injection**: Services and repositories injected into handlers
- **Layered Architecture**: Clear separation between HTTP, business, and data layers
- **Error Mapping**: Business errors mapped to appropriate HTTP status codes
- **Factory Pattern**: Constructor functions for creating instances

## 🔄 Future Enhancements

This project serves as a foundation for more advanced features:

- **Authentication & Authorization**: JWT tokens, user roles
- **Middleware**: Logging, rate limiting, CORS
- **Testing**: Unit tests, integration tests
- **Caching**: Redis for performance optimization
- **API Documentation**: Swagger/OpenAPI documentation
- **Monitoring**: Health checks, metrics, logging
- **CI/CD**: Automated testing and deployment
- **Google Cloud Integration**: Cloud SQL, Cloud Run deployment

## 📚 Learning Resources

This project demonstrates key Go web development concepts:

- **Gin Framework**: HTTP routing and middleware
- **PostgreSQL Integration**: Database connections, migrations, queries
- **Repository Pattern**: Data access abstraction
- **Clean Architecture**: Separation of concerns
- **Environment Configuration**: Managing different environments
- **Docker**: Containerized development environment
- **Go Modules**: Dependency management
- **Error Handling**: Go's explicit error handling
- **JSON Handling**: Marshaling and unmarshaling

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is for educational purposes.

---

**Built with ❤️ and Go**