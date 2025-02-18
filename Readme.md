# Go-Gin-Try

This is a sample project based on tutorial https://go.dev/doc/tutorial/web-service-gin which is created using Go and Gin framework. We have lined further enhancement to make it more production ready.

## Features

- [x] Controller defined in separate directory
- [x] Model defined in seperate directory
- [x] Routes defined in seperate directory
- [ ] Services defined in seperate directory (Coming Soon)
- [ ] Utils  (Coming Soon, planned structure)
- [ ] Config (Coming Soon, planned structure)
- [ ] Middlewares (Coming Soon, planned structure)

## Project Structure

```
.
├── main.go
├── controllers
│   ├── albumController.go
├── models
│   ├── album.go
├── routes
│   ├── routes.go
├── README.md
```

## Clone and Run

To clone and run this project, follow these steps:

1. Clone the repository:
    ```
    git clone https://github.com/codebysachin/go-gin-try.git
    ```

2. Navigate to the project directory:
    ```
    cd go-gin-try
    ```

3. Build and run the project:
    ```
    go run main.go
    ```

## Sample Curl request
```curl http://localhost:8080/albums/2```
