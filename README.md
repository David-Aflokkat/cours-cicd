# Introduction

Cats demo REST API used to manage a local database of 🐈.


# Dev

## Run locally

``` bash
go run .
```

Then you can browse to the Swagger UI on http://localhost:8080/swagger

## Swagger UI setup

Done following [Swagger official doc](https://github.com/swagger-api/swagger-ui/blob/master/docs/usage/installation.md#plain-old-htmlcssjs-standalone).

## Regenerate the OpenApi file

``` bash
npx swagger-cli bundle -o swagger-ui/openapi.json openapi.yml
```

## Docker

``` bash
docker build -t test2 .
```

# Testing
