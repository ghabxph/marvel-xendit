# Marvel Characters

Just a simple tool that fetches all characters from official marvel api.

## Requirements

* Either:
  * Go 1.15
  * Make (optional)
* Or:
  * Docker
  * Docker Compose
  
## Building and running the application

First, you need to be a [marvel developer](https://developer.marvel.com/) and get your app
key and secret. Then create a `config.yaml` file in this directory.

``` yaml
# Marvel App Key
public_key: "<your-marvel-app-key>"

# Marvel Secret Key
private_key: "<your-marvel-secret-key"
```

You got two options: Use docker or to build the application by yourself.  We  utilize  GNU
Make to make things very convenient for you. Note that before doing the next steps  below,
be sure that `config.yaml` file is present or things will simply fall apart.

### With Docker

Simply run `make dock-build dock-up` to build and run the docker container. If  you  don't
have GNU Make installed on your machine, then simply  execute  the  commands  manually  by
yourself.

``` sh

# Builds the docker container
docker-compose build

# Runs the docker container
docker-compose up
```

### Without Docker

Simply run `make clean build run` to create a clean build and run your application at  the
same time. If you don't have GNU Make installed on your machine, then simply  execute  the
commands manually by yourself.

``` sh

# Removes existing bin folder
rm -fvr ./bin

# Builds the application
go build -o bin/marvel ./cmd/marvel

# Runs the application
./bin/marvel
```

*(Shortcut: `make clean run` does the same thing, but I want to show you the full process)*
  
## Running the test

Simply run `make test` to run the test for the four important parts of our program, or if
you don't have GNU Make installed on your machine, then you can run the test for yourself.

``` sh

# Tests the HTTP endpoint
go test ./internal/gateway

# Tests our main business logic
go test ./internal/marvel

# Tests our in-memory cache
go test ./internal/memorydb

# Tests our scraper
go test ./internal/scraper
```

