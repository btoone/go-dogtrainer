## Overview

A prototype of a simple Go service.

## Usage

```bash
go run server.go

curl -i -X GET http://localhost:8080/dogs
curl -i -X DELETE http://localhost:8080/dogs
```

## Tests

```bash
# Run all tests
go test ./...

# Run a specific test file
go test handlers/dog_test.go
```
