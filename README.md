# Go api example

An example CRUD api using Go 1.20.7 and the following dependencies and tools:

- Gin (HTTP router and handler)
- Zerolog (Logging)
- Gookit config (Configuration per environment)
- Go MongoDB driver (Persistence)
- Unit testing with internal testing tool and Testify for assertions and mocks
- Swagoo for api documentation (Swagger)
- A Dockerfile to create a Docker container with the Api ready to use

### Endpoints

GET: `http://localhost:9090/api/v1/users`

Finds all active users.

GET: `http://localhost:9090/api/v1/users/{id}`

Gets an user by its id. Returns 404 if the user was not found.

GET: `http://localhost:9090/api/v1/search`

Search users. Parameters:
- firstName: User first name (complete or initial characters)
- lastName: User last name (complete or initial characters)
- email: User email (complete or initial characters)
- page: Page number, starting from 1
- size: Page size, starting from 1

POST: `http://localhost:9090/api/v1/users`

Creates an user. Example request body:

`
{
    "firstName": "Foo",
    "lastName": "Bar",
    "email": "foobar@email.com"
}
`

Returns 201 is the user creation was successful.

PUT: `http://localhost:9090/api/v1/users/{id}`

Updates an existent user by it's id. Example request body:

`
{
    "firstName": "Foo",
    "lastName": "Bar",
    "email": "foobar@email.com"
}

Returns 200 with the updated user if it was successful. This endpoints also allows to active deleted (inactive) users.

DELETE: `http://localhost:9090/api/v1/users/{id}`

Deletes an existent user by it's id. The deletion is logical, so the user registry is preserved in the persistence, but it will be unavailable in the find endpoints. Later you can reactivate the user with the PUT endpoint.

Returns 200 with the deleted user if it was successful.

### Compile and run

First time? Get the required dependencies:

`go mod tidy`

Run the app:

`go run main.go`

### Run the tests

To run the test, use:

`go test -coverprofile=cover.out ./...`

You can generate the coverage report using the following command:

 `go tool cover -html=cover.out` or  `go tool cover -func=cover.out`

 ### Generate documentation (Swagger)

Install Swagoo cmd

`go install github.com/swaggo/swag/cmd/swag@latest`

Run `swag init` in order to generate and update documentation. Yoy will find Swagger endpoint visiting /docs/index.html

### Create the Docker container

Build the image

`docker build FULL_PATH_TO_PROJECT_DOCKERFILE -t IMAGE_NAME --build-arg ARTIFACT=example-api --build-arg VERSION=0.0.1-SNAPSHOT --build-arg PROFILE=DEVELOP`

You can change the environment values according the project, the version you want to deploy, and the project environment.

- IMAGE_NAME: The project name (Ex: example-api)
- ARTIFACT and VERSION: Created binary artifact name and artifact version
- PROFILE: The environment to execute. UPPERCASE (affects configurations)
- PORT: Application port (ex: 9090)

You can check this values later invoking the /health endpoint.

Run the image:

`docker run -p PORT:PORT IMAGE_NAME`

### Liveness endpoint

If you need to configure a liveness check for this api, you can use the health endpoint:

GET: `http://localhost:9090/health`

### Mongo DB document

Database: example
Collection: users

You can change this values in the config file.

Example users document:

`
db.users.insertOne({
    "_id": ObjectId("712d1e835ebee16872a109a4"),
    "reference": "1f047809-6869-41b4-9d2e-0423b9e4b2fc",
    "first_name": "Foo",
    "last_name": "Bar",
    "is_active": true,
    "email": "foobar@foobar.com.ar",
    "created_date": ISODate("2023-02-01T23:58:18Z"),
    "updated_date": ISODate("2023-02-01T23:58:18Z")
})
`

Since the api uses the field "references" as unique identificator, it's a good idea to set this field as unique index.

`
db.users.createIndex(
{
      "reference": 1
  },
  {
      unique: true
  }
)
`

### Q & A

TBD

### Changelog

0.0.1

* Initial version
