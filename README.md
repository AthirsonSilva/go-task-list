# Task list

Simple Web Restful API made with Golang, MongoDB and Chi framework

## Prerequisites

- Go 1.22 or higher
- MongoDB 

## Project Features

| Feature                           | Status       |
|-----------------------------------|--------------|
| Consistent API design             | 🟢 Ready     |
| Use of DTOs                       | 🟢 Ready     |
| Authentication with JWT           | 🟢 Ready     |
| Caching with Redis                | 🟢 Ready     |
| Data generation                   | 🟢 Ready     |
| Documentation with SwaggerUI      | 🟢 Ready     |
| Pagination, sorting and searching | 🟢 Ready     |
| Mailing service                   | 🟢 Ready     |
| Layered architecture              | 🟢 Ready     |
| Error Handling                    | 🟢 Ready     |
| API versioning                    | 🟢 Ready     |
| CSV and PDF exporting             | 🔴 Not Ready |
| AWS S3 service integration        | 🟢 Ready     |
| File upload and download          | 🟢 Ready     |
| Rate Limiting                     | 🟢 Ready     |
| Data Encryption                   | 🟢 Ready     |
| Asynchronous/ background tasks    | 🟢 Ready     |
| Logging                           | 🟢 Ready     |
| CI/ CD with Docker and Railway    | 🟢 Ready     |


## Getting Started

#### Clone the repository:

```bash
git clone https://github.com/athirsonsilva/go-task-list.git
```

#### Navigate to the project directory:

```bash
cd go-task-list
```

#### Install the dependencies:

```bash
go mod download
```

#### Build and run the project:

```bash
go build -o app ./cmd/server/main.go && ./app
```

#### The API will start running on http://localhost:8080.

## Documentation

The API documentation is available at http://localhost:8080/swagger/index.html

![Swagger UI](swagger.png)
