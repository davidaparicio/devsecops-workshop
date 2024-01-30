# Workshop Snowcamp2024 Petstore API

You work in a software development team within a company that manages an online pet sale service called "Petstore." The Product Manager of the company has identified a need to improve the current system and wants to develop an API to make the service more accessible to external partners. Security for this API is a major concern, and you have decided to use the 42Crunch tools to ensure its security throughout the development process.

The Petstore API implements the petstore.json specification with some conformance and security issues. You will be able to find them and to protect the API with 42Crunch tools (Audit/Scan/Firewall). There are more documentations about the Petstore API in the directory 'ressources' with an EDR diagramm and the Postman collection to execute some request.

## 1. Prerequisites

Before starting on the project, you need to have the following prerequisites installed on your machine : 

- Golang 1.18 or higher (https://go.dev/doc/install)
- An IDE
    - VSCode : (https://code.visualstudio.com/download)
    - IntelliJ : (https://www.jetbrains.com/help/idea/installation-guide.html)
- Docker (https://docs.docker.com/engine/install/)
- Docker-Compose
- Curl/Postman [Recommended to test the API]
- Oapi-codegen (https://github.com/deepmap/oapi-codegen) [Required to run 'make build' in ./api else run 'make build-go']

If you have some issues during the 'go mod tidy' for the 'make build'. You can perform these commands:
```Go
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
go get github.com/oapi-codegen/runtime
go get github.com/labstack/echo/v4@v4.11.1
go get github.com/labstack/echo/v4/middleware@v4.11.1
```

## 2. Building and running the API

With all the prerequisites installed, you are able now to run the api: 

```
cd ./api
make build
./api
```

When running the last command you should see the following:

```
   ____    __
  / __/___/ /  ___
 / _// __/ _ \/ _ \
/___/\__/_//_/\___/ v4.8.0
High performance, minimalist Go web framework
https://echo.labstack.com
____________________________________O/_______
                                    O\
⇨ http server started on [::]:4010

```

[Optionnal] Update the host file to simplify requests to the API:

```
127.0.0.1    petstore-secured
127.0.0.1    petstore
```

The API is up and is listening on the port 4010. To test that everything is fine you can execute the command and get the response below:

```
~/dev/petstore/api$ curl 'http://petstore:4010/version'
{"commitId":"3271bd2a72002b6a730d6da541c5da355390c4ff","version":"v1.0.0"}
```

## Vizualize the specification (OpenAPI)

There are 3 differents ways to visualize the specification file.

### Binary

Download the binary in the following link :

https://github.com/swagger-api/swagger-ui/releases/latest

### UI

Open the browser, copy the following link and drag and drop the file './petstore.json':

https://github.com/swagger-api/swagger-ui/releases/latest


### Docker

Execute the following command to have an instance of Swagger UI listening on port 8080

```
docker run --rm -p 8080:8080 -e SWAGGER_JSON=/petstore/petstore.json -v $(pwd):/petstore swaggerapi/swagger-ui
```

## 42Crunch Tools

To be able to execute 42Crunch tools you must have an account in the platform. You can register here : (https://platform.42crunch.com/).

## Workshop 1 : Static Analysis [Audit] 

Head over to the ./audit directory to complete this section of the workshop.

## Workshop 2 : Dynamic Analysis [Scan]

Navigate to the ./scan directory in order to work on this segment of the workshop.

## Workshop 3 : Protection [FW]

To protect the API, follow the readme file present in the 'fw' directory