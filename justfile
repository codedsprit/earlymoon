_default:
    @just -l

alias b := build
alias t := tidy

# Builds the binary
[linux]
build:
   @mkdir -p bin
   @go build -ldflags ' -s -w' -o "bin/earlymoon" cmd/earlymoon/main.go

# Builds  the binary
[macos]
build:
   @mkdir -p bin
   @go build -ldflags '-extldflags "-static"' -o "bin/earlymoon" cmd/earlymoon/main.go

tidy:
   @go mod tidy
