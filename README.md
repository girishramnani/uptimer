# uptimer
A service which checks for availablity of 3rd party services like facebook, amazon etc.

## How to

- Start a postgres server or use the `docker-compose up -d` to start a dockerized version of postgres
- `go build main.go`
- `./main`
- Go to "http://localhost:8080/api/v1/services"
