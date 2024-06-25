FROM golang:alpine

# Set the Current Working Directory inside the container
WORKDIR /app/earlymoon

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build earlymoon
RUN go build -o earlymoon cmd/earlymoon/main.go 


# Ensure the binary is in the correct location and is executable
RUN ls -la /app/earlymoon/earlymoon


# Run earlymoon binary program produced by `go build`
CMD ["/app/earlymoon/earlymoon"]
