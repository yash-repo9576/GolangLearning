# Library Management API

A RESTful API built with Go, Gin framework, and MongoDB for managing a library's book inventory. This API allows you to perform basic CRUD operations on books, including checking out and returning books.

## Tech Stack

- Go 1.16 or higher
- Gin Web Framework
- MongoDB
- Docker & Docker Compose

## Project Structure
```
.
├── main.go
├── docker-compose.yml
├── database/
│   └── database.go
├── models/
│   └── book.go
├── handlers/
│   └── handlers.go
└── utils/
    └── utils.go
```

## Prerequisites

- Go 1.16 or higher
- Docker and Docker Compose
- Git

## Installation & Setup

1. Clone the repository:
```bash
git clone <your-repository-url>
cd <project-directory>
```

2. Install Go dependencies:
```bash
go mod init your-module-name
go get github.com/gin-gonic/gin
go get go.mongodb.org/mongo-driver/mongo
```

3. Start MongoDB using Docker:
```bash
docker-compose up -d
```

4. Run the application:
```bash
go run main.go
```

The server will start on `localhost:8080`

## API Endpoints

### GET `/books`
- Returns a list of all books in the library
- Response: Array of book objects

### GET `/book/:id`
- Returns a specific book by ID
- Parameters: `id` (path parameter)
- Response: Book object or error message

### POST `/create`
- Adds a new book to the library
- Request Body:
```json
{
    "title": "New Book",
    "author": "Author Name",
    "quantity": 1
}
```
- Response: Book object with generated ID

### PUT `/checkout`
- Checks out a book (decreases quantity by 1)
- Parameters: `id` (query parameter)
- Response: Updated book object or error message

### PUT `/return`
- Returns a book (increases quantity by 1)
- Parameters: `id` (query parameter)
- Response: Updated book object or error message

### DELETE `/delete/:id`
- Removes a book from the library
- Parameters: `id` (path parameter)
- Response: Success message or error message

## Database Configuration

MongoDB connection details:
- Host: localhost
- Port: 27017
- Username: admin
- Password: password
- Database: library
- Collection: books

## Docker Commands

Start containers:
```bash
docker-compose up -d
```

Stop containers:
```bash
docker-compose down
```

View logs:
```bash
docker logs library_mongodb
```

Reset data (warning: this will delete all data):
```bash
docker-compose down -v
docker-compose up -d
```

## Example API Usage

### Get all books
```bash
curl http://localhost:8080/books
```

### Create a new book
```bash
curl -X POST http://localhost:8080/create \
    -H "Content-Type: application/json" \
    -d '{"title":"New Book","author":"Author Name","quantity":1}'
```

### Check out a book
```bash
curl -X PUT "http://localhost:8080/checkout?id=<book_id>"
```

### Return a book
```bash
curl -X PUT "http://localhost:8080/return?id=<book_id>"
```

## Error Handling

The API returns appropriate HTTP status codes and error messages:
- 200: Success
- 201: Created
- 202: Accepted
- 400: Bad Request
- 404: Not Found
- 500: Internal Server Error

## Development

To run in development mode with hot reload:
```bash
go install github.com/cosmtrek/air@latest
air
```

## Testing

Run all tests:
```bash
go test ./...
```
