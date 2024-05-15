# Configuration Guide - Go Server with PostgreSQL

This guide shows how to set up a Go server using Gorilla mux and a PostgreSQL database, whit step-by-step instruction on configuring the development enviroment.

## Prerequisites

Make sure you have the following programs installed before you begin:

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)

To start the project:

- [Air](https://github.com/cosmtrek/air/)
- [Gorm](https://gorm.io/)

## Step 1: Setting up the Go server

1. Install Gorilla mux, a powerful routing library in Go:
```shell
go get -u github.com/gorilla/mux
```

2. Create a Go server using Gorilla mux. Example:

```go
package main

import (
    "fmt"
    "net/http"

    "github.com/gorilla/mux"

    func main(){
        router := mux.NewRouter()
        reuter.handlerFunc("/", HomeHandler)
        http.ListenAndServe(":3000", router)
    }

    func HomeHandler(w http.ResponseWriter, r *http.request){
        fmt.Fprintf(w, "Bienvenido al servidor Go!")
    }
)
```

3. Configure the port on which the server listens to port 3000.

## Step 2: Setting up Air 
1. Install Air, a tool to automatically restart the server whenever changes are detected in the code:
```shell
go install github.com/cosmtrek/air@latest
```

2. Run Air to start the server:
```shell
Air
```

## Step 3: Configuring the database:

1. Download a PostgreSQL images using Docker:
```shell
docker pull postgres:alpine
```

2. Create the PostgreSQL container:
```shell
docker run -d --name postgres-container -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 5432:5432 postgres:alpine
```

3. Enter the PostgreSQL container:
```shell
docker exec -it postgres-container bash
```

4. Enter the database server:
```shell
psql -U root --password
```

5. Create the database:
```sql
CREATE DATABASE gorm;
```

6. Verify that the database has been created:
```sql
\l
```

7. To see the tables created in the database, enter the database and execute:
```sql
\c gorm
\dt
```

## Extra Step: 
- Compile the project for Linux:
```bash
GOOS=linux GOARCH=amd64 go build -o my_program_linux main.go
```

- Compile the project for macOS:
```bash
GOOS=darwin GOARCH=amd64 go build -o my_program_macos main.go
```

- Compile the project for Windows:
```bash
GOOS=windows GOARCH=amd64 go build -o my_program_win.exe main.go
```
