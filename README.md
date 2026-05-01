# Books API

A simple RESTful API for managing books, built with Go, Gin framework, GORM, and MySQL.

## Features

- Full CRUD operations for books
- Concurrent search functionality
- MySQL database with GORM
- Docker support for easy deployment
- Environment-based configuration

## Tech Stack

- **Go** 1.26.2
- **Gin** - HTTP web framework
- **GORM** - ORM library
- **MySQL** 8 - Database
- **Docker** - Containerization

## Prerequisites

- Go 1.26.2 or higher
- Docker and Docker Compose (for containerized deployment)
- MySQL 8 (if running locally without Docker)

## Project Structure

```
go_task/
├── config/
│   └── db_config.go      # Database connection configuration
├── controllers/
│   └── book_controller.go # API endpoint handlers
├── models/
│   └── book.go           # Book data model
├── main.go               # Application entry point
├── Dockerfile            # Docker configuration
├── docker-compose.yml    # Docker Compose configuration
├── go.mod                # Go module dependencies
└── .env                  # Environment variables (create this)
```

## Quick Start

### Using Docker (Recommended)

1. **Clone the repository**
   ```bash
   cd go_task
   ```

2. **Run with Docker Compose**
   ```bash
   docker-compose up --build
   ```

3. **Access the API**
   ```
   http://localhost:8080
   ```

The application and MySQL database will start automatically.

### Local Development

1. **Install dependencies**
   ```bash
   go mod download
   ```

2. **Create `.env` file**
   ```env
   DB_USER=root
   DB_PASSWORD=root
   DB_HOST=localhost
   DB_NAME=booksdb
   ```

3. **Set up MySQL database**
   ```sql
   CREATE DATABASE booksdb;
   ```

4. **Run the application**
   ```bash
   go run main.go
   ```

The API will be available at `http://localhost:8080`.

## API Endpoints

### Books

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/books` | Get all books |
| GET | `/books/:id` | Get a book by ID |
| POST | `/books` | Create a new book |
| PUT | `/books/:id` | Update a book |
| DELETE | `/books/:id` | Delete a book |
| GET | `/books/search?q=keyword` | Search books by title or description |

### Book Model

```json
{
  "bookId": "string",
  "authorId": "string",
  "publisherId": "string",
  "title": "string",
  "publicationDate": "string",
  "isbn": "string",
  "pages": 0,
  "genre": "string",
  "description": "string",
  "price": 0.0,
  "quantity": 0
}
```

## Example Requests

### Create a Book

```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{
    "bookId": "B001",
    "authorId": "A001",
    "publisherId": "P001",
    "title": "The Go Programming Language",
    "publicationDate": "2015-11-16",
    "isbn": "978-0134190440",
    "pages": 400,
    "genre": "Technology",
    "description": "A comprehensive guide to Go",
    "price": 39.99,
    "quantity": 50
  }'
```

### Get All Books

```bash
curl http://localhost:8080/books
```

### Get Book by ID

```bash
curl http://localhost:8080/books/B001
```

### Update a Book

```bash
curl -X PUT http://localhost:8080/books/B001 \
  -H "Content-Type: application/json" \
  -d '{
    "price": 34.99,
    "quantity": 45
  }'
```

### Delete a Book

```bash
curl -X DELETE http://localhost:8080/books/B001
```

### Search Books

```bash
curl "http://localhost:8080/books/search?q=programming"
```

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| DB_USER | Database username | root |
| DB_PASSWORD | Database password | root |
| DB_HOST | Database host | localhost |
| DB_NAME | Database name | booksdb |

## Docker Configuration

The `docker-compose.yml` file defines two services:

- **app**: The Go application (port 8080)
- **mysql**: MySQL database (port 3307 → 3306)

MySQL data is persisted using a Docker volume.

## Development

### Adding New Features

1. Define models in `models/`
2. Create controllers in `controllers/`
3. Register routes in `main.go`
4. Update database migrations if needed

### Running Tests

```bash
go test ./...
```

