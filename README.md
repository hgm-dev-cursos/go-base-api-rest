# REST API example application

This is an example of a Golang REST API README

# Application ENVs

To find all necessary application envs, please take a look into *.env.example* file. To run locally, please copy all envs in *.env.example* into a new *.env* file

# Makefile Commands

### Install Dependencies

    make install

### Run the app

    make run

### Run the tests

    make tests

### Generate Mocks for Application Deps

    make mock-dependencies

### Generate Swagger Documentation

    make swagger-docs

# REST API

The REST API to the example app is described below.

## Swagger Documentation

[http://localhost:8000/docs](http://localhost:8000/docs)

## Create a new Example

### Request

`POST /api/example`

    curl --request POST \
        --url http://localhost:8000/api/example \
        --header 'Authorization: test-authorization-header' \
        --header 'Content-Type: application/json' \
        --data '{
            "name": "test"
        }'

### Response

    Content-Type: application/json
    Date: Tue, 13 Dec 2022 03:32:01 GMT
    Content-Length: 60

    {"id":"639b89588b70ecfe5f2dee2c","name":"test"}
