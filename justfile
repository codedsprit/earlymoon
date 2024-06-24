#Go parameters
GOCMD := "go"
GOBUILD := "go build"
GOMOD := "go mod"
GOTEST := "go test"
GOFLAGS := "-v"

# Check if the operating system is Darwin (macOS)
ldflags := `if [ $(go env GOOS) != "darwin" ]; then echo '-extldflags "-static"'; else echo '-s -w'; fi`

# Tasks
all: build

build:
  {{GOBUILD}} {{GOFLAGS}} -ldflags '{{ldflags}}' -o "earlymoon" cmd/earlymoon/main.go

tidy:
  {{GOMOD}} tidy
