# Library Management API

A simple RESTful API built with Go and Gin framework for managing a library's book inventory. This API allows you to perform basic CRUD operations on books, including checking out and returning books.

## Features

- Get all books in the library
- Get a specific book by ID
- Create new books
- Remove books from the library
- Check out books (decrease quantity)
- Return books (increase quantity)

## Prerequisites

- Go 1.16 or higher
- Gin framework (`github.com/gin-gonic/gin`)

## Installation

1. Clone the repository
2. Install dependencies:
```bash
go mod init your-module-name
go get github.com/gin-gonic/gin
```

3. Run the application:
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
    "id": "4",
    "title": "New Book",
    "author": "Author Name",
    "quantity": 1
}
```
- Response: Success message

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

## Data Structure

Book object structure:
```go
type book struct {
    ID       string `json:"id"`
    Title    string `json:"title"`
    Author   string `json:"author"`
    Quantity int    `json:"quantity"`
}
```

## Example Usage

### Get all books
```bash
curl http://localhost:8080/books
```

### Create a new book
```bash
curl -X POST http://localhost:8080/create \
    -H "Content-Type: application/json" \
    -d '{"id":"4","title":"New Book","author":"Author Name","quantity":1}'
```

### Check out a book
```bash
curl -X PUT "http://localhost:8080/checkout?id=1"
```

### Return a book
```bash
curl -X PUT "http://localhost:8080/return?id=1"
```

## Error Handling

The API returns appropriate HTTP status codes and error messages:
- 200: Success
- 201: Created
- 202: Accepted
- 400: Bad Request
- 404: Not Found

## Contributing

Feel free to submit issues and enhancement requests.

## License

This project is licensed under the MIT License - see the LICENSE file for details.