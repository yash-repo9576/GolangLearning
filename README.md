# GolangLearning
Golang Learning Repository - CRUD APP BookStore

1. Get all Books - GET:
```bash
$ curl localhost:8080/books
```

2. Create a Book - POST:
```bash
$ curl localhost:8080/create-book --include --header "Content-Type: application/json" -d @body.json --request "POST"
```
