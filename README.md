# Go Albums API

A RESTful API for managing music albums built with Go and the Gin web framework. This project demonstrates clean architecture principles with proper separation of concerns using handlers, services, and route organization.

## 🏗️ Architecture

This project follows a layered architecture pattern:

```
┌─────────────────┐
│   HTTP Layer    │  ← Gin routes, middleware
├─────────────────┤
│  Handler Layer  │  ← HTTP concerns (parsing, validation, responses)
├─────────────────┤
│  Service Layer  │  ← Business logic, validation, orchestration
├─────────────────┤
│   Data Layer    │  ← In-memory storage (could be database)
└─────────────────┘
```

### Project Structure

```
├── main.go              # Application entry point
├── go.mod              # Go module definition
├── handlers/           # HTTP handlers (HTTP concerns only)
│   └── albums.go       # Album-related HTTP handlers
├── services/           # Business logic layer
│   └── albums.go       # Album business logic and validation
└── routes/             # Route configuration
    └── routes.go       # Route setup and middleware configuration
```

## 🚀 Features

- **RESTful API** for album management
- **Clean Architecture** with separation of concerns
- **Business Logic Validation**:
  - Title and artist cannot be empty
  - Price must be greater than 0
  - No duplicate albums (same title + artist)
- **Proper Error Handling** with specific business errors
- **Dependency Injection** for testability
- **Route Grouping** and organization
- **API Versioning** (`/api/v1/`)

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
- Git

### Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd learn-go
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Run the application**
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

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

## 🏛️ Architecture Decisions

### Why This Structure?

1. **Separation of Concerns**: Each layer has a single responsibility
   - Handlers: HTTP-specific logic only
   - Services: Business logic and validation
   - Routes: Configuration and dependency injection

2. **Testability**: Services can be unit tested without HTTP concerns

3. **Scalability**: Easy to add new resources (users, playlists, etc.)

4. **Maintainability**: Changes to business logic don't affect HTTP handling

### Design Patterns Used

- **Dependency Injection**: Services are injected into handlers
- **Repository Pattern**: Service acts as a repository for album data
- **Handler Pattern**: HTTP handlers delegate to business services
- **Error Mapping**: Business errors mapped to appropriate HTTP status codes

## 🔄 Future Enhancements

This project serves as a foundation for more advanced features:

- **Database Integration**: Replace in-memory storage with PostgreSQL/MongoDB
- **Authentication & Authorization**: JWT tokens, user roles
- **Middleware**: Logging, rate limiting, CORS
- **Testing**: Unit tests, integration tests
- **Configuration**: Environment-based configuration
- **Docker**: Containerization for deployment
- **API Documentation**: Swagger/OpenAPI documentation
- **Caching**: Redis for performance optimization

## 📚 Learning Resources

This project demonstrates key Go web development concepts:

- **Gin Framework**: HTTP routing and middleware
- **Go Modules**: Dependency management
- **Struct Methods**: Object-oriented patterns in Go
- **Error Handling**: Go's explicit error handling
- **JSON Handling**: Marshaling and unmarshaling
- **Clean Architecture**: Separation of concerns
- **RESTful Design**: HTTP methods and status codes

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