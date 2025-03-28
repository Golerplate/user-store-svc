FROM golang:1.23 as builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go get -u github.com/pressly/goose/v3/cmd/goose && \
    go install github.com/pressly/goose/v3/cmd/goose && \
    go install github.com/cosmtrek/air@v1.40.4 && \
    go mod download

# Copy the source code into the container
COPY internal/database/migrations /go/src/app/internal/database/migrations

COPY . .

# Set the command to run air for hot reloading
CMD ["air"]

# Expose the port your app runs on
EXPOSE 50051