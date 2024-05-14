# Guía de Configuración - Servidor Go con PostgreSQL

Se muestra cómo configurar un servidor Go utilizando Gorilla mux y una base de datos PostgreSQL, con instrucciones paso a paso sobre cómo configurar el entorno de desarrollo.

## Requisitos previos

Asegúrate de tener instalado los siguientes programas antes de comenzar:

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/)
- [Air](https://github.com/cosmtrek/air/)
- [Gorm](https://gorm.io/)

## Paso 1: Configuración del servidor Go

1. Instalar Gorilla Mux, una poderosa librería para enrutamiento en Go:
```shell
go get -u github.com/gorilla/mux
```

2. Crea un servidor Go utilizando Gorilla Mux. Ejemplo:

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

3. Configurar el puerto en el que escucha el servidor en el puerto 3000

## Paso 2: Configuración de Air

1. Instalar Air, una herramienta para reiniciar automáticamente el servidor cada vez que se detectan cambios en el código:
```shell
go install github.com/cosmtrek/air@latest
```

2. Ejecutar Air para iniciar el servidor:
```shell
Air
```

## Paso 3: Configurar la base de datos

1. Descargar una imagen de PostgreSQL mediante Docker:
```shell
docker pull postgres:alpine
```

2. Crea el contenedor de PostgreSQL:
```shell
docker run -d --name postgres-container -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 5432:5432 postgres:alpine
```

3. Ingresar al contenedor de PostgreSQL:
```shell
docker exec -it postgres-container bash
```

4. Ingresar al servidor de la base de datos:
```shell
psql -U root --password
```

5. Crea la base de datos:
```sql
CREATE DATABASE gorm;
```

6. Revisa que se haya creado la base de datos:
```sql
\l
```

7. Para ver las tablas creadas en la base de datos, ingresa a la base de datos y ejecuta:
```sql
\c gorm
\dt
```