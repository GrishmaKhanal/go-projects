# Simple Go Projects to Get Started with Golang

## 1. Simple HTTP Server
A basic example to handle HTTP POST requests.

### How to Run:
1. Run the server directly:  
   ```bash
   go run main.go
2. Build and execute the binary:
    ```bash
    go build main.go
    ./main    # On Linux/MacOS
    ./main.exe # On Windows

## 2. CRUD API
A simple CRUD API project using the gorilla/mux library.
gorilla/mux is a HTTP router and URL matcher for building web applications and APIs in Go

### How to Run:
1. Initialize the module
    ```bash
    go mod init example.com/crud-api
2. Install dependencies
    ```bash
    go get github.com/gorilla/mux@latest
    go mod tidy #clean up
