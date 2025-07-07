# Golang image based on Debian
FROM golang:1.18

# Set the Current Working Directory inside the container
WORKDIR /go/src/stajprojesi

# Label for metadata
LABEL authors="Merve_Bayram"

# Copy everything from the current directory to the container's WORKDIR
COPY . .

# Ensure go modules are enabled and tidy up the modules
RUN go mod tidy

# Use Go proxy (optional)
ENV GOPROXY=https://proxy.golang.org,direct

# Download dependencies
RUN go mod download

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]
