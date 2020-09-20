# Portal-API sample project 

A simple project to test portal rest features using Golang + Gin + Mongodb

## Goal & Requirement 

```
Use the API to sign up
Use the API to log in

```

go mod init 

```
$ go get github.com/gin-gonic/gin # the web framework

$ go get github.com/joho/godotenv # environment variables

$ go get gopkg.in/mgo.v2 # mongo driver

```

Setup mongodb DB and authorization

```
export DATABASE_HOST=localhost
export DATABASE_NAME=portal_api
export DATABASE_USER=rully
export DATABASE_PWD=secret

```

## Running the application using makefile

```
make run
```

## Accessing (read Routing API)

- Application will be accessible on http://localhost:5000/api/v1