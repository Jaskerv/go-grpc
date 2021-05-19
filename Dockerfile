# Stage 1 - Start from golang base image
FROM golang:1.16.4-alpine as builder

# Set the current working directory inside the container
WORKDIR /build

# Copy go.mod, go.sum files and download deps
COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG project
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -a -v -o server $project

# Stage 2 - Start a new stage from busybox
FROM busybox:latest

# Set the current working directory inside the container
WORKDIR /dist

# Copy the build artifacts from the previous stage
COPY --from=builder /build/server .

# Run the executable
CMD ["./server"]
