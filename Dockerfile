FROM golang:1.15.8-alpine as build

# Copies the context
COPY . /app

# Sets the working dir
WORKDIR /app

# Compiles the binary
RUN go build -o bin/marvel ./cmd/marvel


# Excluding the source code
FROM golang:1.15.8-alpine

# Copies the binary
COPY --from=build /app/bin/marvel /app/marvel

# Copies the config
COPY --from=build /app/config.yaml /app/config.yaml

# Copies the swagger document
COPY --from=build /app/swagger.yaml /app/swagger.yaml

# Sets the workdir
WORKDIR /app

# Downgrade the user to non-root
USER 1000

# Exposes port
EXPOSE 8080

# Sets default command
CMD ["./marvel"]
