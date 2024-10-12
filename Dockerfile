FROM golang:1.20-alpine

WORKDIR /app

# Copy the go.mod and go.sum files if they already exist
COPY go.mod go.sum ./

# If go.mod and go.sum do not exist, initialize a new Go module
RUN if [ ! -f go.mod ]; then go mod init Company-Management; fi

# Download dependencies (if go.mod or go.sum are available)
RUN go mod tidy

# Copy the entire project (everything else)
COPY . .

RUN go build -o /Company-Management

EXPOSE 30000

# Run the executable
CMD ["/Company-Management"]
